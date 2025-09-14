package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-chat-server/internal/handler"
	"go-chat-server/internal/hub"
	"go-chat-server/pkg/config"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Create hub
	chatHub := hub.New()

	// Start hub in a goroutine
	go chatHub.Run()

	// Create handlers
	wsHandler := handler.NewWebSocketHandler(chatHub)
	staticHandler := handler.NewStaticHandler(cfg.TemplateDir, cfg.StaticDir)

	// Setup routes
	router := mux.NewRouter()

	// Home page
	router.HandleFunc("/", staticHandler.ServeHome)

	// WebSocket endpoint
	router.HandleFunc("/ws", wsHandler.ServeWS)

	// Static files
	router.PathPrefix("/static/").Handler(staticHandler.ServeStatic())

	// Create HTTP server
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on %s", server.Addr)
		log.Printf("Chat server running at http://%s", cfg.Address())
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
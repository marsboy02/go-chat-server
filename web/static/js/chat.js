class ChatClient {
    constructor() {
        this.ws = null;
        this.username = '';
        this.isConnected = false;
        this.userCount = 0;
        this.shouldAutoScroll = true;

        // DOM elements
        this.joinForm = document.getElementById('joinForm');
        this.chatContainer = document.getElementById('chatContainer');
        this.usernameInput = document.getElementById('usernameInput');
        this.joinBtn = document.getElementById('joinBtn');
        this.messageInput = document.getElementById('messageInput');
        this.sendBtn = document.getElementById('sendBtn');
        this.messages = document.getElementById('messages');
        this.connectionStatus = document.getElementById('connectionStatus');
        this.userCountEl = document.getElementById('user-count');

        this.initializeEventListeners();
    }

    initializeEventListeners() {
        // Join form events
        this.joinBtn.addEventListener('click', () => this.joinChat());
        this.usernameInput.addEventListener('keypress', (e) => {
            if (e.key === 'Enter') this.joinChat();
        });

        // Chat events
        this.sendBtn.addEventListener('click', () => this.sendMessage());
        this.messageInput.addEventListener('keypress', (e) => {
            if (e.key === 'Enter') this.sendMessage();
        });

        // Input validation
        this.usernameInput.addEventListener('input', () => this.validateJoinForm());
        this.messageInput.addEventListener('input', () => this.validateMessageForm());

        // Scroll detection for auto-scroll behavior
        this.setupScrollDetection();
    }

    validateJoinForm() {
        const username = this.usernameInput.value.trim();
        this.joinBtn.disabled = username.length === 0;
    }

    validateMessageForm() {
        const message = this.messageInput.value.trim();
        this.sendBtn.disabled = !this.isConnected || message.length === 0;
    }

    joinChat() {
        const username = this.usernameInput.value.trim();
        if (!username) return;

        this.username = username;
        this.connect();
    }

    connect() {
        const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
        const wsUrl = `${protocol}//${window.location.host}/ws?username=${encodeURIComponent(this.username)}`;

        this.updateConnectionStatus('connecting');

        try {
            this.ws = new WebSocket(wsUrl);
            this.setupWebSocketHandlers();
        } catch (error) {
            console.error('Failed to create WebSocket connection:', error);
            this.handleConnectionError();
        }
    }

    setupWebSocketHandlers() {
        this.ws.onopen = () => {
            console.log('WebSocket connected');
            this.isConnected = true;
            this.updateConnectionStatus('connected');
            this.showChatContainer();
            this.validateMessageForm();
        };

        this.ws.onmessage = (event) => {
            try {
                // Handle multiple messages separated by newlines
                const messages = event.data.split('\n');
                messages.forEach(messageData => {
                    if (messageData.trim()) {
                        const message = JSON.parse(messageData);
                        this.handleMessage(message);
                    }
                });
            } catch (error) {
                console.error('Error parsing message:', error, event.data);
            }
        };

        this.ws.onclose = (event) => {
            console.log('WebSocket closed:', event.code, event.reason);
            this.isConnected = false;
            this.updateConnectionStatus('disconnected');
            this.validateMessageForm();

            // Auto-reconnect after a delay (unless it was a normal closure)
            if (event.code !== 1000) {
                setTimeout(() => {
                    if (!this.isConnected) {
                        console.log('Attempting to reconnect...');
                        this.connect();
                    }
                }, 3000);
            }
        };

        this.ws.onerror = (error) => {
            console.error('WebSocket error:', error);
            this.handleConnectionError();
        };
    }

    handleMessage(message) {
        switch (message.type) {
            case 'chat':
                this.displayChatMessage(message);
                break;
            case 'join':
                this.displaySystemMessage(message);
                this.updateUserCount(1);
                break;
            case 'leave':
                this.displaySystemMessage(message);
                this.updateUserCount(-1);
                break;
            case 'error':
                this.displayErrorMessage(message);
                break;
            default:
                console.warn('Unknown message type:', message);
        }
    }

    displayChatMessage(message) {
        const messageEl = document.createElement('div');
        messageEl.className = `message ${message.username === this.username ? 'own' : 'other'}`;

        const contentEl = document.createElement('div');
        contentEl.className = 'message-content';
        contentEl.textContent = message.content;

        const infoEl = document.createElement('div');
        infoEl.className = 'message-info';

        const usernameEl = document.createElement('span');
        usernameEl.className = 'username';
        usernameEl.textContent = message.username;

        const timestampEl = document.createElement('span');
        timestampEl.className = 'timestamp';
        timestampEl.textContent = this.formatTimestamp(message.timestamp);

        infoEl.appendChild(usernameEl);
        infoEl.appendChild(timestampEl);

        messageEl.appendChild(contentEl);
        messageEl.appendChild(infoEl);

        this.messages.appendChild(messageEl);

        // Force scroll to bottom after adding message
        setTimeout(() => this.scrollToBottom(), 10);
    }

    displaySystemMessage(message) {
        const messageEl = document.createElement('div');
        messageEl.className = 'message system';

        const contentEl = document.createElement('div');
        contentEl.className = 'message-content';
        contentEl.textContent = message.content;

        messageEl.appendChild(contentEl);
        this.messages.appendChild(messageEl);

        // Force scroll to bottom after adding message
        setTimeout(() => this.scrollToBottom(), 10);
    }

    displayErrorMessage(message) {
        const messageEl = document.createElement('div');
        messageEl.className = 'message system';

        const contentEl = document.createElement('div');
        contentEl.className = 'message-content';
        contentEl.style.color = '#dc3545';
        contentEl.textContent = `⚠️ ${message.content}`;

        messageEl.appendChild(contentEl);
        this.messages.appendChild(messageEl);

        // Force scroll to bottom after adding message
        setTimeout(() => this.scrollToBottom(), 10);
    }

    sendMessage() {
        const content = this.messageInput.value.trim();
        if (!content || !this.isConnected) return;

        const message = {
            type: 'chat',
            content: content,
            username: this.username,
            timestamp: new Date().toISOString()
        };

        try {
            this.ws.send(JSON.stringify(message));
            this.messageInput.value = '';
            this.validateMessageForm();

            // Ensure auto-scroll when user sends a message
            this.shouldAutoScroll = true;
        } catch (error) {
            console.error('Failed to send message:', error);
            this.handleConnectionError();
        }
    }

    showChatContainer() {
        this.joinForm.style.display = 'none';
        this.chatContainer.style.display = 'flex';
        this.messageInput.focus();

        // Setup scroll detection after container is visible
        if (this.scrollDetectionSetup) {
            this.scrollDetectionSetup();
        }

        // Force scroll to bottom when chat starts
        this.shouldAutoScroll = true;
        setTimeout(() => this.scrollToBottom(), 100);
    }

    updateConnectionStatus(status) {
        this.connectionStatus.className = `status-${status}`;

        switch (status) {
            case 'connected':
                this.connectionStatus.textContent = '연결됨';
                break;
            case 'connecting':
                this.connectionStatus.textContent = '연결 중...';
                break;
            case 'disconnected':
                this.connectionStatus.textContent = '연결 끊김';
                break;
        }
    }

    updateUserCount(change) {
        this.userCount = Math.max(0, this.userCount + change);
        this.userCountEl.textContent = `${this.userCount} users online`;
    }

    handleConnectionError() {
        this.isConnected = false;
        this.updateConnectionStatus('disconnected');
        this.validateMessageForm();
    }

    setupScrollDetection() {
        // This will be called after chat container is shown
        this.scrollDetectionSetup = () => {
            const messagesContainer = this.messages.parentElement;
            if (messagesContainer) {
                messagesContainer.addEventListener('scroll', () => {
                    const { scrollTop, scrollHeight, clientHeight } = messagesContainer;
                    // Check if user is near the bottom (within 50px)
                    this.shouldAutoScroll = scrollTop + clientHeight >= scrollHeight - 50;
                });
            }
        };
    }

    scrollToBottom() {
        const messagesContainer = this.messages.parentElement;
        if (!messagesContainer) return;

        // Try multiple methods to ensure scrolling works
        const scrollToEnd = () => {
            // Method 1: scrollTop
            messagesContainer.scrollTop = messagesContainer.scrollHeight;

            // Method 2: scrollIntoView on last message
            const lastMessage = this.messages.lastElementChild;
            if (lastMessage) {
                lastMessage.scrollIntoView({ behavior: 'instant', block: 'end' });
            }
        };

        // Use both requestAnimationFrame and setTimeout for reliability
        requestAnimationFrame(scrollToEnd);
        setTimeout(scrollToEnd, 50);
    }

    formatTimestamp(timestamp) {
        const date = new Date(timestamp);
        return date.toLocaleTimeString('ko-KR', {
            hour: '2-digit',
            minute: '2-digit'
        });
    }
}

// Initialize the chat client when the page loads
document.addEventListener('DOMContentLoaded', () => {
    new ChatClient();
});

// Handle page visibility changes to manage connections
document.addEventListener('visibilitychange', () => {
    if (document.hidden) {
        // Page is hidden, we might want to reduce activity
        console.log('Page hidden');
    } else {
        // Page is visible again
        console.log('Page visible');
    }
});
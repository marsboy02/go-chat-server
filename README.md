# 💬 Go Chat Server

[![CI](https://github.com/YOUR_USERNAME/go-chat-server/actions/workflows/ci.yml/badge.svg)](https://github.com/YOUR_USERNAME/go-chat-server/actions/workflows/ci.yml)
[![Docker](https://github.com/YOUR_USERNAME/go-chat-server/actions/workflows/docker.yml/badge.svg)](https://github.com/YOUR_USERNAME/go-chat-server/actions/workflows/docker.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/YOUR_USERNAME/go-chat-server)](https://goreportcard.com/report/github.com/YOUR_USERNAME/go-chat-server)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

현대적인 실시간 채팅 서버로, Go와 WebSocket을 기반으로 구축되었으며, 반응형 웹 인터페이스와 Docker 컨테이너화를 지원합니다. Go 표준 프로젝트 레이아웃을 따르고 프로덕션 준비가 완료된 아키텍처를 제공합니다.

## ✨ 주요 기능

### 🌐 실시간 채팅

- **WebSocket 기반 양방향 통신**: 지연 없는 실시간 메시지 전송
- **자동 재연결**: 네트워크 끊김 시 자동 복구 기능
- **연결 상태 표시**: 실시간 연결 상태 모니터링
- **메시지 타입 시스템**: 채팅, 입장/퇴장, 시스템 메시지 구분

### 🎨 모던 웹 인터페이스

- **반응형 디자인**: 데스크톱부터 모바일까지 완벽 지원
- **부드러운 애니메이션**: 메시지 등장 효과와 스크롤 최적화
- **고정형 입력창**: 메시지가 많아도 항상 하단 고정
- **스마트 스크롤**: 사용자 의도에 맞는 자동/수동 스크롤

### 🏗️ 확장 가능한 아키텍처

- **Hub-and-Spoke 패턴**: 중앙집중식 연결 관리
- **Go 표준 프로젝트 구조**: cmd, internal, pkg 디렉토리 구조
- **인터페이스 기반 설계**: 테스트 가능하고 확장 가능한 구조
- **Graceful Shutdown**: 안전한 서버 종료 처리

### 🐳 완전한 컨테이너화

- **멀티스테이지 Docker 빌드**: 최적화된 프로덕션 이미지
- **개발/프로덕션 환경 분리**: docker-compose로 환경별 설정
- **보안 강화**: 비루트 사용자 실행, 최소 권한 원칙
- **헬스체크 내장**: 컨테이너 오케스트레이션 준비

## 🚀 빠른 시작

### 방법 1: Docker로 실행 (권장)

```bash
# 저장소 클론
git clone https://github.com/YOUR_USERNAME/go-chat-server.git
cd go-chat-server

# Docker Compose로 시작
make docker-up

# 브라우저에서 http://localhost:8080 접속
```

### 방법 2: 로컬 개발 환경

```bash
# 필요 조건: Go 1.20 이상 설치

# 저장소 클론 및 설정
git clone https://github.com/YOUR_USERNAME/go-chat-server.git
cd go-chat-server

# 의존성 설치
go mod tidy

# 서버 실행
make run

# 브라우저에서 http://localhost:8080 접속
```

### 사용법

1. 웹 브라우저에서 `http://localhost:8080` 접속
2. 사용자 이름 입력 후 '참여하기' 클릭
3. 채팅 메시지 입력 및 전송
4. 여러 탭이나 브라우저로 다중 사용자 테스트 가능

## 🏗️ 프로젝트 아키텍처

### 디렉토리 구조

```
go-chat-server/
├── cmd/server/                    # 애플리케이션 진입점
│   └── main.go                   # 서버 실행 및 설정
├── internal/                     # 프라이빗 애플리케이션 코드
│   ├── client/                   # WebSocket 클라이언트 관리
│   │   └── client.go             # 개별 클라이언트 연결 처리
│   ├── handler/                  # HTTP/WebSocket 핸들러
│   │   ├── websocket.go          # WebSocket 연결 처리
│   │   └── static.go             # 정적 파일 서빙
│   ├── hub/                      # 연결 허브 (브로드캐스트 시스템)
│   │   └── hub.go                # 중앙 메시지 관리
│   ├── message/                  # 메시지 타입 및 직렬화
│   │   └── message.go            # 메시지 구조체 정의
│   └── types/                    # 공통 인터페이스
│       └── interfaces.go         # Hub/Client 인터페이스
├── pkg/config/                   # 설정 관리
│   └── config.go                # 환경변수 처리
├── web/                         # 프론트엔드 에셋
│   ├── static/                  # CSS, JavaScript 파일
│   │   ├── css/style.css        # 반응형 채팅 UI 스타일
│   │   └── js/chat.js           # WebSocket 클라이언트 로직
│   └── templates/               # HTML 템플릿
│       └── index.html           # 메인 채팅 페이지
├── docker/                      # Docker 설정 파일
│   ├── .env.example             # 환경변수 템플릿
│   ├── docker-compose.dev.yml   # 개발 환경
│   ├── docker-compose.prod.yml  # 프로덕션 환경
│   └── redis.conf               # Redis 설정 (선택사항)
├── .github/                     # GitHub 설정
│   ├── workflows/               # CI/CD 파이프라인
│   ├── ISSUE_TEMPLATE/          # 이슈 템플릿
│   ├── CONTRIBUTING.md          # 기여 가이드
│   └── SECURITY.md              # 보안 정책
├── Dockerfile                   # 멀티스테이지 Docker 빌드
├── docker-compose.yml           # 기본 컨테이너 설정
├── Makefile                     # 개발 자동화 스크립트
└── .goreleaser.yaml             # 자동 릴리스 설정
```

### 핵심 컴포넌트

#### 🔄 Hub 패턴 (internal/hub/)

- **중앙집중식 관리**: 모든 WebSocket 연결을 한 곳에서 관리
- **스레드 세이프**: 고루틴과 채널을 활용한 동시성 처리
- **브로드캐스트**: 연결된 모든 클라이언트에게 메시지 전송
- **생명주기 관리**: 클라이언트 등록/해제 자동 처리

#### 🔗 클라이언트 관리 (internal/client/)

- **개별 연결 처리**: 각 WebSocket 연결의 독립적 관리
- **읽기/쓰기 펌프**: 별도 고루틴으로 메시지 송수신 처리
- **자동 정리**: 연결 해제 시 리소스 정리 및 허브 통지
- **에러 핸들링**: 네트워크 오류 및 예외 상황 처리

#### 💬 메시지 시스템 (internal/message/)

- **구조화된 메시지**: chat, join, leave, error 타입 구분
- **JSON 직렬화**: 클라이언트-서버 간 표준화된 통신
- **타임스탬프**: 모든 메시지에 정확한 시간 정보
- **타입 안전성**: Go의 타입 시스템을 활용한 메시지 검증

#### 🌐 웹 인터페이스 (web/)

- **모던 CSS**: Flexbox 기반 반응형 디자인
- **ES6+ JavaScript**: 클래스 기반 WebSocket 클라이언트
- **상태 관리**: 연결 상태, 사용자 수 실시간 업데이트
- **UX 최적화**: 스마트 스크롤, 자동 재연결, 입력 검증

## 🐳 Docker 배포

### 개발 환경

```bash
make docker-dev-up    # 개발 환경 시작 (소스 코드 마운트)
make docker-dev-logs  # 개발 환경 로그 확인
make docker-dev-down  # 개발 환경 종료
```

### 프로덕션 환경

```bash
make docker-prod-up    # Traefik & SSL이 포함된 프로덕션 환경 시작
make docker-prod-logs  # 프로덕션 로그 확인
make docker-prod-down  # 프로덕션 환경 종료
```

### Docker 특징

- **멀티스테이지 빌드**: Go 빌더와 scratch 기반 최종 이미지
- **보안 강화**: 비루트 사용자 실행, 최소 권한
- **헬스체크**: 컨테이너 상태 모니터링 내장
- **멀티아키텍처**: AMD64, ARM64 지원

## 🛠️ 개발 가이드

### 사용 가능한 명령어

```bash
# 개발 실행
make run              # 서버 시작
make dev              # 개발 모드로 시작
make test             # 테스트 실행
make test-coverage    # 커버리지와 함께 테스트 실행

# 코드 품질
make fmt              # 코드 포맷팅
make vet              # go vet 실행
make lint             # golangci-lint 실행
make check            # 모든 품질 검사 실행

# 빌드
make build            # 현재 플랫폼용 빌드
make build-all        # 모든 플랫폼용 빌드

# Docker
make docker-build     # Docker 이미지 빌드
make docker-up        # docker-compose로 시작
make docker-logs      # 컨테이너 로그 확인
```

### 개발 환경 요구사항

- **Go 1.20 이상** (최신 1.21 권장)
- **Docker & Docker Compose** (컨테이너 개발용)
- **Make** (빌드 자동화용, 선택사항)

### 환경 설정

환경변수를 통한 설정:

```bash
PORT=8080                    # 서버 포트
HOST=localhost               # 서버 호스트
DEV_MODE=true               # 개발 모드
TEMPLATE_DIR=web/templates   # HTML 템플릿 디렉토리
STATIC_DIR=web/static       # 정적 파일 디렉토리
```

### 개발 워크플로우

1. **저장소 포크 및 클론**
2. **기능 브랜치 생성**: `git checkout -b feature/your-feature`
3. **개발 서버 실행**: `make run` 또는 `make docker-dev-up`
4. **코드 변경 및 테스트**: `make test`, `make check`
5. **커밋 및 푸시**: Conventional Commits 형식 사용
6. **Pull Request 생성**: 제공된 템플릿 사용

## 🚀 프로덕션 배포

### GitHub Container Registry 활용

프로젝트는 GitHub Actions를 통해 자동으로 컨테이너 이미지를 빌드하고 배포합니다.

```bash
# 최신 이미지 사용
docker pull ghcr.io/YOUR_USERNAME/go-chat-server:latest

# 특정 버전 사용
docker pull ghcr.io/YOUR_USERNAME/go-chat-server:v1.0.0

# 컨테이너 실행
docker run -p 8080:8080 ghcr.io/YOUR_USERNAME/go-chat-server:latest
```

### 프로덕션 배포 단계

1. **환경 설정**: `docker/.env.example`을 `.env`로 복사 후 설정
2. **도메인 설정**: Traefik 라벨에서 도메인 구성
3. **SSL 설정**: Let's Encrypt 자동 인증서 설정
4. **실행**: `make docker-prod-up`으로 전체 스택 시작

### CI/CD 파이프라인

- **지속적 통합**: 푸시/PR시 자동 테스트 및 빌드
- **보안 스캔**: Gosec, Trivy를 통한 코드/이미지 취약점 검사
- **자동 릴리스**: 태그 푸시 시 GoReleaser로 자동 배포
- **멀티플랫폼**: Linux, Windows, macOS 바이너리 자동 생성

## 🧪 테스트

### 기본 테스트 실행

```bash
# 모든 테스트 실행
make test

# 커버리지와 함께 테스트 실행
make test-coverage

# 특정 패키지 테스트
go test -v ./internal/hub

# 레이스 컨디션 검출
go test -race ./...
```

### 테스트 구조

- **유닛 테스트**: 각 패키지별 핵심 로직 테스트
- **통합 테스트**: WebSocket 연결 및 메시지 전송 테스트
- **도커 테스트**: 컨테이너 빌드 및 실행 테스트
- **성능 테스트**: 대량 연결 및 메시지 처리 테스트

## 🤝 기여하기

### 기여 방법

1. **저장소 포크**: GitHub에서 프로젝트 포크
2. **기능 브랜치 생성**: `git checkout -b feature/amazing-feature`
3. **변경사항 개발**: 코드 작성 및 테스트
4. **품질 검사**: `make check` 실행으로 모든 검사 통과
5. **커밋**: Conventional Commits 형식으로 커밋
6. **푸시**: 본인 포크에 브랜치 푸시
7. **Pull Request**: 제공된 템플릿을 사용해 PR 생성

더 자세한 내용은 [기여 가이드](.github/CONTRIBUTING.md)를 참고하세요.

### 이슈 및 지원

- 🐛 [버그 신고](https://github.com/YOUR_USERNAME/go-chat-server/issues/new?template=bug_report.md)
- 💡 [기능 제안](https://github.com/YOUR_USERNAME/go-chat-server/issues/new?template=feature_request.md)
- ❓ [질문하기](https://github.com/YOUR_USERNAME/go-chat-server/issues/new?template=question.md)

## 📋 로드맵

### v1.x - 기본 기능 완성

- [x] **실시간 채팅**: WebSocket 기반 메시징
- [x] **반응형 UI**: 모바일/데스크톱 지원
- [x] **Docker 지원**: 컨테이너화 및 배포 자동화
- [x] **CI/CD**: GitHub Actions 파이프라인

### v2.x - 사용자 관리

- [ ] **사용자 인증**: 회원가입/로그인 시스템
- [ ] **사용자 프로필**: 아바타, 상태 메시지
- [ ] **권한 관리**: 관리자/사용자 역할

### v3.x - 고급 채팅 기능

- [ ] **채팅방/채널**: 다중 채팅방 지원
- [ ] **개인 메시지**: 1:1 메시징
- [ ] **메시지 히스토리**: 데이터베이스 연동 채팅 기록
- [ ] **파일 공유**: 이미지/파일 업로드 기능

### v4.x - 엔터프라이즈 기능

- [ ] **이모지 지원**: 커스텀 이모지 시스템
- [ ] **메시지 검색**: 전체 텍스트 검색
- [ ] **관리자 패널**: 채팅방 관리 인터페이스
- [ ] **API 제한**: Rate limiting 및 보안 강화
- [ ] **메시지 암호화**: End-to-end 암호화

## 📄 라이센스

이 프로젝트는 MIT 라이센스 하에 배포됩니다. 자세한 내용은 [LICENSE](LICENSE) 파일을 참고하세요.

## 🙏 감사의 말

### 핵심 라이브러리

- **[gorilla/websocket](https://github.com/gorilla/websocket)**: 안정적이고 성능이 뛰어난 WebSocket 구현
- **[gorilla/mux](https://github.com/gorilla/mux)**: 강력한 HTTP 라우터 및 URL 매처

### 개발 도구

- **GitHub Actions**: 자동화된 CI/CD 파이프라인
- **Docker**: 컨테이너화 및 배포 솔루션
- **GoReleaser**: 크로스 플랫폼 릴리스 자동화

### 커뮤니티

- **Go 커뮤니티**: 훌륭한 도구와 라이브러리 제공
- **오픈소스 생태계**: 프로젝트 구조와 모범 사례 참고

## 🌟 프로젝트 하이라이트

### 기술적 특징

- **🏗️ 확장 가능한 아키텍처**: Hub-and-Spoke 패턴으로 수천 개의 동시 연결 지원
- **⚡ 고성능**: Go의 고루틴과 채널을 활용한 효율적인 동시성 처리
- **🔒 보안 우선**: 비루트 컨테이너 실행, 입력 검증, CORS 설정
- **📱 크로스 플랫폼**: 웹, 모바일, 데스크톱 모든 환경 지원

### 개발자 경험

- **📖 명확한 문서화**: 상세한 README, 기여 가이드, API 문서
- **🧪 포괄적인 테스트**: 유닛, 통합, 성능 테스트 포함
- **🔄 자동화된 워크플로우**: 빌드, 테스트, 배포 완전 자동화
- **🎯 개발자 친화적**: Hot-reload, 도커 개발 환경, Makefile 지원

---

<div align="center">

## 💬 Go Chat Server

**현대적이고 확장 가능한 실시간 채팅 솔루션**

[![Go](https://img.shields.io/badge/Go-1.21-blue?style=for-the-badge&logo=go)](https://golang.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue?style=for-the-badge&logo=docker)](https://www.docker.com/)
[![WebSocket](https://img.shields.io/badge/WebSocket-Real--time-green?style=for-the-badge)](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API)

### 🚀 [데모 보기](http://localhost:8080) • 📖 [문서](CLAUDE.md) • 🤝 [기여하기](.github/CONTRIBUTING.md)

---

**이 프로젝트가 도움이 되었다면 ⭐ 스타를 눌러주세요!**

_Made with ❤️ and Go by the open source community_

</div>

# Talkie

Real-time walkie-talkie voice chat application. Record and send voice messages in chat rooms.

## Tech Stack

- **Backend:** Go (net/http), in-memory store
- **Frontend:** Vue 2, Bulma CSS, Vue Router, Composition API
- **Deployment:** Docker, Kubernetes

## Project Structure

```
gonorrhea/
  main.go              # CLI entry point
  go.mod
  Dockerfile
  server/
    server.go           # HTTP server, route wiring
  api/
    store.go            # In-memory data store
    auth.go             # Auth endpoints
    chats.go            # Messages + audio upload
    middleware.go       # CORS middleware
    health_check.go     # Health check types
  util/
    util.go             # Environment helpers
  ui/
    package.json
    src/
      main.js
      api.js            # API client (auth, messages, storage)
      App.vue
      components/
        Home.vue         # Login / register
        Login.vue        # Login form
        User.vue         # Auth state wrapper
        UserProfile.vue  # User info + sign out
        ChatRoom.vue     # Chat view (messages + voice recording)
        ChatMessage.vue  # Single message bubble
```

## Getting Started

### Prerequisites

- Go 1.20+
- Node.js 18+

### Backend

```bash
# Install dependencies
go mod download

# Run server (default port 3333)
go run main.go serve
```

Environment variables:

| Variable | Default | Description |
|---|---|---|
| `PORT` | `3333` | Server port |
| `DATA_DIR` | `./data` | Directory for uploaded audio files |

### Frontend

```bash
cd ui
npm install
npm run serve
```

The dev server runs on `http://localhost:8080` and proxies API calls to the backend.

### Docker

```bash
docker build -t talkie .
docker run -p 3333:3333 talkie
```

## API Endpoints

| Method | Path | Description |
|---|---|---|
| GET | `/` | Version info |
| POST | `/api/auth/anonymous` | Anonymous sign-in |
| POST | `/api/auth/register` | Email/password registration |
| POST | `/api/auth/login` | Email/password sign-in |
| POST | `/api/auth/logout` | Sign out |
| GET | `/api/chats/:id/messages?limit=N` | List messages |
| POST | `/api/chats/:id/messages` | Send message |
| POST | `/api/chats/:id/audio/:msgId` | Upload audio (multipart) |
| GET | `/files/audio/:chatId/:file` | Serve uploaded audio |

### Example: Send a message

```bash
curl -X POST http://localhost:3333/api/chats/room1/messages \
  -H "Content-Type: application/json" \
  -d '{"text": "Hello!", "sender": "user123"}'
```

### Example: Get messages

```bash
curl http://localhost:3333/api/chats/room1/messages?limit=5
```

## Kubernetes

See [docs/minikube.md](docs/minikube.md) for local cluster setup with Minikube.

Kubernetes manifests are in [docs/kubernetes.yaml](docs/kubernetes.yaml).

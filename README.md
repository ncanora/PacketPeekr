# Peekr 👀

A real-time local network monitoring dashboard. Peekr passively observes traffic on your network interface, tracks connected devices, flags unknown devices, and uses an LLM to explain suspicious behavior — all from a clean web UI.

> **Passive only** — Peekr never sends packets or probes your network. It only listens.

---

## Stack

| Layer | Technology |
|---|---|
| Packet Capture | Go + [gopacket](https://github.com/google/gopacket) |
| Backend API | Go + [Gin](https://gin-gonic.com/) |
| Real-time Events | Go + [Gorilla WebSockets](https://github.com/gorilla/websocket) |
| Database | SQLite via [go-sqlite3](https://github.com/mattn/go-sqlite3) |
| Frontend | React + TypeScript (Vite) |
| Styling | Tailwind CSS v3 |
| Charts | Recharts |
| LLM Analysis | Anthropic API (planned) |

---

## Prerequisites

### Windows
1. **Go 1.21+**
   - Download: https://golang.org/dl/
   - Verify: `go version`

2. **Node.js 22+**
   - Download: https://nodejs.org/
   - Or via nvm: `nvm install 22 && nvm use 22`
   - Verify: `node -v`

3. **Npcap** (required for packet capture on Windows)
   - Download: https://npcap.com/#download
   - During install, check **"WinPcap API Compatible Mode"**
   - Leave all other options default

### macOS / Linux
Npcap is not needed — libpcap is built in. Just install Go and Node.

---

## Getting Started

### 1. Clone the repo
```bash
git clone https://github.com/ncanora/peekr.git
cd peekr
```

### 2. Backend
```bash
cd backend
go mod tidy
go run cmd/peekr/main.go
```
Backend runs on `http://localhost:8080`

> **Windows note:** Must run as Administrator for packet capture permissions.

### 3. Frontend
```bash
cd frontend
npm install
npm run dev
```
Frontend runs on `http://localhost:5173`

---

## Project Structure

```
peekr/
├── backend/
│   ├── cmd/peekr/        # Entry point
│   └── internal/
│       ├── capture/      # Packet capture engine (gopacket)
│       ├── devices/      # SQLite device store
│       └── api/          # REST + WebSocket server (Gin)
├── frontend/
│   └── src/
│       ├── components/   # React UI components
│       └── hooks/        # Custom React hooks (WebSocket, fetch)
└── README.md
```

---

## Features (Roadmap)

- [x] Project scaffold
- [ ] Packet capture — MAC/IP/protocol detection
- [ ] Device tracking — first seen, last seen, traffic volume
- [ ] Unknown device alerts via WebSocket
- [ ] REST API — device list, traffic stats
- [ ] React dashboard — live device table, traffic graphs
- [ ] LLM behavioral analysis — explain suspicious patterns
- [ ] Local anomaly detection model (CICIDS dataset)

---

## Notes

- Peekr only inspects packet **headers** (MAC, IP, port, protocol flags). It never reads packet payloads.
- Built and tested on Windows 11 with Npcap. Also compatible with Linux/macOS.
- Run the backend as Administrator on Windows to allow raw packet capture.

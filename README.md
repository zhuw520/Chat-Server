# Chat Server

## Project Structure
chat-server/
├── go.mod
├── main.go
├── server/
│   ├── server.go
│   └── connection.go
├── handler/
│   └── websocket.go
├── model/
│   ├── message.go
│   └── user.go
├── middleware/
│   └── rate_limit.go
├── storage/
│   └── memory_store.go
├── utils/
│   ├── ip_utils.go
│   └── validator.go
├── config/
│   └── constants.go
├── protocol/
│   └── message_proto.go
└── monitor/
    └── stats_collector.go

## Features
- WebSocket real-time chat
- User connection management
- Message broadcasting
- Memory message storage (120 max)
- Rate limiting (6 messages/min)
- IP filtering
- Message length validation (78 chars)
- Automatic message cleanup
- Online user statistics

## Quick Start

1. Change directory:
cd /storage/emulated/0/chat-server

2. Download dependencies:
go mod tidy

3. Run server:
go run main.go

4. Server info:
- WebSocket: ws://localhost:8092/ws
- HTTP: http://localhost:8092/

## Message Protocol

Join chat:
{"type":"join","userId":"user123"}

Send message:
{"type":"message","userId":"user123","message":"hello","timestamp":"15:04"}

## Configuration

Default settings:
- Port: 8092
- Max messages: 120
- Rate limit: 6/min
- Message limit: 78 chars
- Ban duration: 300 seconds
- Timeout: 90 seconds

## Testing

Test WebSocket connection:
wscat -c ws://localhost:8092/ws

Test HTTP:
curl http://localhost:8092/

## Building

Build binary:
go build -o chat-server main.go

Run binary:
./chat-server

## Notes
- Pure memory storage (no database)
- No message persistence
- No admin panel
- No file upload
- No user authentication
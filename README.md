# Go Chat – Real-time Chat Application API

Go Chat is a backend API for a real-time chat application. It provides user authentication, message management, and WebSocket-based real-time communication, allowing users to send and receive messages instantly.

---

## Features

- **User Authentication**: Secure sign up and log in with JWT tokens.
- **Real-time Messaging**: WebSocket-based instant messaging between users.
- **Message Management**: Send, view, and manage chat messages with pagination.
- **User Management**: User registration and profile management.
- **Healthcheck**: Simple endpoint to check if the server is running.

---

## Tech Stack

- **Language**: Go (Golang) 1.24.0
- **Framework**: Echo (HTTP web framework)
- **Database**: PostgreSQL (via GORM ORM)
- **Authentication**: JWT (JSON Web Tokens)
- **Real-time Communication**: WebSocket (Gorilla WebSocket)
- **Validation**: Go Playground Validator
- **Other**: godotenv, pgx driver

---

## Architecture

The project follows a **layered architecture** pattern:

- **Models**: Data structures and database logic (`models/`)
- **Controllers**: Request handling and business logic (`controllers/`)
- **Services**: Business logic and data processing (`services/`)
- **Routes**: API endpoint definitions (`routes/`)
- **Middleware**: Authentication and request pre-processing (`middleware/`)
- **DTOs**: Data Transfer Objects for API requests/responses (`dto/`)
- **Sockets**: WebSocket handling for real-time communication (`sockets/`)
- **Utils**: Utility functions (JWT, password hashing, etc.)

---

## API Endpoints

All endpoints (except `/ping`, `/login`, `/signup`, `/chat`) require a valid JWT in the `Authorization: Bearer <token>` header.

### Healthcheck

- **GET `/ping`**
  - **Response**: `"pong"`

### Authentication

- **POST `/signup`**
  - **Body**: `{ "email": string, "password": string, "name": string }`
  - **Response**: `{ "data": { "user": { ... }, "token": string } }`

- **POST `/login`**
  - **Body**: `{ "email": string, "password": string }`
  - **Response**: `{ "data": { "user": { ... }, "token": string } }`

### Messages

- **GET `/messages`**
  - **Query Parameters**: 
    - `offset` (optional): Number of messages to skip (default: 0)
    - `limit` (optional): Number of messages to return (default: 20)
  - **Response**: `{ "data": { "messages": [ ... ] } }`

### WebSocket

- **GET `/chat`** (WebSocket endpoint)
  - **Description**: Real-time messaging endpoint
  - **Protocol**: WebSocket
  - **Message Format**: JSON with action and message content

---

## WebSocket Message Format

Messages sent through the WebSocket should follow this format:

```json
{
  "id": 0,
  "message": "Hello, world!",
  "action": "add"
}
```

- add `id` for edit, delete

Available actions:
- `add` - Send a new message
- `edit` - Edit an existing message
- `delete` - Delete a message

---

## Environment Variables

Create a `.env` file in the project root with the following variables:

- `DB_URL` – PostgreSQL connection string (e.g., `postgres://user:password@localhost:5432/gochat`)
- `JWT_KEY` – Secret key for signing JWT tokens
- `PORT` – (Optional) Port for the server (default: `:8080`)

---

## Getting Started

1. **Clone the repository**
2. **Install dependencies**:  
   ```
   go mod download
   ```
3. **Set up your `.env` file** (see above)
4. **Run the server**:  
   ```
   go run main.go
   ```
5. **API is now available at** `http://localhost:8080` (or your specified port)

---

## Database Schema

### Users Table
- `id` (Primary Key)
- `name` (String)
- `email` (String, Unique)
- `password` (String, Hashed)

### Messages Table
- `id` (Primary Key)
- `message` (String)
- `sender_id` (Foreign Key to Users)
- `created_at` (Timestamp)
- `updated_at` (Timestamp)

---

## Real-time Features

The application uses WebSocket connections to provide real-time messaging:

- **Hub**: Manages all connected WebSocket clients
- **Client**: Represents individual WebSocket connections
- **Broadcasting**: Messages are broadcast to all connected clients
- **Connection Management**: Automatic client registration/unregistration

---

## Testing

The project includes test files for authentication and message functionality:

- `controllers/auth_test.go` - Authentication controller tests
- `controllers/messages_test.go` - Message controller tests
- `mock/` - Mock implementations for testing

Run tests with:
```
go test ./...
```

---

## License

MIT
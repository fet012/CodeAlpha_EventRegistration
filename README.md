# CodeAlpha Event Registration System

A REST API for managing events and user registrations, built with Go, Gin, and MongoDB.

## Tech Stack
- **Language:** Go
- **Framework:** Gin
- **Database:** MongoDB
- **Auth:** JWT

## Getting Started

### Prerequisites
- Go 1.21+
- MongoDB Atlas account

### Installation
1. Clone the repo
```bash
   git clone https://github.com/fet012/CodeAlpha_EventRegistration
   cd CodeAlpha_EventRegistration
```

2. Install dependencies
```bash
   go mod tidy
```

3. Create a `.env` file in the root directory
```env
   MONGODB_URI=your_mongodb_uri
   DB_NAME=event_registration
   JWT_SECRET=your_jwt_secret
   ADMIN_EMAIL=admin@example.com
   ADMIN_PASSWORD=yourpassword
   PORT=8080
```

4. Run the server
```bash
   go run main.go
```

## API Endpoints

### Public
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/admin/login` | Admin login, returns JWT |
| GET | `/events/:id` | Get event details |
| POST | `/events/:id/register` | Register for an event |

### Protected (Admin only — requires Bearer token)
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/events` | Create a new event |
| GET | `/events` | List all events |
| GET | `/events/:id/registrations` | View registrations for an event |


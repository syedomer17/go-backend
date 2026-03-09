# Go Task Manager Backend

A full-featured RESTful API backend built with Go (Golang) for task management with user authentication, email verification, and JWT-based authorization.

## 🚀 Features

### Authentication & Authorization
- ✅ User registration with email verification
- ✅ OTP-based email verification using Redis
- ✅ Secure login with JWT tokens (Access & Refresh tokens)
- ✅ Token refresh mechanism
- ✅ Password hashing with bcrypt
- ✅ Cookie-based authentication
- ✅ Protected routes with middleware

### Task Management
- ✅ Create, read, update, and delete tasks
- ✅ User-specific task isolation
- ✅ Task completion status tracking
- ✅ Pagination and filtering support

### Additional Features
- ✅ Rate limiting middleware
- ✅ Request logging with Uber Zap logger
- ✅ CORS configuration
- ✅ MongoDB database integration
- ✅ Redis caching for OTP storage
- ✅ Email notification system
- ✅ Health check endpoint

## 🛠️ Tech Stack

- **Language**: Go 1.25.0
- **Web Framework**: Gin
- **Database**: MongoDB
- **Cache**: Redis (Upstash)
- **Authentication**: JWT (golang-jwt/jwt)
- **Password Hashing**: bcrypt
- **Logger**: Uber Zap
- **Other**: CORS, Rate Limiting

## 📦 Dependencies

```go
- github.com/gin-gonic/gin - Web framework
- github.com/gin-contrib/cors - CORS middleware
- go.mongodb.org/mongo-driver - MongoDB driver
- github.com/redis/go-redis/v9 - Redis client
- github.com/golang-jwt/jwt/v5 - JWT authentication
- golang.org/x/crypto - Password hashing
- github.com/joho/godotenv - Environment variable management
- github.com/ulule/limiter/v3 - Rate limiting
- go.uber.org/zap - Structured logging
```

## 📁 Project Structure

```
server/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   ├── db.go                # MongoDB connection
│   │   ├── env.go               # Environment configuration
│   │   └── redis.go             # Redis connection
│   ├── controllers/
│   │   ├── auth_controller.go   # Authentication handlers
│   │   └── task_controller.go   # Task management handlers
│   ├── middleware/
│   │   ├── auth_middleware.go   # JWT authentication middleware
│   │   ├── rate_limit.go        # Rate limiting middleware
│   │   └── request_logger.go    # Request logging middleware
│   ├── models/
│   │   ├── user.go              # User model
│   │   └── task.go              # Task model
│   ├── routes/
│   │   └── routes.go            # API routes definition
│   └── utils/
│       ├── email.go             # Email sending utility
│       ├── hash.go              # Password hashing utility
│       ├── jwt.go               # JWT token generation/validation
│       └── otp.go               # OTP generation utility
├── pkg/
│   └── logger/
│       └── logger.go            # Logger configuration
├── go.mod                       # Go module dependencies
└── go.sum                       # Dependency checksums
```

## ⚙️ Environment Variables

Create a `.env` file in the `server/` directory with the following variables:

```env
# Server Configuration
PORT=8080

# MongoDB
MONGO_URI=mongodb://localhost:27017/task_manager
# or use MongoDB Atlas
# MONGO_URI=mongodb+srv://username:password@cluster.mongodb.net/task_manager

# JWT Secret
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production

# Redis (Upstash)
UPSTASH_REDIS_REST_URL=https://your-redis-url.upstash.io
UPSTASH_REDIS_REST_TOKEN=your-redis-token

# SMTP Email Configuration
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your-email@gmail.com
SMTP_PASS=your-app-specific-password
EMAIL_FROM=noreply@yourdomain.com
```

## 🚦 Installation & Setup

### Prerequisites
- Go 1.25+ installed
- MongoDB running locally or MongoDB Atlas account
- Redis instance (Upstash or local)
- SMTP credentials for email sending

### Steps

1. **Clone the repository**
```bash
git clone git@github.com:syedomer17/go-backend.git
cd go-backend
```

2. **Navigate to server directory**
```bash
cd server
```

3. **Install dependencies**
```bash
go mod download
```

4. **Create environment file**
```bash
cp .env.example .env
# Edit .env with your configuration
```

5. **Run the server**
```bash
go run cmd/api/main.go
```

Or build and run:
```bash
go build -o bin/api cmd/api/main.go
./bin/api
```

The server will start on `http://localhost:8080` (or your configured PORT)

## 📡 API Endpoints

### Health Check
```
GET /health - Check server status
```

### Authentication Routes (`/auth`)

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/auth/register` | Register new user | No |
| POST | `/auth/verify-email` | Verify email with OTP | No |
| POST | `/auth/login` | Login user | No |
| POST | `/auth/refresh` | Refresh access token | No |
| POST | `/auth/logout` | Logout user | No |

### Task Routes (`/tasks`)

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/tasks` | Create new task | Yes |
| GET | `/tasks` | Get all user tasks | Yes |
| PUT | `/tasks/:id` | Update task | Yes |
| DELETE | `/tasks/:id` | Delete task | Yes |

## 📝 API Usage Examples

### 1. Register User
```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

**Response:**
```json
{
  "message": "user register successfully"
}
```
*An OTP will be sent to the provided email*

### 2. Verify Email
```bash
curl -X POST http://localhost:8080/auth/verify-email \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "otp": "123456"
  }'
```

**Response:**
```json
{
  "message": "email verify successfully"
}
```

### 3. Login
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -c cookies.txt \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

**Response:**
```json
{
  "message": "Login Successfull"
}
```
*Sets `accessToken` and `refreshToken` cookies*

### 4. Create Task
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -b cookies.txt \
  -d '{
    "title": "Complete project documentation",
    "description": "Write comprehensive README for the project"
  }'
```

**Response:**
```json
{
  "task": {
    "id": "507f1f77bcf86cd799439011",
    "title": "Complete project documentation",
    "description": "Write comprehensive README for the project",
    "completed": false,
    "userId": "507f191e810c19729de860ea",
    "createdAt": "2026-03-09T10:30:00Z",
    "updatedAt": "2026-03-09T10:30:00Z"
  }
}
```

### 5. Get All Tasks
```bash
curl -X GET http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -b cookies.txt
```

**Response:**
```json
{
  "tasks": [
    {
      "id": "507f1f77bcf86cd799439011",
      "title": "Complete project documentation",
      "description": "Write comprehensive README for the project",
      "completed": false,
      "userId": "507f191e810c19729de860ea",
      "createdAt": "2026-03-09T10:30:00Z",
      "updatedAt": "2026-03-09T10:30:00Z"
    }
  ]
}
```

### 6. Update Task
```bash
curl -X PUT http://localhost:8080/tasks/507f1f77bcf86cd799439011 \
  -H "Content-Type: application/json" \
  -b cookies.txt \
  -d '{
    "title": "Complete project documentation",
    "description": "Write comprehensive README for the project",
    "completed": true
  }'
```

### 7. Delete Task
```bash
curl -X DELETE http://localhost:8080/tasks/507f1f77bcf86cd799439011 \
  -H "Content-Type: application/json" \
  -b cookies.txt
```

### 8. Refresh Token
```bash
curl -X POST http://localhost:8080/auth/refresh \
  -H "Content-Type: application/json" \
  -b cookies.txt
```

### 9. Logout
```bash
curl -X POST http://localhost:8080/auth/logout \
  -H "Content-Type: application/json" \
  -b cookies.txt
```

## 🔒 Security Features

- **Password Hashing**: Passwords are hashed using bcrypt before storage
- **JWT Authentication**: Access tokens expire in 24 hours, refresh tokens in 7 days
- **HTTP-Only Cookies**: Tokens stored in HTTP-only cookies to prevent XSS attacks
- **Rate Limiting**: Prevents brute force attacks
- **CORS Configuration**: Restricts API access to allowed origins
- **OTP Expiration**: Email verification OTPs expire in 10 minutes

## 🏗️ Database Schema

### Users Collection
```javascript
{
  _id: ObjectId,
  name: String,
  email: String (unique, indexed),
  password: String (hashed),
  isVerified: Boolean,
  createdAt: DateTime
}
```

### Tasks Collection
```javascript
{
  _id: ObjectId,
  title: String,
  description: String,
  completed: Boolean,
  userId: ObjectId (reference to Users),
  createdAt: DateTime,
  updatedAt: DateTime
}
```

## 🧪 Testing

Run the server in development mode:
```bash
go run cmd/api/main.go
```

For production build:
```bash
go build -o bin/api cmd/api/main.go
./bin/api
```

## 🐛 Error Handling

The API uses standard HTTP status codes:

- `200` - Success
- `400` - Bad Request (validation errors)
- `401` - Unauthorized (authentication required)
- `404` - Not Found
- `500` - Internal Server Error

Error response format:
```json
{
  "error": "Error message description"
}
```

## 🚀 Deployment

### Production Considerations

1. **Environment Variables**: Use secure environment variable management
2. **HTTPS**: Enable HTTPS in production
3. **Database**: Use MongoDB Atlas or managed MongoDB service
4. **Redis**: Use Upstash or managed Redis service
5. **CORS**: Configure allowed origins for your frontend
6. **Rate Limiting**: Adjust limits based on your needs
7. **Logging**: Configure log levels and storage

### Docker Deployment (Optional)

Create a `Dockerfile`:
```dockerfile
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY server/go.* ./
RUN go mod download
COPY server/ ./
RUN go build -o main cmd/api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```

Build and run:
```bash
docker build -t go-task-manager .
docker run -p 8080:8080 --env-file .env go-task-manager
```

## 📄 License

MIT License - feel free to use this project for learning and commercial purposes.

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!

## 📧 Contact

For questions or support, please open an issue in the repository.

---

**Built with ❤️ using Go** 

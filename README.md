<div align="center">

# Sagara Booking API

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-v2-00ACD7?style=for-the-badge&logo=fiber&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=for-the-badge&logo=postgresql&logoColor=white)
![Heroku](https://img.shields.io/badge/Heroku-430098?style=for-the-badge&logo=heroku&logoColor=white)

**A production-ready RESTful API for sports field booking management**

[Live Demo](https://sagara-booking-api-f264e78236b6.herokuapp.com/api) • [API Documentation](https://sagara-booking-api-f264e78236b6.herokuapp.com/swagger/index.html) • [Report Issue](https://github.com/HIUNCY/sagara-booking-api/issues)

</div>

---

## 📋 Table of Contents

- [Overview](#-overview)
- [Key Features](#-key-features)
- [Technology Stack](#-technology-stack)
- [System Architecture](#-system-architecture)
- [Getting Started](#-getting-started)
- [API Documentation](#-api-documentation)
- [Testing](#-testing)
- [Deployment](#-deployment)
- [Contributing](#-contributing)
- [License](#-license)

---

## 🎯 Overview

Sagara Booking API is a robust, enterprise-grade backend solution designed for managing sports field bookings. Built with Go and following clean architecture principles, this API provides a scalable foundation for booking management systems with comprehensive authentication, authorization, and business logic validation.

**Developed as part of the Backend Developer Assessment for PT Sagara Asia Teknologi**

### Live Endpoints

- **Base URL**: `https://sagara-booking-api-f264e78236b6.herokuapp.com/api`
- **Swagger Documentation**: `https://sagara-booking-api-f264e78236b6.herokuapp.com/swagger/index.html`

> **Note**: The application runs on Heroku's free tier. Initial requests after periods of inactivity may experience a cold start delay of 10-20 seconds.

---

## ✨ Key Features

### Authentication & Authorization
- 🔐 **JWT-based Authentication** - Secure token-based authentication system
- 👥 **Role-Based Access Control (RBAC)** - Granular permissions for admin and user roles
- 🔑 **Password Encryption** - Industry-standard password hashing

### Field Management
- ✅ **Complete CRUD Operations** - Full create, read, update, delete functionality
- 🏟️ **Public Field Listing** - Anonymous access to view available fields
- 🛡️ **Admin-Only Modifications** - Protected endpoints for field management

### Booking System
- 📅 **Smart Scheduling** - Automatic overlap detection and prevention
- 🔄 **Status Management** - Structured booking lifecycle (pending → paid)
- ⚡ **Real-time Validation** - Instant feedback on booking conflicts

### Payment Integration
- 💳 **Mock Payment Gateway** - Simulated payment processing for testing
- 📊 **Transaction Tracking** - Complete payment history and status updates

### Code Quality
- 🏗️ **Clean Architecture** - Separation of concerns with clear layer boundaries
- 📝 **Comprehensive Documentation** - Auto-generated Swagger/OpenAPI specs
- ✅ **Unit Testing** - Test coverage for critical business logic
- 🐳 **Docker Support** - Containerized deployment ready

---

## 🛠 Technology Stack

| Category | Technology |
|----------|------------|
| **Language** | Go 1.21+ |
| **Web Framework** | Fiber v2 |
| **Database** | PostgreSQL (Neon Cloud) |
| **ORM** | GORM |
| **Authentication** | JWT (golang-jwt) |
| **Documentation** | Swagger/OpenAPI (swaggo) |
| **Deployment** | Heroku |
| **Testing** | Go testing package |

---

## 🏗 System Architecture

### Project Structure

```
sagara-booking-api/
│
├── cmd/
│   └── api/
│       └── main.go              # Application entry point & route configuration
│
├── internal/
│   ├── core/
│   │   ├── domain/              # Business entities (User, Field, Booking)
│   │   └── port/                # Interfaces, DTOs, and contracts
│   │
│   ├── handler/                 # HTTP request handlers (presentation layer)
│   ├── service/                 # Business logic implementation
│   └── repository/              # Data access layer (GORM implementations)
│
├── pkg/
│   ├── database/                # Database connection & configuration
│   ├── middleware/              # JWT authentication & authorization
│   └── util/                    # Utility functions (hashing, token generation)
│
├── docs/                        # Auto-generated Swagger documentation
├── .env.example                 # Environment variable template
├── Dockerfile                   # Container configuration
├── Procfile                     # Heroku deployment configuration
└── README.md
```

### Architecture Layers

```
┌─────────────────────────────────────────┐
│         HTTP/REST Interface             │
│            (Fiber Handlers)             │
├─────────────────────────────────────────┤
│         Business Logic Layer            │
│              (Services)                 │
├─────────────────────────────────────────┤
│        Data Access Layer                │
│           (Repositories)                │
├─────────────────────────────────────────┤
│            Database                     │
│          (PostgreSQL)                   │
└─────────────────────────────────────────┘
```

---

## 🚀 Getting Started

### Prerequisites

Ensure you have the following installed:

- **Go**: Version 1.21 or higher ([Download](https://golang.org/dl/))
- **PostgreSQL**: Version 12+ ([Download](https://www.postgresql.org/download/))
- **Git**: For cloning the repository
- **Docker** (Optional): For containerized deployment

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/HIUNCY/sagara-booking-api.git
   cd sagara-booking-api
   ```

2. **Install dependencies**
   ```bash
   go mod download
   go mod verify
   ```

3. **Configure environment variables**

   Create a `.env` file in the project root:
   ```bash
   cp .env.example .env
   ```

   Update the `.env` file with your configuration:
   ```env
   # Database Configuration
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASSWORD=your_secure_password
   DB_NAME=sagara_booking
   DB_PORT=5432
   
   # JWT Configuration
   JWT_SECRET=your_jwt_secret_key_min_32_chars
   
   # Server Configuration (Optional)
   PORT=8080
   ```

4. **Initialize the database**

   The application will automatically run migrations on startup. Ensure your PostgreSQL server is running and the database exists.

### Running the Application

#### Local Development

```bash
go run cmd/api/main.go
```

The API will be available at:
- **Base URL**: `http://localhost:8080/api`
- **Swagger UI**: `http://localhost:8080/swagger/index.html`

#### Docker Deployment

1. **Build the Docker image**
   ```bash
   docker build -t sagara-booking-api:latest .
   ```

2. **Run the container**
   ```bash
   docker run -d \
     --name sagara-booking-api \
     -p 8080:8080 \
     -e DB_HOST=your_db_host \
     -e DB_USER=postgres \
     -e DB_PASSWORD=your_password \
     -e DB_NAME=sagara_booking \
     -e DB_PORT=5432 \
     -e JWT_SECRET=your_jwt_secret \
     sagara-booking-api:latest
   ```

3. **View logs**
   ```bash
   docker logs -f sagara-booking-api
   ```

---

## 📚 API Documentation

### Authentication Endpoints

| Method | Endpoint | Description | Authentication |
|--------|----------|-------------|----------------|
| `POST` | `/api/register` | Register a new user or admin | Public |
| `POST` | `/api/login` | Authenticate and receive JWT token | Public |

### Field Management Endpoints

| Method | Endpoint | Description | Required Role |
|--------|----------|-------------|---------------|
| `GET` | `/api/fields` | Retrieve all available fields | Public |
| `GET` | `/api/fields/:id` | Get detailed field information | Public |
| `POST` | `/api/fields` | Create a new field | Admin |
| `PUT` | `/api/fields/:id` | Update field information | Admin |
| `DELETE` | `/api/fields/:id` | Remove a field | Admin |

### Booking Endpoints

| Method | Endpoint | Description | Required Role |
|--------|----------|-------------|---------------|
| `POST` | `/api/bookings` | Create a new booking (with overlap validation) | User/Admin |
| `GET` | `/api/bookings` | Retrieve user's booking history | User/Admin |
| `GET` | `/api/bookings/:id` | Get specific booking details | User/Admin |

### Payment Endpoints

| Method | Endpoint | Description | Required Role |
|--------|----------|-------------|---------------|
| `POST` | `/api/payments` | Process payment for a booking (mock) | User/Admin |

### Importing Postman Collection

A ready-to-use Postman collection is included in the repository:

```bash
Sagara Booking API.postman_collection.json
```

**Import Instructions:**
1. Open Postman
2. Click **Import** in the top-left corner
3. Select the JSON file from the project directory
4. Configure the environment variables for your setup

---

## 🧪 Testing

The project includes comprehensive unit tests for critical components.

### Run All Tests

```bash
go test ./...
```

### Run Tests with Coverage

```bash
go test ./... -cover
```

### Generate Coverage Report

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Test Structure

- **Handler Tests**: HTTP request/response validation
- **Service Tests**: Business logic verification
- **Utility Tests**: Helper function validation
- **Mock Testing**: No external dependencies required

---

## 🚢 Deployment

### Heroku Deployment

The application is configured for Heroku deployment with an included `Procfile`.

1. **Create a Heroku app**
   ```bash
   heroku create your-app-name
   ```

2. **Set environment variables**
   ```bash
   heroku config:set DB_HOST=your_neon_db_host
   heroku config:set DB_USER=your_db_user
   heroku config:set DB_PASSWORD=your_db_password
   heroku config:set DB_NAME=your_db_name
   heroku config:set DB_PORT=5432
   heroku config:set JWT_SECRET=your_jwt_secret
   ```

3. **Deploy**
   ```bash
   git push heroku main
   ```

4. **Open your application**
   ```bash
   heroku open
   ```

### Updating Swagger Documentation

If you modify API handlers, regenerate Swagger documentation:

```bash
# Install swag CLI
go install github.com/swaggo/swag/cmd/swag@latest

# Generate documentation
swag init -g cmd/api/main.go -o docs
```

---

## 🤝 Contributing

Contributions are welcome! Please follow these guidelines:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Code Standards

- Follow Go conventions and best practices
- Write tests for new features
- Update documentation as needed
- Keep commits atomic and descriptive

---

## 📄 License

This project is developed as part of an assessment for PT Sagara Asia Teknologi.

---

## 👨‍💻 Author

**Muhamad Zainul Kamal**

- GitHub: [@HIUNCY](https://github.com/HIUNCY)
- Project Link: [https://github.com/HIUNCY/sagara-booking-api](https://github.com/HIUNCY/sagara-booking-api)

---

## 🙏 Acknowledgments

- PT Sagara Asia Teknologi for the opportunity
- Go Fiber community for excellent documentation
- GORM team for the robust ORM solution

---

<div align="center">

**Built with ❤️ using Go and Fiber**

</div>
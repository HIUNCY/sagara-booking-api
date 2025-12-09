# Sagara Tech - Backend Developer Intern Test

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-v2-black?style=flat)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Neon-336791?style=flat&logo=postgresql)
![Heroku](https://img.shields.io/badge/Deployed-Heroku-430098?style=flat&logo=heroku)

This repository contains the solution for the **Backend Developer Take Home Test** at PT. Sagara Asia Teknologi. The project is a RESTful API for a sports venue booking system, built using **Go (Golang)**, **Fiber**, and **Clean Architecture**.

## 🌐 Live Demo & Documentation

The application is deployed on Heroku and includes interactive Swagger documentation.

* **Base URL:** `https://sagara-booking-api-f264e78236b6.herokuapp.com/api`
* **Swagger UI (Test API Here):** [Click to Open Swagger Documentation](https://sagara-booking-api-f264e78236b6.herokuapp.com/swagger/index.html)

> **Note:** The server might sleep on inactivity (Heroku Free Tier). Please wait 10-20 seconds for the first request.

## 🔑 Testing Credentials

You can use these credentials to test Role-Based Access Control (RBAC):

| Role | Email | Password | Permissions |
| :--- | :--- | :--- | :--- |
| **Admin** | `admin@sagara.id` | `admin123` | Can Create/Edit/Delete Fields, View Bookings |
| **User** | `user@test.com` | `user123` | Can View Fields, Create Booking |

*Or you can register a new user via `POST /api/register`.*

## ✨ Key Features

### 1. Authentication & Authorization
* **JWT Implementation:** Secure login with Bearer Token.
* **RBAC Middleware:** Strict separation between `admin` and `user` capabilities.

### 2. Field Management (CRUD)
* **Admin:** Full access to Create, Update, and Delete sports fields.
* **Public:** Users can view the list of available fields.

### 3. Booking System (Core Logic)
* **Overlap Validation:** The system strictly prevents double booking. It checks if a requested time slot overlaps with any existing booking for the specific field.
* **Status Management:** Default status is `pending`.

### 4. Payment (Bonus Feature)
* Mock endpoint to update booking status from `pending` to `paid`.

### 5. Clean Architecture
The project follows the separation of concerns principle:
* **Handler:** HTTP transport layer (Fiber).
* **Service:** Business logic layer.
* **Repository:** Data access layer (GORM).
* **Core/Port:** Domain entities and interface definitions.

## 🛠️ Tech Stack

* **Language:** Go (Golang)
* **Framework:** Fiber v2
* **Database:** PostgreSQL (Cloud via Neon Tech)
* **ORM:** GORM
* **Documentation:** Swagger (Swaggo)
* **Deployment:** Heroku

## 📂 Project Structure

```text
├── cmd/api/main.go          # Application Entry Point & Route Config
├── internal/
│   ├── core/
│   │   ├── domain/          # Database Entities (User, Field, Booking)
│   │   └── port/            # Interfaces & DTOs (Request/Response structs)
│   ├── handler/             # HTTP Handlers
│   ├── service/             # Business Logic (Validation, Calculation)
│   └── repository/          # Database Queries
├── pkg/
│   ├── database/            # Postgres Connection
│   ├── middleware/          # JWT & Role Middleware
│   └── util/                # Helper functions
└── docs/                    # Swagger Generated Files
```

## 🚀 Local Installation
If you want to run this project locally:
### 1. Clone the repository
```Bash
git clone https://github.com/HIUNCY/sagara-booking-api
cd sagara-booking-api
```

### 2. Install Dependencies
```Bash
go mod tidy
```

### 3. Setup Environment Variables
* Copy ```.env.example``` file to ```.env```
```bash
cp .env.example .env
```

* Open ```.env``` file and adjust the values ​​to your local configuration, especially for the database connection 
```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=passwordmu
DB_NAME=sagara_booking
DB_PORT=5432
JWT_SECRET=rahasia_negara_sagara
```

### 4. Run the Application
```Bash
go run cmd/api/main.go
```
Access Swagger at: http://localhost:8080/swagger/index.html
## 📡 API Endpoints Overview
### Auth
| Method | Endpoint | Description | Auth |
| :--- | :--- | :--- | :--- |
| **POST** | `/api/register` | Register New User / Admin | Public |
| **POST** | `/api/login` | Login & Get Token | Public |
### Fields
| Method | Endpoint | Description | Auth |
| :--- | :--- | :--- | :--- |
| **GET** | `/api/fields` | List All Fields | Public |
| **GET** | `/api/fields/:id` | Get Field Detail | Public |
| **POST** | `/api/fields/` | Create New Field | Admin |
| **PUT** | `/api/fields/:id` | Update Field | Admin |
| **DELETE** | `/api/fields/:id` | Delete Field | Admin |
### Bookings
| Method | Endpoint | Description | Auth |
| :--- | :--- | :--- | :--- |
| **POST** | `/api/bookings` | Create Booking (Check Overlap) | User/Admin |
| **GET** | `/api/bookings` | Get Booking History | User/Admin |
| **GET** | `/api/bookings/:id` | Get Booking Detail | User/Admin |
| **POST** | `/api/payments` | Pay Booking (Mock) | User/Admin |

---
Author: **Muhamad Zainul Kamal**

<!-- ========================================================= -->
<!--        E-Learning Management System (Go + PostgreSQL)    -->
<!-- ========================================================= -->

# 📚 E-Learning Management System

<p align="center">
  <img src="https://readme-typing-svg.herokuapp.com?font=Fira+Code&size=26&pause=1000&color=00C2FF&center=true&vCenter=true&width=900&lines=Scalable+E-Learning+Backend+Built+With+Go;PostgreSQL+Powered+REST+API;Clean+Architecture+%7C+Dockerized+%7C+SaaS+Ready" />
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
  <img src="https://img.shields.io/badge/PostgreSQL-15+-336791?style=for-the-badge&logo=postgresql&logoColor=white" />
  <img src="https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker&logoColor=white" />
  <img src="https://img.shields.io/badge/Architecture-Clean-blue?style=for-the-badge" />
  <img src="https://img.shields.io/github/license/coderemon24/E-Learning?style=for-the-badge" />
</p>

---

# 🚀 Overview

A modern, scalable, RESTful **E-Learning Management System** built with:

- 🧠 Go (Golang)
- 🗄 PostgreSQL
- 🐳 Docker
- 🔐 JWT Authentication
- 🧱 Clean Architecture

Designed for:
- Online Course Platforms
- University LMS
- SaaS-based Learning Systems
- API-first education platforms

---

# ✨ Core Features

## 👤 Authentication
- User Registration
- Login System
- JWT Token Authentication
- Secure Password Hashing (bcrypt)

## 🎓 Course Management
- Create / Update / Delete Course
- Course Categories
- Instructor Assignment

## 📚 Module & Lesson Structure
- Course → Modules → Lessons
- Hierarchical Learning System

## 📝 Enrollment System
- Student Course Enrollment
- Enrollment Validation
- Access Control

## 📊 Progress Tracking
- Lesson Completion Tracking
- Course Progress Percentage

## 🔐 Role-Based Access
- Admin
- Instructor
- Student

---

# 🧱 Architecture
Client (Web / Mobile)
|
v
REST API Layer (Handlers / Controllers)
|
v
Service Layer (Business Logic)
|
v
Repository Layer (Database Access)
|
v
PostgreSQL Database


Clean layered architecture ensures:
- Maintainability
- Scalability
- Testability
- Clear separation of concerns

---

# 🗂 Project Structure

```
E-Learning/
├── cmd/
│   └── main.go                # Application entry point
│
├── internal/
│   ├── config/                # Configuration & environment setup
│   ├── controllers/           # HTTP handlers (request/response layer)
│   ├── middlewares/           # Authentication & custom middleware
│   ├── models/                # Database models / structs
│   ├── repositories/          # Database interaction layer
│   ├── routes/                # Route definitions
│   └── services/              # Business logic layer
│
├── docker/                    # Docker configuration files
├── docker-compose.yml         # Multi-container Docker setup
├── go.mod                     # Go module definition
└── README.md                  # Project documentation
```

# 🐳 Docker Setup (Recommended)

## 1️⃣ Clone Repository

```bash
git clone https://github.com/coderemon24/E-Learning.git
cd E-Learning
docker compose up -d --build

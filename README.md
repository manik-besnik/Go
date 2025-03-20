# 🚀 Go Learning Routine (10 Minutes a Day)

> **Goal:** Learn Go by building real-world projects in small, focused steps.

---

## 📅 Weekly Schedule

### **Day 1 - Syntax & Basics**
- `var` / `const` / `:=` declarations
- Data types (int, string, bool, slices, maps)
- Basic control flow (`if`, `for`, `switch`)

### **Day 2 - Functions & Structs**
- Functions with multiple return values
- Structs (Go’s version of classes)
- Pointers in Go (`*` and `&`)

### **Day 3 - Packages & Modules**
- `go mod init`
- Importing built-in and custom packages
- Folder structure (like Laravel’s organization)

### **Day 4 - HTTP Server**
- Create a simple HTTP server with `net/http`
- Handle GET and POST requests

### **Day 5 - REST API Routes**
- Set up routes with `gorilla/mux` or `chi`
- Implement simple CRUD handlers

### **Day 6 - Database Integration**
- Connect to MySQL/Postgres with `gorm` or `database/sql`
- Basic DB operations (Create, Read)

### **Day 7 - Go Routines & Channels**
- Learn Go concurrency with `go` keyword
- Simple example with `go routines` and channels

---

## 🔁 Repeat & Expand (Week 2+)

- Week 2+: Expand the API (Auth, Middleware, Error handling)
- Week 3+: CLI project or Microservice patterns
- Optional: Testing with `testing` and `testify`

---

# 🗓️ Go Learning Routine - Weeks 2, 3 & 4

---

## 📅 **Week 2 - Building a REST API (Intermediate)**

### **Day 8 - Middleware**
- Build a simple logging middleware
- Learn about `http.Handler` and chaining middleware

### **Day 9 - CRUD - Create & Read**
- Create `POST` endpoint to insert data into DB
- Create `GET` endpoint to fetch single/multiple records

### **Day 10 - CRUD - Update & Delete**
- Implement `PUT` (update) and `DELETE` handlers

### **Day 11 - Models & Validation**
- Create a `User` or `Task` struct as your model
- Add basic validation (e.g., check for empty fields)

### **Day 12 - Environment Variables**
- Load config (DB creds, ports) using `os.Getenv` or a `.env` loader package like `joho/godotenv`

### **Day 13 - Refactor Handlers**
- Organize your project:
  - `/handlers`
  - `/models`
  - `/routes`
  - `/main.go`

### **Day 14 - Error Handling**
- Centralize error handling
- Return JSON error responses with proper HTTP status codes

---

## 📅 **Week 3 - Advanced API & Testing**

### **Day 15 - Dependency Injection**
- Pass DB connection to handlers
- Learn how to keep handlers testable

### **Day 16 - JSON Marshalling/Unmarshalling**
- Learn how to decode incoming JSON (`json.NewDecoder`)
- Encode responses using `json.Marshal`

### **Day 17 - Custom Response Structs**
- Create a standardized API response (e.g., `{ "data": ..., "message": ..., "error": ... }`)

### **Day 18 - Authentication (JWT)**
- Use a JWT package (e.g., `github.com/golang-jwt/jwt`)
- Implement a simple login route with token generation

### **Day 19 - Protect Routes (Auth Middleware)**
- Middleware to validate JWTs and protect certain endpoints

### **Day 20 - Unit Testing Basics**
- Write unit tests using Go's built-in `testing` package
- Test simple functions and handlers

### **Day 21 - API Testing**
- Write API tests with `net/http/httptest`

---

## 📅 **Week 4 - Final Touches & Extra**

### **Day 22 - Logging**
- Add structured logging with `log` or `logrus`

### **Day 23 - Pagination**
- Implement basic pagination logic in your `GET` endpoints

### **Day 24 - Dockerize the App**
- Create a `Dockerfile` to containerize your Go app

### **Day 25 - Graceful Shutdown**
- Handle `SIGINT` and gracefully shut down your HTTP server

### **Day 26 - CLI Tool Bonus**
- Build a simple CLI utility with Go's `flag` package

### **Day 27 - Go Routines Refresher**
- Revisit Go routines with a real-world example (e.g., background job)

### **Day 28 - Deployment**
- Deploy your app to a VPS or cloud platform (e.g., Render, DigitalOcean, Fly.io)

---

## ✅ **End Goal**
- By the end of Week 4, you’ll have:
  - A production-ready REST API or microservice
  - Working knowledge of Go best practices
  - Intro to Go testing, concurrency, and deployment

---

## ⏰ Pro Tip
- Stay consistent, even **10 mins/day** is enough!
- Use the official docs: [https://go.dev/doc](https://go.dev/doc)

---

## 🎯 Goal
By the end of 4 weeks, you’ll be comfortable building Go APIs or CLI tools and deploying them.

---

# Go API with Redis, OTP, JWT, and Rate Limiting

This project demonstrates how to build a REST API in Go with:

- Redis for OTP storage and rate limiting
- JWT for access tokens and refresh tokens
- Gin for routing and middleware
- Rate limiting to prevent brute force attacks

## Running the project

```bash
docker-compose up --build

API will be available at: http://localhost:8080

Endpoints

POST /login — send OTP
POST /verify — verify OTP and get tokens
POST /refresh — refresh access token
POST /logout — logout
GET /me — protected route

Usage examples:

get otp:

% curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"phone": "79991234567"}

verify otp:
curl -X POST http://localhost:8080/verify -H "Content-Type: application/json" -d '{"phone": "79991234567", "otp": "543269"}'
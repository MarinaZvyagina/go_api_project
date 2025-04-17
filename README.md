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


# RSS Aggregator

A simple RSS aggregator built with Go for learning purposes.

## Purpose

This project is designed to learn Go basics including:
- HTTP server setup with Chi router
- Database operations with PostgreSQL
- SQLC for type-safe SQL queries
- API authentication with API keys
- HashiCorp Vault integration for secure database credentials

## Features

- User registration and authentication
- RSS feed management
- API key-based authentication
- Secure database credential management with Vault

## API Endpoints

### Users
- `POST /v1/users` - Create a new user
- `GET /v1/users` - Get current user (requires API key)

### Health Check
- `GET /v1/healthy` - Health check endpoint
- `GET /v1/err` - Error testing endpoint

## Authentication

API endpoints require an API key in the Authorization header:
```
Authorization: ApiKey your-api-key-here
```

## Getting Started

1. Set up your environment variables (PORT, database credentials)
2. Configure HashiCorp Vault for database secrets
3. Run database migrations
4. Start the server:
   ```bash
   go run .
   ```

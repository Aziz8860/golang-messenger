# Golang Simple Messenger
### Architecture
![architecture](https://github.com/user-attachments/assets/e58188eb-2d7a-4004-8ef5-0dea47c9fc72)

### List of APIs
- POST /user/v1/register
- POST /user/v1/login
- PUT /user/v1/refresh-token
- DELETE /user/v1/logout
- GET /message/v1/history
- GET /message/v1/send

### Packages used
- Fiber
- Gorm
- JWT

### Getting Started
1. Run docker compose

```
cd elk_stack
docker-compose up -d
```

2. Run main.go

```
go run main.gp
```

Learning from fastcampus

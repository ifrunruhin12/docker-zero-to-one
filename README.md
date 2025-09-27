# Docker Learning Journey 🐳

This is my docker learning journey where I'm learning Docker through playing a game which starts from level 0 and ends on level 69.

## 🎮 Game Structure

Each level focuses on specific Docker concepts and gradually builds complexity:

- **Levels 0-10**: Docker basics (containers, images, Dockerfiles)
- **Levels 11-30**: Advanced concepts (volumes, networks, compose)
- **Levels 31-50**: Production scenarios (multi-stage builds, orchestration)
- **Levels 51-69**: Expert topics (security, optimization, scaling)

## 📁 Repository Structure

```
docker-learning-journey/
├── docker-level3/          # Basic Ubuntu container with tools
├── docker-level6/          # Simple Go web application
├── docker-level8/          # Port mapping and SSE experiments
├── docker-level10/         # Full-stack app with PostgreSQL + Nginx
└── mydata/                 # Persistent volume experiments
```

## 🚀 Key Learning Projects

### Level 10: Full-Stack Application

Complete Docker Compose setup with:

- **PostgreSQL** database with initialization scripts
- **Go API** server with REST endpoints (`/users`, `/ping`)
- **Nginx** reverse proxy handling traffic routing
- **Custom network** for inter-container communication

**Try it:** `cd docker-level10 && docker-compose up`

### Level 8: Port Experiments

Exploration of Docker port mapping, SSE streaming, and container networking fundamentals.

### Level 6: First Web App

Simple Go web server containerized and exposed through Docker port mapping.

## 🛠️ Technologies Used

- **Docker & Docker Compose** - Containerization
- **Go** - Backend applications
- **PostgreSQL** - Database
- **Nginx** - Reverse proxy and web server
- **Ubuntu** - Base images for custom containers

## 📚 What I've Learned

- Container lifecycle management
- Dockerfile best practices
- Volume mounting and data persistence
- Network configuration and container communication
- Multi-container orchestration with Docker Compose
- Port mapping and reverse proxy setup
- Database initialization and connection handling

## 🏃‍♂️ Quick Start

```bash
# Clone the repository
git clone https://github.com/ifrunruhin12/docker-zero-to-one.git
cd docker-zero-to-one

# Try the full-stack app
cd docker-level10
docker-compose up

# Access the application
curl http://localhost:8080/api/ping  # Returns {"message": "pong"}
curl http://localhost:8080/api/users # Returns user list from database
```

---

_"Learning Docker one level at a time!"_ 🎯

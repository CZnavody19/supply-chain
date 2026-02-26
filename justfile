default:
    @just --list

# Start all services
up:
    docker compose up -d

# Stop all services
down:
    docker compose down

# Rebuild and start (use after modifying Dockerfile or dependencies)
build:
    docker compose up --build -d

# Follow logs for all services
logs:
    docker compose logs -f

# Follow logs for frontend only
logs-fe:
    docker compose logs -f frontend

# Follow logs for backend only
logs-be:
    docker compose logs -f backend

# Restart all services
restart:
    docker compose restart

# Wipe everything, including Neo4j database volumes
clean:
    docker compose down -v
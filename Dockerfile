# Stage 1: Frontend build
FROM node:22-alpine AS frontend-build

WORKDIR /frontend

COPY frontend/package*.json ./
RUN npm install

COPY frontend/ ./
RUN npm run build

# Stage 2: Backend build
FROM golang:1.24.2 AS backend-build

WORKDIR /app

COPY api/ ./
RUN apt-get update && apt-get install -y gcc libc6-dev
RUN go build -o server .

# Stage 3: Final image with GCC for CGO
FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc libc6-dev && \
    rm -rf /var/lib/apt/lists/*

COPY --from=backend-build /app/server .
COPY --from=frontend-build /frontend/dist ./dist

EXPOSE 8080

CMD ["./server"]
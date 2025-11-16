FROM node:22-alpine AS frontend-build

WORKDIR /frontend

COPY ./frontend .

RUN npm install

RUN npm run build

FROM golang:1.25.3 AS backend-build

WORKDIR /app

COPY go.mod go.sum main.go router.go ./

COPY core ./core

COPY --from=frontend-build /frontend/dist ./frontend/dist

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o recipebook .

FROM gcr.io/distroless/static-debian12

COPY --from=backend-build /app/recipebook .

EXPOSE 8080

CMD ["./recipebook"]
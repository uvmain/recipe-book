# recipebook Web Application

This is a **recipebook Web Application** built with a **Go REST API backend** and a **Vite/Vue frontend**. The application uses **SQLite** to store recipes and is designed to be served on a local or hosted server with a reverse proxy for routing.

## Features

- Go REST API backend
- Vue.js frontend built with Vite
- SQLite database for image metadata storage and auth session storage
- SSL/TLS certificate generation for local development

## Project Structure

- **api/** - Go backend REST API
- **frontend/** - Vite/Vue3 frontend application

## Prerequisites

- Go (tested on 1.24)
- Node.js & npm (tested on Node 22 LTS)

## Getting Started

1. **Clone the repository:**

```bash
git clone https://github.com/uvmain/gallery
cd gallery
```

2. **Install dependencies:**
```bash
npm run deps
```
This command downloads the npm dependencies for the local dev and frontend components, and the go dependencies for the api component.

3. **Set up environment variables:**
Update the following environment variables in the /package.json (for local dev) and in the docker_compose.yml (or sibling .env file):
```plaintext
ADMIN_USER=your_admin_name
ADMIN_PASSWORD=your_admin_password
```

4. **Generate SSL/TLS certificates (optional for local development):**
```bash
npm run create-cert
```

5. **Start development environment::**
```bash
npm run dev
```
This command:
- Starts the Caddy server for reverse proxy
- Runs the frontend in development mode (HMR)
- Runs the backend in development mode

Open [recipebook.localhost](https://[recipebook.localhost) in your browser

## Scripts

- `npm run backend:dev` - Runs the Go backend in development mode.
- `npm run frontend:dev` - Runs the frontend in development mode.
- `npm run dev` - Starts Caddy and concurrently runs the frontend and backend.
- `npm run build` - Builds both the backend and frontend for production.
- `npm run create-cert` - Generates and installs certificates for local SSL.
- `npm run preview` - Runs the built distributable with Caddy for SSL and routing.

## Deployment

```
docker build -t recipebook:latest . && docker compose up -d
```

## Contributing

Feel free to submit issues or pull requests for improvements.

## License

This project is licensed under the MIT License.
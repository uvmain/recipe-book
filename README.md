## Setup

Make sure to install the dependencies:

```bash
# npm
npm install
```

## Development Server

Start the development server on `http://localhost:3000`:

```bash
# npm
npm run dev

```

## Production

Build the application for production:

```bash
# npm
npm run build

```

## Docker

Build a docker image:
```bash
git rev-parse HEAD | xargs -I % docker build -t recipebook:% .
```

Run a docker image:
```bash
git rev-parse HEAD | xargs -I % docker run --publish "0.0.0.0:3000:3000" recipebook:% .
```

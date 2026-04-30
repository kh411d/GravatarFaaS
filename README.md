# GravatarFaaS
Providing Gravatar API Function as a Service

## Go public proxy for Gravatar V3

This service exposes a public HTTP endpoint and forwards requests to the Gravatar V3 Profiles API.

### Endpoints

- **Local development**: `GET /gravatar/{profileIdentifier}` on `http://localhost:8080`
- **Vercel deployment**: `GET /api/gravatar/{profileIdentifier}`

### Local Development

```bash
export GRAVATAR_API_KEY="your_gravatar_api_key"
go run main.go
curl http://localhost:8080/gravatar/205e460b479e2e5b48aec07710c08d50
```

### Vercel Deployment

1. Install Vercel CLI:
   ```bash
   npm install -g vercel
   ```

2. Deploy:
   ```bash
   vercel
   ```

3. Set environment variable:
   ```bash
   vercel env add GRAVATAR_API_KEY
   ```

4. Access the deployed function:
   ```bash
   curl https://your-project.vercel.app/api/gravatar/205e460b479e2e5b48aec07710c08d50
   ```

### Build

```bash
go build -o gravatarfaas
```

### Run

```bash
export GRAVATAR_API_KEY="your_gravatar_api_key"
./gravatarfaas
```

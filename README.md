# GravatarFaaS
Providing Gravatar API Function as a Service

## Go public proxy for Gravatar V3

This service exposes a public HTTP endpoint and forwards requests to the Gravatar V3 Profiles API.

### Endpoints

- **Local development**: `GET /profile` on `http://localhost:3000`
- **Vercel deployment**: `GET /api/profile`

### Local Development

```bash
export GRAVATAR_API_KEY="your_gravatar_api_key" GRAVATAR_PROFILE_ID="your_gravatar_profile_id"
go run main.go
curl http://localhost:3000/profile
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
   vercel env add GRAVATAR_PROFILE_ID
   ```

4. Access the deployed function:
   ```bash
   curl https://your-project.vercel.app/api/profile
   ```

### Build

```bash
go build -o gravatarfaas
```

### Run

```bash
export GRAVATAR_API_KEY="your_gravatar_api_key" GRAVATAR_PROFILE_ID="your_gravatar_profile_id"
./gravatarfaas
```

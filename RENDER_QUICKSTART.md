# Render Quick Start

## Penting: Render Free Tier Sleep

Render akan **SLEEP** setelah 15 menit tidak ada traffic. Ini berarti:
- ❌ Scraper tidak bisa jalan di Render (akan sleep)
- ✅ API bisa jalan di Render (wake up saat ada request)

## Solusi: Hybrid Setup

- **Scraper** → GitHub Actions (gratis, scheduled)
- **API** → Render (gratis, on-demand)
- **Database** → MongoDB Atlas (gratis)

## 5-Minute Setup

### 1. Setup MongoDB Atlas (3 min)
- Go to https://www.mongodb.com/cloud/atlas
- Create free account
- Create cluster (Free M0)
- Get connection string

### 2. Setup GitHub Actions untuk Scraper (2 min)
```bash
# Push ke GitHub
cd comic-scraper
git init
git add .
git commit -m "Initial commit"
git remote add origin https://github.com/YOUR_USERNAME/comic-scraper.git
git push -u origin main

# Di GitHub repo:
# Settings → Secrets → Actions
# Add: MONGODB_URI, BASE_URL
```

Workflow sudah siap di `.github/workflows/scrape.yml`

### 3. Deploy API ke Render (5 min)
1. Go to https://render.com
2. Sign up (gratis, no CC)
3. Click "New +" → "Web Service"
4. Connect GitHub repository
5. Configure:
   - Name: `comic-api`
   - Environment: `Go`
   - Build Command: `go build -o api server.go api.go database.go models.go`
   - Start Command: `./api`
6. Add Environment Variables:
   - `MONGODB_URI`: (dari MongoDB Atlas)
   - `BASE_URL`: `https://komiku.id/`
   - `PORT`: `8080`
7. Click "Deploy Web Service"

### 4. Test API (1 min)
```bash
# Render akan berikan URL seperti: https://comic-api.onrender.com
curl https://comic-api.onrender.com/api/stats
curl https://comic-api.onrender.com/api/comics
```

## Total Time: ~10-15 minutes

## Keep API Awake (Optional)

Karena Render sleep, setup Uptime Robot (gratis):
1. Go to https://uptimerobot.com
2. Create monitor
3. URL: `https://comic-api.onrender.com/api/stats`
4. Interval: 5 minutes

## Cost: $0

- GitHub Actions: Free
- MongoDB Atlas: Free
- Render: Free

## Architecture

```
GitHub Actions (Scraper) → MongoDB Atlas ← Render (API)
     (Scheduled)              (Shared)            (On-demand)
```

## Next Steps

1. Test scraper manual dari GitHub Actions tab
2. Test API endpoints
3. (Opsional) Setup Uptime Robot
4. (Opsional) Build frontend

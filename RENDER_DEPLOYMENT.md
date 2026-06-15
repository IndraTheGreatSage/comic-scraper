# Render Deployment Guide

## Important: Render Free Tier Limitation

**Render Free Tier akan SLEEP setelah 15 menit tidak ada traffic.**

Ini berarti:
- ❌ **Tidak cocok untuk scraper** (butuh jalan terus)
- ✅ **Cocok untuk API server** (hanya aktif saat ada request)
- ✅ **Gratis** (tidak butuh credit card untuk free tier)

## Solusi Hybrid (Recommended)

Karena Render tidak cocok untuk scraper, gunakan kombinasi ini:

1. **Scraper** → GitHub Actions (gratis, scheduled)
2. **API Server** → Render (gratis, on-demand)
3. **Database** → MongoDB Atlas (gratis, shared)

Ini memberikan:
- Gratis 100%
- Scraper jalan otomatis setiap hari
- API aktif saat dibutuhkan
- Tidak ada sleep issue untuk scraper

## Setup Guide

### Step 1: Setup MongoDB Atlas

1. Go to https://www.mongodb.com/cloud/atlas
2. Create free account
3. Create cluster (Free M0 tier)
4. Create database user
5. Network Access → Add IP: `0.0.0.0/0`
6. Get connection string

### Step 2: Deploy Scraper ke GitHub Actions

Lihat `GITHUB_ACTIONS_DEPLOYMENT.md` untuk setup lengkap.

Singkatnya:
```bash
# Push ke GitHub
git init
git add .
git commit -m "Initial commit"
git remote add origin https://github.com/IndraTheGreatSage/comic-scraper.git
git push -u origin main

# Add secrets di GitHub:
# - MONGODB_URI
# - BASE_URL
```

Workflow sudah disiapkan di `.github/workflows/scrape.yml`

### Step 3: Deploy API ke Render

#### Option A: Go API (Recommended)

1. **Prepare Project**
```bash
cd comic-scraper
```

2. **Create render.yaml**
```yaml
services:
  - type: web
    name: comic-api
    env: node
    buildCommand: go build -o api server.go api.go database.go models.go
    startCommand: ./api
    envVars:
      - key: MONGODB_URI
        sync: false
      - key: BASE_URL
        value: https://komiku.id/
      - key: PORT
        value: 8080
```

3. **Push to GitHub**
```bash
git add render.yaml
git commit -m "Add Render config"
git push
```

4. **Deploy ke Render**
   - Go to https://render.com
   - Sign up (gratis, no CC untuk free tier)
   - Click "New +"
   - Select "Web Service"
   - Connect GitHub repository
   - Select `comic-scraper` repo
   - Configure:
     - Name: `comic-api`
     - Environment: `Go`
     - Build Command: `go build -o api server.go api.go database.go models.go`
     - Start Command: `./api`
   - Add Environment Variables:
     - `MONGODB_URI`: (dari MongoDB Atlas)
     - `BASE_URL`: `https://komiku.id/`
     - `PORT`: `8080`
   - Click "Deploy Web Service"

#### Option B: Node.js API

Gunakan project `comic-scraper-nodejs`:

1. **Prepare Project**
```bash
cd comic-scraper-nodejs
```

2. **Create render.yaml**
```yaml
services:
  - type: web
    name: comic-api
    env: node
    buildCommand: npm run build
    startCommand: npm start
    envVars:
      - key: MONGODB_URI
        sync: false
      - key: BASE_URL
        value: https://komiku.id/
```

3. **Push to GitHub**
```bash
git init
git add .
git commit -m "Initial commit"
git remote add origin https://github.com/YOUR_USERNAME/comic-scraper-nodejs.git
git push -u origin main
```

4. **Deploy ke Render**
   - Go to https://render.com
   - Click "New +"
   - Select "Web Service"
   - Connect GitHub repository
   - Select `comic-scraper-nodejs` repo
   - Configure:
     - Name: `comic-api`
     - Environment: `Node`
     - Build Command: `npm run build`
     - Start Command: `npm start`
   - Add Environment Variables:
     - `MONGODB_URI`: (dari MongoDB Atlas)
     - `BASE_URL`: `https://komiku.id/`
   - Click "Deploy Web Service"

### Step 4: Test API

Setelah deploy selesai, Render akan memberikan URL seperti:
`https://comic-api.onrender.com`

Test endpoints:
```bash
# Get stats
curl https://comic-api.onrender.com/api/stats

# Get comics
curl https://comic-api.onrender.com/api/comics

# Search
curl https://comic-api.onrender.com/api/search?q=one%20punch
```

### Step 5: Keep API Awake (Optional)

Karena Render akan sleep setelah 15 menit, kamu bisa:

**Option A: Use Uptime Robot (Gratis)**
1. Go to https://uptimerobot.com
2. Create monitor
3. URL: `https://comic-api.onrender.com/api/stats`
4. Interval: 5 minutes
5. Ini akan ping API setiap 5 menit, mencegah sleep

**Option B: Cron Job (di tempat lain)**
```bash
# Add ke crontab (di laptop/VPS lain)
*/5 * * * * curl https://comic-api.onrender.com/api/stats
```

**Option C: Biarkan sleep**
- API akan wake up saat ada request (butuh 30-60 detik)
- Biasanya cukup untuk personal use

## Architecture

```
┌─────────────────┐
│ GitHub Actions │ (Scraper - Scheduled)
│  (Free)         │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ MongoDB Atlas  │ (Database - Free)
│  (Free)         │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│    Render       │ (API Server - Free)
│  (Free)         │
└─────────────────┘
```

## Cost Summary

| Service | Free Tier | Cost |
|---------|-----------|------|
| GitHub Actions | 2000 min/month | $0 |
| MongoDB Atlas | 512MB storage | $0 |
| Render | Web service (sleeps) | $0 |
| **Total** | | **$0** |

## Advantages

- **100% Gratis** - Tidak ada biaya sama sekali
- **No Credit Card** - Render free tier tidak butuh CC
- **Auto-scaling** - Render scale up/down otomatis
- **Easy Setup** - Connect GitHub, auto-deploy
- **SSL Included** - HTTPS otomatis

## Disadvantages

- **API Sleep** - Render akan sleep setelah 15 menit tidak aktif
- **Cold Start** - Butuh 30-60 detik untuk wake up
- **Scraper di GitHub Actions** - Tidak 24/7, hanya scheduled

## Troubleshooting

### Deployment Failed

**Build Error:**
- Check Render logs
- Verify build command
- Check dependencies

**Runtime Error:**
- Check environment variables
- Verify MongoDB connection
- Check logs in Render dashboard

### API Always Sleeping

**Solution:**
- Setup Uptime Robot (gratis)
- Atau biarkan sleep (wake up saat ada request)

### MongoDB Connection Failed

**Solution:**
- Verify MONGODB_URI
- Check IP whitelist di MongoDB Atlas (0.0.0.0/0)
- Verify database user permissions

### GitHub Actions Not Running

**Solution:**
- Check workflow syntax
- Verify secrets
- Check Actions tab for errors

## Comparison with Other Options

| Platform | API | Scraper | 24/7 | No CC | Cost |
|----------|-----|---------|------|-------|------|
| Render + GitHub Actions | Yes (sleeps) | Yes (scheduled) | No | Yes | $0 |
| Vercel + GitHub Actions | Yes | Yes (scheduled) | No | Yes | $0 |
| Oracle Cloud | Yes | Yes | Yes | No | $0 |
| Local Machine | Yes | Yes | Yes (if PC on) | Yes | $0 |

## Next Steps

1. Setup MongoDB Atlas
2. Deploy scraper ke GitHub Actions
3. Deploy API ke Render
4. Test API endpoints
5. (Opsional) Setup Uptime Robot
6. (Opsional) Build frontend

## Conclusion

Render + GitHub Actions adalah kombinasi terbaik untuk:
- Gratis 100% tanpa credit card
- Scraper otomatis (scheduled)
- API on-demand (tidak perlu 24/7)
- Mudah setup dan maintenance

Satu-satunya kekurangan adalah API akan sleep, tapi ini bisa diatasi dengan Uptime Robot atau diterima untuk personal use.

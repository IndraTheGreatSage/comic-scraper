# GitHub Actions Deployment (100% Gratis, No Credit Card)

GitHub Actions adalah opsi terbaik untuk scraping tanpa biaya dan tanpa credit card.

## Keuntungan

- **100% Gratis** - 2000 minutes/month untuk public repo
- **No Credit Card Required** - Cukup GitHub account
- **Easy Setup** - Hanya perlu GitHub repo
- **Scheduled Runs** - Setup cron job untuk scraping rutin
- **MongoDB Atlas Free** - Bisa connect ke MongoDB Atlas gratis

## Limitations

- **Not 24/7** - Hanya jalan sesuai schedule
- **Time Limit** - 2000 minutes/month (public repo)
- **Execution Time** - Setiap run max 6 jam

## Setup Guide

### Step 1: Create GitHub Repository

```bash
cd comic-scraper
git init
git add .
git commit -m "Initial commit"
git remote add origin https://github.com/IndraTheGreatSage/comic-scraper.git
git push -u origin main
```

### Step 2: Setup MongoDB Atlas

1. Go to https://www.mongodb.com/cloud/atlas
2. Create free account (no credit card required)
3. Create cluster (Free M0 tier)
4. Create database user
5. Network Access → Add IP: `0.0.0.0/0`
6. Get connection string

### Step 3: Create GitHub Secrets

1. Di GitHub repo, go to **Settings → Secrets and variables → Actions**
2. Add secrets:

**MONGODB_URI:**
```
mongodb+srv://username:password@cluster.mongodb.net/comic_db?retryWrites=true&w=majority
```

**BASE_URL:**
```
https://komiku.id/
```

### Step 4: Create GitHub Actions Workflow

Create file: `.github/workflows/scrape.yml`

```yaml
name: Comic Scraper

on:
  schedule:
    # Run every day at 2 AM UTC
    - cron: '0 2 * * *'
  workflow_dispatch: # Allow manual trigger

jobs:
  scrape:
    runs-on: ubuntu-latest
    
    steps:
      - name: Checkout code
        uses: actions/checkout@v3
      
      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Install dependencies
        run: go mod download
      
      - name: Run scraper
        env:
          MONGODB_URI: ${{ secrets.MONGODB_URI }}
          BASE_URL: ${{ secrets.BASE_URL }}
        run: |
          go run main.go scraper.go database.go models.go
      
      - name: Upload logs
        if: always()
        uses: actions/upload-artifact@v3
        with:
          name: scraper-logs
          path: |
            *.log
```

### Step 5: Push to GitHub

```bash
git add .github/workflows/scrape.yml
git commit -m "Add GitHub Actions workflow"
git push
```

### Step 6: Test Workflow

1. Go to GitHub repo → Actions tab
2. Select "Comic Scraper" workflow
3. Click "Run workflow" → "Run workflow"
4. Monitor progress

### Step 7: Schedule Setup

Workflow sudah di-setup untuk jalan setiap hari jam 2 AM UTC. Untuk ubah schedule:

Edit `.github/workflows/scrape.yml`:
```yaml
schedule:
  # Setiap 6 jam sekali
  - cron: '0 */6 * * *'
  
  # Atau setiap hari jam 8 AM WIB (UTC+7)
  - cron: '0 1 * * *'
```

## Advanced Setup

### Multiple Scrapers

Jika mau scrape berbagai jenis komik:

```yaml
name: Comic Scraper

on:
  schedule:
    - cron: '0 2 * * *'  # Manga
    - cron: '0 4 * * *'  # Manhwa
    - cron: '0 6 * * *'  # Manhua
  workflow_dispatch:

jobs:
  scrape-manga:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      - run: go mod download
      - env:
          MONGODB_URI: ${{ secrets.MONGODB_URI }}
          BASE_URL: ${{ secrets.BASE_URL }}
        run: go run main.go scraper.go database.go models.go --type=manga
  
  scrape-manhwa:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      - run: go mod download
      - env:
          MONGODB_URI: ${{ secrets.MONGODB_URI }}
          BASE_URL: ${{ secrets.BASE_URL }}
        run: go run main.go scraper.go database.go models.go --type=manhwa
```

### Add API Server

Untuk serve API data, kamu bisa deploy ke Vercel (gratis) dan connect ke MongoDB Atlas yang sama:

1. Deploy API ke Vercel (lihat Node.js version)
2. Scraper jalan di GitHub Actions
3. Keduanya connect ke MongoDB Atlas yang sama

### Monitoring

Add notification ke Discord/Slack:

```yaml
- name: Notify on success
  if: success()
  uses: 8398a7/action-slack@v3
  with:
    status: custom
    custom_payload: |
      {
        text: '✅ Comic scraper completed successfully'
      }
  env:
    SLACK_WEBHOOK_URL: ${{ secrets.SLACK_WEBHOOK }}

- name: Notify on failure
  if: failure()
  uses: 8398a7/action-slack@v3
  with:
    status: custom
    custom_payload: |
      {
        text: '❌ Comic scraper failed'
      }
  env:
    SLACK_WEBHOOK_URL: ${{ secrets.SLACK_WEBHOOK }}
```

## Cost Calculation

**GitHub Actions:**
- Free tier: 2000 minutes/month
- Scraping 1 jam/hari = 30 hours/month = 1800 minutes
- **Masih dalam free tier!**

**MongoDB Atlas:**
- Free tier: 512MB storage
- Cukup untuk ribuan komik
- **Gratis selamanya**

**Total: $0**

## Troubleshooting

### Workflow tidak jalan
- Check schedule timezone (UTC)
- Pastikan cron syntax benar
- Cek Actions tab untuk error logs

### MongoDB connection failed
- Verify MONGODB_URI secret
- Check IP whitelist di MongoDB Atlas
- Pastikan database user permissions

### Time limit exceeded
- Kurangi jumlah page per run
- Split jadi multiple workflows
- Upgrade ke GitHub Pro ($4/month untuk 10,000 minutes)

### Rate limiting dari source website
- Tambah delay di scraper
- Kurangi concurrency
- Spread scraping ke multiple schedules

## Comparison with Other Options

| Platform | Free Tier | 24/7 | No CC | Setup Difficulty |
|----------|-----------|------|-------|------------------|
| GitHub Actions | 2000 min/month | No | Yes | Easy |
| Local Machine | Unlimited | Yes (if PC on) | Yes | Easy |
| Oracle Cloud | 2 VM ARM | Yes | No | Medium |
| Railway | $5 credit/month | Yes | No | Easy |
| Render | Web service (sleeps) | No | No | Easy |

## Next Steps

1. Setup MongoDB Atlas
2. Create GitHub repo dan push code
3. Add secrets ke GitHub
4. Create workflow file
5. Test manual trigger
6. Setup schedule
7. (Opsional) Deploy API ke Vercel
8. (Opsional) Setup monitoring

## Conclusion

GitHub Actions adalah **solusi terbaik** untuk scraping gratis tanpa credit card. Walaupun tidak 24/7, scheduled scraping biasanya cukup untuk kebanyakan use case.

Untuk 24/7 scraping tanpa credit card, opsi terbaik adalah **local machine** (jalan di laptop/PC kamu sendiri).

# Local Machine Setup (100% Gratis, No Cloud Required)

Jalan scraper di laptop/PC kamu sendiri. Ini opsi paling gratis dan paling mudah.

## Keuntungan

- **100% Gratis** - Tidak butuh cloud sama sekali
- **No Credit Card** - Tidak butuh payment method apapun
- **Full Control** - Full access ke system
- **Unlimited Scraping** - Tidak ada time limit
- **24/7** - Jika PC selalu on

## Kekurangan

- Butuh PC selalu on untuk 24/7 scraping
- Butuh setup manual
- Butuh backup manual

## Setup Guide

### Step 1: Install Prerequisites

#### Install MongoDB

**Windows:**
1. Download dari https://www.mongodb.com/try/download/community
2. Install MongoDB Community Server
3. Install MongoDB Compass (opsional, untuk GUI)
4. Start MongoDB service:
```cmd
net start MongoDB
```

**macOS:**
```bash
brew tap mongodb/brew
brew install mongodb-community
brew services start mongodb-community
```

**Linux:**
```bash
sudo apt-get install mongodb
sudo systemctl start mongodb
sudo systemctl enable mongodb
```

#### Install Go

**Windows:**
1. Download dari https://go.dev/dl/
2. Install Go
3. Restart terminal/command prompt

**macOS:**
```bash
brew install go
```

**Linux:**
```bash
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### Step 2: Setup Project

```bash
cd comic-scraper
go mod download
```

### Step 3: Configure Environment

Create `.env` file:
```env
MONGODB_URI=mongodb://localhost:27017
DB_NAME=comic_db
BASE_URL=https://komiku.id/
PORT=8080
```

### Step 4: Start MongoDB

**Windows:**
```cmd
net start MongoDB
```

**macOS/Linux:**
```bash
sudo systemctl start mongodb
# atau
brew services start mongodb-community
```

### Step 5: Run Scraper

```bash
go run main.go scraper.go database.go models.go
```

### Step 6: Run API Server (Optional)

Buka terminal baru:
```bash
go run server.go api.go database.go models.go
```

API akan jalan di http://localhost:8080

## Scheduled Scraping

### Windows (Task Scheduler)

1. Open Task Scheduler
2. Create Basic Task
3. Name: "Comic Scraper"
4. Trigger: Daily at 2 AM
5. Action: Start a program
6. Program: `C:\Program Files\Go\bin\go.exe`
7. Arguments: `run C:\path\to\comic-scraper\main.go scraper.go database.go models.go`
8. Start in: `C:\path\to\comic-scraper\`

### Linux/macOS (Cron)

```bash
crontab -e

# Add line untuk scraping setiap hari jam 2 AM
0 2 * * * cd /path/to/comic-scraper && /usr/local/go/bin/go run main.go scraper.go database.go models.go >> /var/log/comic-scraper.log 2>&1
```

## Auto-Start on Boot

### Windows (Startup Folder)

1. Create batch file `start-scraper.bat`:
```batch
@echo off
cd C:\path\to\comic-scraper
go run main.go scraper.go database.go models.go
```

2. Copy ke: `shell:startup`

### Linux (Systemd Service)

```bash
sudo nano /etc/systemd/system/comic-scraper.service
```

Paste:
```ini
[Unit]
Description=Comic Scraper Service
After=network.target mongodb.service

[Service]
Type=simple
User=YOUR_USERNAME
WorkingDirectory=/path/to/comic-scraper
ExecStart=/usr/local/go/bin/go run main.go scraper.go database.go models.go
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

Enable service:
```bash
sudo systemctl daemon-reload
sudo systemctl enable comic-scraper
sudo systemctl start comic-scraper
```

## Backup

### Manual Backup

```bash
# Backup
mongodump --db=comic_db --out=./backup/

# Restore
mongorestore --db=comic_db ./backup/comic_db
```

### Automated Backup (Cron)

```bash
crontab -e

# Backup setiap hari jam 3 AM
0 3 * * * mongodump --db=comic_db --out=./backup/$(date +\%Y\%m\%d)/
```

### Backup ke Cloud (Google Drive/Dropbox)

Install rclone:
```bash
# Linux/macOS
curl https://rclone.org/install.sh | sudo bash

# Windows: Download dari rclone.org
```

Setup rclone:
```bash
rclone config
# Follow prompts untuk connect ke Google Drive/Dropbox
```

Backup script:
```bash
#!/bin/bash
# backup-to-cloud.sh

# Backup MongoDB
mongodump --db=comic_db --out=./backup/

# Upload ke cloud
rclone copy ./backup/ gdrive:comic-backup/$(date +%Y%m%d)/

# Cleanup old backups (keep 7 days)
find ./backup/ -type d -mtime +7 -exec rm -rf {} \;
```

Add ke cron:
```bash
crontab -e
0 3 * * * /path/to/backup-to-cloud.sh
```

## Monitoring

### Simple Log Monitoring

```bash
# Watch logs in real-time
tail -f scraper.log

# Atau jika menggunakan systemd
sudo journalctl -u comic-scraper -f
```

### Email Notification (Linux)

Install mailutils:
```bash
sudo apt-get install mailutils
```

Add ke scraper script untuk kirim email setelah selesai:
```bash
echo "Comic scraping completed at $(date)" | mail -s "Scraper Report" your@email.com
```

## Access from Other Devices

### Local Network

Jika mau akses API dari device lain di network yang sama:

1. Pastikan firewall allow port 8080
2. Cari IP lokal: `ipconfig` (Windows) atau `ifconfig` (Linux/macOS)
3. Akses dari device lain: `http://YOUR_LOCAL_IP:8080`

### ngrok (Public Access)

Untuk akses dari luar network:

```bash
# Install ngrok
# Windows: Download dari ngrok.com
# Linux/macOS: curl -s https://ngrok-agent.s3.amazonaws.com/ngrok.asc | sudo tee /etc/apt/trusted.gpg.d/ngrok.asc
#          echo "deb https://ngrok-agent.s3.amazonaws.com buster main" | sudo tee /etc/apt/sources.list.d/ngrok.list
#          sudo apt update && sudo apt install ngrok

# Expose port 8080
ngrok http 8080
```

Akses via URL yang diberikan ngrok.

## Performance Optimization

### Increase Scraping Speed

Edit `scraper.go`:
```go
c.Limit(&colly.LimitRule{
    Parallelism: 4,  // Increase from 2 to 4
    Delay: 1 * time.Second,  // Decrease from 2s to 1s
})
```

**Warning:** Terlalu cepat bisa kena block dari source website.

### Reduce Memory Usage

Jika PC RAM kecil, kurangi concurrency di scraper.

## Troubleshooting

### MongoDB tidak jalan

**Windows:**
```cmd
net start MongoDB
```

**Linux/macOS:**
```bash
sudo systemctl start mongodb
```

### Port 8080 sudah dipakai

Ganti port di `.env`:
```env
PORT=8081
```

### Scraper crash

Check logs:
```bash
tail -f scraper.log
```

Atau jika menggunakan systemd:
```bash
sudo journalctl -u comic-scraper -n 50
```

### PC sleep/hibernate

Disable sleep/hibernate:
- Windows: Power settings → Never sleep
- Linux: `systemctl mask sleep.target suspend.target hibernate.target hybrid-sleep.target`

## Cost: $0

Tidak ada biaya sama sekali. Semua jalan di hardware kamu sendiri.

## Comparison

| Aspect | Local | GitHub Actions | Oracle Cloud |
|--------|--------|----------------|--------------|
| Cost | $0 | $0 | $0 (butuh CC) |
| 24/7 | Yes (if PC on) | No | Yes |
| Setup | Easy | Easy | Medium |
| Maintenance | Manual | Auto | Auto |
| Control | Full | Limited | Full |

## Recommendation

**Pilih Local jika:**
- Punya PC yang bisa selalu on
- Mau full control
- Tidak mau setup cloud

**Pilih GitHub Actions jika:**
- PC tidak bisa selalu on
- Mau auto-schedule
- Tidak mau maintenance

**Pilih Oracle Cloud jika:**
- Punya credit card (untuk verifikasi saja)
- Mau 24/7 tanpa PC selalu on
- Mau auto-maintenance

# Upload ke GitHub - Step by Step

## Step 1: Create GitHub Repository

1. Go to https://github.com
2. Login (atau sign up jika belum punya)
3. Click **+** di pojok kanan atas → **New repository**
4. Isi:
   - **Repository name:** `comic-scraper` (atau nama lain)
   - **Description:** Comic scraper with MongoDB integration
   - **Public/Private:** Pilih **Public** (untuk GitHub Actions gratis)
5. Click **Create repository**

## Step 2: Initialize Git di Local

```bash
cd c:\Users\CHIIO\Downloads\comic-scraper
git init
```

## Step 3: Add Files ke Git

```bash
git add .
```

## Step 4: Commit Changes

```bash
git commit -m "Initial commit - Comic scraper with MongoDB"
```

## Step 5: Connect ke GitHub Repository

```bash
git remote add origin https://github.com/IndraTheGreatSage/comic-scraper.git
```

**Username GitHub kamu: IndraTheGreatSage**

## Step 6: Push ke GitHub

```bash
git branch -M main
git push -u origin main
```

## Step 7: Setup GitHub Secrets

Setelah push berhasil:

1. Buka repository di GitHub
2. Click **Settings** tab
3. Di sidebar, klik **Secrets and variables** → **Actions**
4. Klik **New repository secret**
5. Add secrets:

**Secret 1: MONGODB_URI**
- Name: `MONGODB_URI`
- Value: Connection string dari MongoDB Atlas
  ```
  mongodb+srv://username:password@cluster.mongodb.net/comic_db?retryWrites=true&w=majority
  ```
- Click **Add secret**

**Secret 2: BASE_URL**
- Name: `BASE_URL`
- Value: `https://komiku.id/`
- Click **Add secret**

## Step 8: Verifikasi Workflow

1. Di GitHub repo, klik **Actions** tab
2. Kamu akan melihat workflow "Comic Scraper"
3. Workflow akan otomatis jalan sesuai schedule (setiap hari jam 2 AM UTC)
4. Untuk test manual, klik workflow → **Run workflow** → **Run workflow**

## Troubleshooting

### Git push failed

**Error: "remote origin already exists"**
```bash
git remote remove origin
git remote add origin https://github.com/YOUR_USERNAME/comic-scraper.git
git push -u origin main
```

**Error: "Authentication failed"**
- Setup GitHub personal access token:
  1. GitHub Settings → Developer settings → Personal access tokens → Tokens (classic)
  2. Generate new token
  3. Select scopes: `repo`
  4. Copy token
- Gunakan token saat push:
```bash
git push https://TOKEN@github.com/YOUR_USERNAME/comic-scraper.git
```

### Workflow tidak muncul di Actions

Pastikan file `.github/workflows/scrape.yml` sudah di-commit dan push:
```bash
git add .github/workflows/scrape.yml
git commit -m "Add GitHub Actions workflow"
git push
```

### MongoDB connection failed

- Verify MONGODB_URI secret di GitHub
- Pastikan IP whitelist di MongoDB Atlas: `0.0.0.0/0`
- Pastikan database user punya permission yang benar

## Next Steps

Setelah berhasil upload ke GitHub:

1. **Test GitHub Actions** - Run manual trigger dari Actions tab
2. **Setup MongoDB Atlas** - Jika belum punya
3. **Deploy ke Render** - Untuk API server
4. **Test API** - Verify API berjalan dengan data dari MongoDB

## Quick Summary

```bash
# 1. Create repo di GitHub (via browser)

# 2. Initialize git
cd c:\Users\CHIIO\Downloads\comic-scraper
git init

# 3. Add dan commit
git add .
git commit -m "Initial commit"

# 4. Connect ke GitHub
git remote add origin https://github.com/IndraTheGreatSage/comic-scraper.git

# 5. Push
git branch -M main
git push -u origin main

# 6. Setup secrets di GitHub Settings
# - MONGODB_URI
# - BASE_URL
```

**Total time: ~5-10 menit**

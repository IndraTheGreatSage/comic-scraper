# Upload ke GitHub via GUI (Tanpa Command Line)

## Opsi 1: GitHub Desktop (Paling Mudah)

### Step 1: Install GitHub Desktop
1. Download dari https://desktop.github.com/
2. Install di Windows
3. Login dengan akun GitHub (IndraTheGreatSage)

### Step 2: Clone Repository
1. Buka GitHub Desktop
2. Click **File** → **Clone Repository**
3. Pilih tab **URL**
4. Repository URL: `https://github.com/IndraTheGreatSage/comic-scraper.git`
5. Local path: `C:\Users\CHIIO\Downloads\comic-scraper`
6. Click **Clone**

### Step 3: Add Files
1. Di GitHub Desktop, kamu akan melihat semua file yang belum di-commit
2. Tulis commit message: "Initial commit - Comic scraper with MongoDB"
3. Click **Commit to main**

### Step 4: Push ke GitHub
1. Click **Push origin** di pojok kanan atas
2. Tunggu sampai selesai

---

## Opsi 2: Upload via Web Browser (Drag & Drop)

### Step 1: Create Repository di GitHub
1. Buka https://github.com/IndraTheGreatSage/comic-scraper
2. Kalau belum ada, create new repository dulu

### Step 2: Upload Files
1. Di repository GitHub, click **uploading an existing file**
2. Drag semua file dari `C:\Users\CHIIO\Downloads\comic-scraper` ke browser
3. Atau click **choose your files**
4. Select semua file:
   - main.go
   - scraper.go
   - database.go
   - api.go
   - server.go
   - models.go
   - go.mod
   - .env
   - render.yaml
   - .github/workflows/scrape.yml
   - Documentation files (.md)
5. Tulis commit message: "Initial commit - Comic scraper with MongoDB"
6. Click **Commit changes**

---

## Opsi 3: VS Code dengan GitHub Integration

### Step 1: Install VS Code
Kalau belum punya, download dari https://code.visualstudio.com/

### Step 2: Install GitHub Extension
1. Buka VS Code
2. Extensions (Ctrl+Shift+X)
3. Search: "GitHub Pull Requests"
4. Install

### Step 3: Open Project
1. File → Open Folder
2. Select: `C:\Users\CHIIO\Downloads\comic-scraper`

### Step 4: Commit dan Push
1. Di sidebar kiri, klik icon **Source Control** (atau Ctrl+Shift+G)
2. Kamu akan melihat semua file yang berubah
3. Tulis message: "Initial commit - Comic scraper with MongoDB"
4. Click **Commit**
5. Klik **Sync Changes** (icon refresh di pojok kanan bawah)

---

## Rekomendasi Saya

**Paling Mudah:** GitHub Desktop
- GUI yang jelas
- Tidak perlu command line
- Auto-sync dengan GitHub

**Paling Cepat:** Upload via Web Browser
- Tidak perlu install apa-apa
- Drag & drop langsung
- Tapi harus upload file per file (tidak bisa folder sekaligus)

**Paling Fleksibel:** VS Code
- Bisa edit code langsung
- Git integration built-in
- Tapi perlu install VS Code

---

## Setelah Upload Berhasil

Setup GitHub Secrets:
1. Buka https://github.com/IndraTheGreatSage/comic-scraper
2. **Settings** → **Secrets and variables** → **Actions**
3. **New repository secret**:
   - Name: `MONGODB_URI`
   - Value: (connection string MongoDB Atlas)
   - Name: `BASE_URL`
   - Value: `https://komiku.id/`

Mau coba opsi yang mana?

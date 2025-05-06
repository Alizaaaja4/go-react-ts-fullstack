# 🚀 Fullstack Web Development - Go (Gin) + React + TypeScript

Panduan langkah demi langkah untuk membangun aplikasi fullstack menggunakan Go (Gin) sebagai backend dan React + TypeScript sebagai frontend, berdasarkan tutorial SantriKoding.

---

## 🧠 Backend – Go (Gin + GORM + Air + DotEnv)

### 1. Inisialisasi Modul Go
```bash
go mod init santrikoding/backend-api
```

### 2. Install Gin Web Framework
```bash
go get -u github.com/gin-gonic/gin@v1.10.0
```

### 3. Install Air (Live Reload untuk Go)
```bash
go install github.com/air-verse/air@latest
air -v
air init
air  # Jalankan server dengan hot reload
```

### 4. Install dotenv support
```bash
go get github.com/joho/godotenv@v1.5.1
```

### 5. Install GORM dan MySQL Driver
```bash
go get -u gorm.io/gorm@v1.25.12
go get -u gorm.io/driver/mysql@v1.5.7
```

---

## ⚛️ Frontend – React + TypeScript + Vite

### 1. Inisialisasi Project React dengan Vite + TypeScript
```bash
npm create vite@5.2.3 frontend-react-ts -- --template react-ts
cd frontend-react-ts
npm install
npm run dev
```

### 2. Install Axios
```bash
npm install axios@1.8.4
```

### 3. Install js-cookie dan type-nya
```bash
npm install js-cookie@3.0.5
npm install -D @types/js-cookie@3.0.6
```

### 4. Install React Router
```bash
npm install react-router@7.4.0
```

### 5. Install TanStack React Query
```bash
npm install @tanstack/react-query@5.72.1
```

---

## 🧱 Struktur Folder Disarankan (Frontend)

```bash
frontend-react-ts/
├── src/
│   ├── assets/
│   ├── components/
│   ├── pages/
│   ├── services/
│   ├── App.tsx
│   └── main.tsx
├── public/
├── vite.config.ts
├── package.json
└── tsconfig.json
```

---

## 📝 Catatan

- `go mod init santrikoding/backend-api` adalah nama modul, bukan nama folder. Boleh diganti sesuai nama repo GitHub atau organisasi kamu.
- Untuk publikasi ke GitHub, disarankan menggunakan:
  ```bash
  go mod init github.com/username/repo-name
  ```
- Gunakan `air` untuk hot reload agar development lebih efisien.
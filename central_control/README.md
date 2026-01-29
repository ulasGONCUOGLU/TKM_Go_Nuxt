# Central Control

Minimal MVP mimarisi iskeleti. HTTP API üzerinden “plan” gönderip bir controller’ın (şu an mock) durumunu yönetir ve sorgular.

---

## Ne İşe Yarıyor?

- **Plan uygulama:** `/plans/apply` endpoint’ine bir plan (faz listesi) gönderilir. Her fazın süresi (saniye) kadar bekleme simüle edilir.
- **Durum sorgulama:** `/controllers/mock/status` ile o anki **aktif faz** ve **kalan süre (saniye)** alınır.
- **Mock controller:** Gerçek donanım yerine bellekte çalışan bir simülasyon kullanılır; planlar arka planda goroutine ile işlenir.

### API Endpoint’leri

| Metod | Endpoint | Açıklama |
|-------|----------|----------|
| `POST` | `/plans/apply` | Plan uygula (JSON body) |
| `GET` | `/controllers/mock/status` | Mock controller durumunu getir |

### Plan JSON Örneği

```json
{
  "id": "plan-1",
  "phases": [
    { "id": 1, "duration": 5 },
    { "id": 2, "duration": 10 }
  ]
}
```

- `id`: Plan tanımlayıcısı  
- `phases`: Sırayla çalışacak fazlar  
  - `id`: Faz numarası  
  - `duration`: Süre (saniye)

---

## Projeyi Çalıştırma (Adım Adım)

### 1. Gereksinimler

- **Go 1.20+** yüklü olmalı. Kontrol için:

  ```bash
  go version
  ```

### 2. Proje Dizinine Geç

  ```bash
  cd c:\Users\ulasg\Desktop\py\TKM2\central_control
  ```

### 3. Bağımlılıkları Kontrol Et

  Bu projede harici bağımlılık yok; `go mod tidy` yeterli:

  ```bash
  go mod tidy
  ```

### 4. Sunucuyu Başlat

  ```bash
  go run ./cmd/server
  ```

  Çıktıda şunu görmelisin:

  ```
  Server running on :8080
  ```

  Sunucu **http://localhost:8080** üzerinde dinler.

### 5. API’yi Test Et

**a) Durum sorgula (başlangıçta):**

  ```bash
  curl http://localhost:8080/controllers/mock/status
  ```

  Örnek cevap: `{"active_phase":0,"remaining_sec":0}`

**b) Plan uygula:**

  ```bash
  curl -X POST http://localhost:8080/plans/apply -H "Content-Type: application/json" -d "{\"id\":\"plan-1\",\"phases\":[{\"id\":1,\"duration\":5},{\"id\":2,\"duration\":10}]}"
  ```

  Beklenen cevap: `{"ok":true}`

**c) Tekrar durum sorgula:**

  ```bash
  curl http://localhost:8080/controllers/mock/status
  ```

  Plan çalışırken `active_phase` ve `remaining_sec` değişir; plan bittikten sonra son fazın değerleri kalır.

### 6. Durdurmak

  Terminalde sunucu çalışan pencerede **Ctrl+C** ile sunucuyu durdurabilirsin.

---

## GitHub'a Yükleme

1. **GitHub'da yeni repo oluştur:** [github.com/new](https://github.com/new)
   - Repository adı: `TKM_Go_Nuxt` (veya istediğin isim)
   - **"Add a README file"** işaretleme (zaten var)
   - **"Add .gitignore"** işaretleme (zaten var)
   - Create repository

2. **Remote ekle ve push et:**
   ```bash
   git remote add origin https://github.com/ulasGONCUOGLU//TKM_Go_Nuxt.git
   git branch -M main
   git push -u origin main
   ```

---

## Proje Yapısı

```
central_control/
├── cmd/server/          # Giriş noktası (main)
├── internal/
│   ├── api/             # HTTP handler’lar, route kayıt
│   ├── controller/      # ControllerManager (plan uygulama, durum)
│   ├── model/           # Plan, Phase, ControllerStatus
│   └── adapter/
│       ├── adapter.go   # ControllerAdapter arayüzü
│       └── mock/        # Mock controller implementasyonu
├── go.mod
└── README.md
```

---

## Özet

1. `go run ./cmd/server` ile sunucuyu başlat.  
2. `POST /plans/apply` ile plan gönder.  
3. `GET /controllers/mock/status` ile anlık faz ve kalan süreyi oku.

İleride `adapter` paketinde gerçek bir donanım adapter’ı yazılıp mock yerine kullanılabilir; API ve controller katmanı aynı kalır.

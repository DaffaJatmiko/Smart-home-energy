# Final Project - Artificial Intelligence menggunakan Golang

## AI-Powered Smart Home Energy Management System

### Deskripsi

Proyek ini mengembangkan Sistem Manajemen Energi Rumah Pintar menggunakan Golang dan model AI Tapas dari Huggingface Model Hub. Sistem ini bertujuan untuk memprediksi dan mengelola penggunaan energi dalam lingkungan rumah pintar. Aplikasi ini menerima data tentang penggunaan energi rumah dan memberikan wawasan serta rekomendasi tentang cara mengoptimalkan konsumsi energi.

### Fitur

1. **Prediksi Konsumsi Energi**: Sistem memprediksi konsumsi energi rumah berdasarkan data historis.
2. **Rekomendasi Penghematan Energi**: Sistem memberikan rekomendasi tentang cara menghemat energi berdasarkan konsumsi energi yang diprediksi.

### Data Input

Data input berupa file CSV dengan kolom berikut:

- `Date`: Tanggal data penggunaan energi.
- `Time`: Waktu data penggunaan energi.
- `Appliance`: Nama alat.
- `Energy_Consumption`: Konsumsi energi alat dalam kWh.
- `Room`: Ruang tempat alat berada.
- `Status`: Status alat (On/Off).

Contoh data:

- Date,Time,Appliance,Energy_Consumption,Room,Status
- 2022-01-01,00:00,Refrigerator,1.2,Kitchen,On
- 2022-01-01,01:00,Refrigerator,1.2,Kitchen,On
- 2022-01-01,08:00,TV,0.8,Living Room,Off
- 2022-01-01,09:00,TV,0.8,Living Room,On
- 2022-01-01,10:00,TV,0.8,Living Room,On

### Penggunaan Model AI

Menggunakan model AI Tapas `tapas-base-finetuned-wtq` dari Hugging Face Model Hub untuk memahami data tabel dan membuat prediksi tentang konsumsi energi masa depan. Model ini menerima file CSV sebagai input dan menghasilkan prediksi untuk total konsumsi energi hari berikutnya.

### Interface

Aplikasi memiliki interface yang dapat berupa aplikasi CLI maupun Web Application. Interface dikembangkan agar mirip dengan chatbot di mana pengguna bisa bertanya mengenai data-data yang ada di file input.

### Dependencies

- Golang
- Gin-Gonic
- Huggingface API

### Cara Menjalankan Aplikasi

1. Clone repositori ini.

   ```bash
   git clone <repository_url>

   ```

2. Pindah ke direktori proyek:

   ```bash
   cd smart_home_energy

   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Pastikan Anda memiliki token Huggingface API yang disimpan dalam file .env dengan format:

```bash
 HUGGINGFACE_TOKEN=your_huggingface_token
```

5. Jalankan aplikasi:

   ```bash
   go run cmd/main.go
   ```

6. Akses aplikasi melalui browser di http://localhost:8080.

## Struktur Proyek

```bash
.
├── cmd
│   └── main.go
├── data-series.csv
├── go.mod
├── go.sum
├── handlers
│   └── api.go
├── model
│   ├── inputs.go
│   └── response.go
├── repository
│   └── ai_model_connect.go
├── service
│   └── ai_service.go
├── util
│   ├── csv_reader.go
│   └── env_loader.go
└── view
    └── index.html
```

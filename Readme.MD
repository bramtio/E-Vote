# 🗳️ E-VOTING APP

Aplikasi E-Voting berbasis Command Line Interface (CLI) yang dikembangkan menggunakan bahasa Go (Golang). Aplikasi ini dirancang untuk membantu proses pemungutan suara secara digital dengan fitur pengelolaan kandidat, voting, pencarian data, pengurutan data, dan statistik hasil pemilihan secara real-time.

## 📖 Deskripsi Program

E-Voting App merupakan aplikasi sederhana yang digunakan untuk melakukan proses pemilihan secara elektronik. Sistem memungkinkan pengguna mengelola data kandidat, melakukan voting, mencari kandidat, mengurutkan data kandidat, dan menampilkan statistik hasil pemilihan.

Aplikasi ini dibuat sebagai implementasi konsep Algoritma dan Pemrograman menggunakan bahasa Go dengan memanfaatkan struktur data array dan struct, serta algoritma sorting dan searching.

---

## 🎯 Tujuan Program

- Mengimplementasikan konsep Struct dan Array dalam bahasa Go.
- Mengimplementasikan operasi CRUD (Create, Read, Update, Delete).
- Mengimplementasikan algoritma Insertion Sort dan Selection Sort.
- Mengimplementasikan algoritma Binary Search.
- Membuat sistem pemungutan suara digital berbasis CLI.
- Menampilkan statistik hasil pemilihan secara real-time.

---

## ✨ Fitur Aplikasi

### 1. Tambah Kandidat
Menambahkan kandidat baru ke dalam sistem.

Data yang diinput:
- Nomor urut
- Nama kandidat
- Visi dan misi

### 2. Tampilkan Kandidat
Menampilkan seluruh kandidat yang tersimpan.

Pilihan pengurutan:
- Berdasarkan nomor urut (Insertion Sort)
- Berdasarkan jumlah suara (Selection Sort)

### 3. Ubah Kandidat
Mengubah informasi kandidat yang sudah ada.

### 4. Hapus Kandidat
Menghapus data kandidat berdasarkan nomor urut.

### 5. Voting
Memberikan suara kepada kandidat tertentu berdasarkan nomor urut.

### 6. Search Kandidat

Metode pencarian:
- Binary Search berdasarkan nomor urut
- Kandidat dengan suara terbanyak

### 7. Statistik

Menampilkan:
- Total suara
- Jumlah suara setiap kandidat
- Persentase suara setiap kandidat

---

## 🗂️ Struktur Data

```go
type Kandidat struct {
    Nomor    int
    Nama     string
    VisiMisi string
    Suara    int
}
```

| Field | Tipe | Keterangan |
|---------|---------|---------|
| Nomor | int | Nomor urut kandidat |
| Nama | string | Nama kandidat |
| VisiMisi | string | Visi dan misi kandidat |
| Suara | int | Jumlah suara kandidat |

---

## 💾 Struktur Penyimpanan Data

Program menggunakan array statis.

```go
const MAX = 100

var data [MAX]Kandidat
var n int
```

Keterangan:

- Maksimal 100 kandidat
- Penyimpanan sementara di memori
- Data akan hilang ketika program ditutup

---

## 🚀 Cara Menjalankan Program

### Clone Repository

```bash
git clone https://github.com/bramtio/E-Vote.git
```

### Masuk ke Folder Project

```bash
cd E-Vote
```

### Jalankan Program

```bash
go run main.go
```

Atau build executable:

```bash
go build
./E-Vote
```

---

## 📋 Panduan Pengguna

Saat program dijalankan akan muncul menu:

```text
====================
      E-VOTING
====================
1. Tambah Kandidat
2. Tampilkan Data
3. Ubah Data
4. Hapus Data
5. Voting
6. Search
7. Statistik
0. Keluar
```

### Menambah Kandidat

1. Pilih menu 1
2. Masukkan nomor kandidat
3. Masukkan nama kandidat
4. Masukkan visi misi

### Melakukan Voting

1. Pilih menu 5
2. Masukkan nomor kandidat
3. Sistem akan menambahkan satu suara

### Mencari Kandidat

1. Pilih menu 6
2. Pilih metode pencarian
3. Masukkan nomor kandidat atau tampilkan suara tertinggi

### Melihat Statistik

1. Pilih menu 7
2. Sistem menampilkan total suara dan persentase perolehan suara

---

## ⚙️ Algoritma yang Digunakan

### Insertion Sort

Digunakan untuk mengurutkan kandidat berdasarkan nomor urut.

Kompleksitas:

```
Best Case  : O(n)
Worst Case : O(n²)
```

### Selection Sort

Digunakan untuk mengurutkan kandidat berdasarkan jumlah suara.

Kompleksitas:

```
Best Case  : O(n²)
Worst Case : O(n²)
```

### Binary Search

Digunakan untuk mencari kandidat berdasarkan nomor urut.

Kompleksitas:

```
O(log n)
```

---

## 🏗️ Struktur Program

```text
main()
│
├── tambahKandidat()
├── tampilkanKandidat()
├── ubahKandidat()
├── hapusKandidat()
├── voting()
├── searchKandidat()
│   ├── cariNomorUrut()
│   └── cariSuaraTerbanyak()
│
├── insertionSortNomor()
├── selectionSortSuara()
└── statistik()
```

---

## 🔧 Teknologi yang Digunakan

- Go (Golang)
- Struct
- Array
- CLI (Command Line Interface)
- Insertion Sort
- Selection Sort
- Binary Search

---

## 📝 Catatan Teknis

### Kelebihan

- Program ringan dan cepat
- Struktur modular
- Mudah dipahami
- Cocok untuk pembelajaran algoritma dasar

### Keterbatasan

- Belum menggunakan database
- Data tidak tersimpan permanen
- Belum ada login pengguna
- Belum ada validasi satu pemilih satu suara
- Masih berbasis terminal

### Pengembangan Selanjutnya

- Login Admin dan Pemilih
- Penyimpanan Database
- GUI/Desktop Application
- Web Application
- Keamanan dan Enkripsi Data

---

## 👨‍💻 Tim Pengembang

**Naufal Rasyid Hakim**  
108072500167

**Bramantyo Pangestu Irawan**  
108072500055

Program Studi Informatika  
Fakultas Informatika  
Universitas Telkom

---

## 📄 Lisensi

Proyek ini dibuat untuk keperluan akademik pada mata kuliah Algoritma dan Pemrograman 2.# 🗳️ E-VOTING APP

Aplikasi E-Voting berbasis Command Line Interface (CLI) yang dikembangkan menggunakan bahasa Go (Golang). Aplikasi ini dirancang untuk membantu proses pemungutan suara secara digital dengan fitur pengelolaan kandidat, voting, pencarian data, pengurutan data, dan statistik hasil pemilihan secara real-time.

## 📖 Deskripsi Program

E-Voting App merupakan aplikasi sederhana yang digunakan untuk melakukan proses pemilihan secara elektronik. Sistem memungkinkan pengguna mengelola data kandidat, melakukan voting, mencari kandidat, mengurutkan data kandidat, dan menampilkan statistik hasil pemilihan.

Aplikasi ini dibuat sebagai implementasi konsep Algoritma dan Pemrograman menggunakan bahasa Go dengan memanfaatkan struktur data array dan struct, serta algoritma sorting dan searching.

---

## 🎯 Tujuan Program

- Mengimplementasikan konsep Struct dan Array dalam bahasa Go.
- Mengimplementasikan operasi CRUD (Create, Read, Update, Delete).
- Mengimplementasikan algoritma Insertion Sort dan Selection Sort.
- Mengimplementasikan algoritma Binary Search.
- Membuat sistem pemungutan suara digital berbasis CLI.
- Menampilkan statistik hasil pemilihan secara real-time.

---

## ✨ Fitur Aplikasi

### 1. Tambah Kandidat
Menambahkan kandidat baru ke dalam sistem.

Data yang diinput:
- Nomor urut
- Nama kandidat
- Visi dan misi

### 2. Tampilkan Kandidat
Menampilkan seluruh kandidat yang tersimpan.

Pilihan pengurutan:
- Berdasarkan nomor urut (Insertion Sort)
- Berdasarkan jumlah suara (Selection Sort)

### 3. Ubah Kandidat
Mengubah informasi kandidat yang sudah ada.

### 4. Hapus Kandidat
Menghapus data kandidat berdasarkan nomor urut.

### 5. Voting
Memberikan suara kepada kandidat tertentu berdasarkan nomor urut.

### 6. Search Kandidat

Metode pencarian:
- Binary Search berdasarkan nomor urut
- Kandidat dengan suara terbanyak

### 7. Statistik

Menampilkan:
- Total suara
- Jumlah suara setiap kandidat
- Persentase suara setiap kandidat

---

## 🗂️ Struktur Data

```go
type Kandidat struct {
    Nomor    int
    Nama     string
    VisiMisi string
    Suara    int
}
```

| Field | Tipe | Keterangan |
|---------|---------|---------|
| Nomor | int | Nomor urut kandidat |
| Nama | string | Nama kandidat |
| VisiMisi | string | Visi dan misi kandidat |
| Suara | int | Jumlah suara kandidat |

---

## 💾 Struktur Penyimpanan Data

Program menggunakan array statis.

```go
const MAX = 100

var data [MAX]Kandidat
var n int
```

Keterangan:

- Maksimal 100 kandidat
- Penyimpanan sementara di memori
- Data akan hilang ketika program ditutup

---

## 🚀 Cara Menjalankan Program

### Clone Repository

```bash
git clone https://github.com/bramtio/E-Vote.git
```

### Masuk ke Folder Project

```bash
cd E-Vote
```

### Jalankan Program

```bash
go run main.go
```

Atau build executable:

```bash
go build
./E-Vote
```

---

## 📋 Panduan Pengguna

Saat program dijalankan akan muncul menu:

```text
====================
      E-VOTING
====================
1. Tambah Kandidat
2. Tampilkan Data
3. Ubah Data
4. Hapus Data
5. Voting
6. Search
7. Statistik
0. Keluar
```

### Menambah Kandidat

1. Pilih menu 1
2. Masukkan nomor kandidat
3. Masukkan nama kandidat
4. Masukkan visi misi

### Melakukan Voting

1. Pilih menu 5
2. Masukkan nomor kandidat
3. Sistem akan menambahkan satu suara

### Mencari Kandidat

1. Pilih menu 6
2. Pilih metode pencarian
3. Masukkan nomor kandidat atau tampilkan suara tertinggi

### Melihat Statistik

1. Pilih menu 7
2. Sistem menampilkan total suara dan persentase perolehan suara

---

## ⚙️ Algoritma yang Digunakan

### Insertion Sort

Digunakan untuk mengurutkan kandidat berdasarkan nomor urut.

Kompleksitas:

```
Best Case  : O(n)
Worst Case : O(n²)
```

### Selection Sort

Digunakan untuk mengurutkan kandidat berdasarkan jumlah suara.

Kompleksitas:

```
Best Case  : O(n²)
Worst Case : O(n²)
```

### Binary Search

Digunakan untuk mencari kandidat berdasarkan nomor urut.

Kompleksitas:

```
O(log n)
```

---

## 🏗️ Struktur Program

```text
main()
│
├── tambahKandidat()
├── tampilkanKandidat()
├── ubahKandidat()
├── hapusKandidat()
├── voting()
├── searchKandidat()
│   ├── cariNomorUrut()
│   └── cariSuaraTerbanyak()
│
├── insertionSortNomor()
├── selectionSortSuara()
└── statistik()
```

---

## 🔧 Teknologi yang Digunakan

- Go (Golang)
- Struct
- Array
- CLI (Command Line Interface)
- Insertion Sort
- Selection Sort
- Binary Search

---

## 📝 Catatan Teknis

### Kelebihan

- Program ringan dan cepat
- Struktur modular
- Mudah dipahami
- Cocok untuk pembelajaran algoritma dasar

### Keterbatasan

- Belum menggunakan database
- Data tidak tersimpan permanen
- Belum ada login pengguna
- Belum ada validasi satu pemilih satu suara
- Masih berbasis terminal

### Pengembangan Selanjutnya

- Login Admin dan Pemilih
- Penyimpanan Database
- GUI/Desktop Application
- Web Application
- Keamanan dan Enkripsi Data

---

## 👨‍💻 Tim Pengembang

**Naufal Rasyid Hakim**  
108072500167

**Bramantyo Pangestu Irawan**  
108072500055

Program Studi Informatika  
Fakultas Informatika  
Universitas Telkom

---

## 📄 Lisensi

Proyek ini dibuat untuk keperluan akademik pada mata kuliah Algoritma dan Pemrograman 2.

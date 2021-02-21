# Design ERD

## Navigasi
- [Gambar](#gambar)
- [Tabel Master](#tabel-master)
  - [Users](#users)
  - [Categories](#categories)
  - [Transactions](#transactions)
- [Tabel Relasi](#tabel-relasi)
  - [Products](#products)
  - [Transaction Items](#transaction-items)

## Gambar
Gambar Design ERD

![Gambar Design ERD](erd.png)

## Tabel Master
Penjelasan Tabel Master, merupakan tabel yang tidak memuat attribute dari tabel lain. Namun Tabel Master digunakan oleh tabel lain sebagai referensi, seperti Tabel Produk yang meminjam attribute`id` dari Tabel Categories. 

### Users

Tabel Users

| Nama Field | Tipe Data | Contoh Nilai | Penjelasan |
| ---------- | --------- | ------------ | ---------- |
| id | INT |
| nama | STRING(255) |
| email | STRING(320) |
| password | STRING(40) (HASHED) |
| token | STRING(60) |
| nohp | STRING(13) |
| foto | STRING(255) (LINK) |
| alamat | STRING(255) |
| status | STRING("CUSTOMER", SUPERADMIN) |
| isVerified | BOOLEAN |

### Categories

### Transactions

## Tabel Relasi
Penjelasan Tabel Relasi, merupakan tabel yang tidak bisa berdiri sendiri karena memerlukan attribute dari tabel lain untuk dapat menambahkan data baru. Contohnya tidak bisa menambahkan data pada Tabel Products jika tidak diberikan nilai `id` dari Tabel Categories.

### Products

### Transaction Items
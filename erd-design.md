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
| id | INT | 1 | ID merupakan nilai unik dari setiap data user, digunakan pada proses transaksi untuk mengetahui transaksi yang dimiliki oleh suatu user.
| nama | STRING(255) | Rizky Alterra Academy | nama pendek ataupun nama panjang dari user digunakan untuk validasi transaksi.
| email | STRING(320) | rizky.alta@gmail.com | email dari user untuk proses login dan register. |
| password | STRING(40) | f4ebfd7a42d9a43a536e2bed9ee4974abf8f8dc8 | password digunakan saat proses login dan register |
| token | STRING(60) | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | token digunakan untuk dapat mengakses sebagian besar endpoint dan didapatkan pada saat proses login dan register. |
| no_hp | STRING(13) | 081234567890 | nomor hp milik user yang bisa digunakan saat validasi transaksi |
| foto | STRING(255) | https://unsplash.com/photos/bC0uxTH7aS0 | link foto dari user. |
| alamat | STRING(255) | Jalan Medan Merdeka Barat No. 12. | alamat dari user untuk melakukan proses pengiriman |
| status | STRING("CUSTOMER", ADMIN) | CUSTOMER | status dari user untuk membedakan hak akses saat mengakses suatu endpoint. |
| is_verified | BOOLEAN | 0 | untuk proses verifikasi sehingga user tidak dapat melakukan transaksi jika belum diverifikasi oleh admin (di luar MVP). |

### Categories

### Transactions

## Tabel Relasi
Penjelasan Tabel Relasi, merupakan tabel yang tidak bisa berdiri sendiri karena memerlukan attribute dari tabel lain untuk dapat menambahkan data baru. Contohnya tidak bisa menambahkan data pada Tabel Products jika tidak diberikan nilai `id` dari Tabel Categories.

### Products

### Transaction Items
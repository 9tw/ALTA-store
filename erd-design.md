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
| password | STRING(40) | f4ebfd7a42d9a43a536e2bed9ee4974ab | password digunakan saat proses login dan register |
| token | STRING(60) | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ | token digunakan untuk dapat mengakses sebagian besar endpoint dan didapatkan pada saat proses login dan register. |
| no_hp | STRING(13) | 081234567890 | nomor hp milik user yang bisa digunakan saat validasi transaksi |
| foto | STRING(255) | https://unsplash.com/photos/bC0uxTH7aS0 | link foto dari user. |
| alamat | STRING(255) | Jalan Medan Merdeka Barat No. 12. | alamat dari user untuk melakukan proses pengiriman |
| status | STRING("CUSTOMER", ADMIN) | CUSTOMER | status dari user untuk membedakan hak akses saat mengakses suatu endpoint. |
| is_verified | BOOLEAN | 0 | untuk proses verifikasi sehingga user tidak dapat melakukan transaksi jika belum diverifikasi oleh admin (di luar MVP). |

### Categories

Tabel Categories

| Nama Field | Tipe Data | Contoh Nilai | Penjelasan |
| ---------- | --------- | ------------ | ---------- |
| id | INT | 1 | ID digunakan untuk membedakan data antar Categories dan digunakan pada tabel Products untuk mengetahui kategori dari suatu product. |
| nama | STRING | alat olahraga | nama dari kategori, untuk ditampilkan di frontend. |
| kode | STRING | SPORT | digunakan pada URI untuk URI yang lebih singkat dan mudah dipahami oleh end user. |

### Transactions

Tabel Transactions

| Nama Field | Tipe Data | Contoh Nilai | Penjelasan |
| ---------- | --------- | ------------ | ---------- |
| id | INT | 1 | ID digunakan untuk membedakan antar data Transactions dan digunakan pada tabel Transaction Items untuk mengetahui item-item yang terdapat pada suatu transaksi. |
| kode_invoice | STRING(25) | 202102201757001 | Sebagai ID yang ditampilkan ke end user untuk mengetahui status dan item-item yang dibeli dari transaksi yang dimiliki. |
| tanggal_invoice | DATETIME | 2021-02-20 17:59:00 | Untuk mengetahui tanggal invoice dibuat. |
| tanggal_tenggat_pembayaran | DATETIME | 2021-02-23 00:00:00 | Untuk mengetahui tenggat pembayaran yang akan memengaruhi status pembayaran jika transaksi tidak kunjung dilakukan. |
| tanggal_selesai_pembayaran | DATETIME | 2021-02-21 15:00:00 | Untuk mengetahui tanggal dari proses pembayaran telah dilakukan oleh user. |
| tanggal_konfirmasi_pembayaran | DATETIME | 2021-02-21 16:00:00 | Untuk mengetahui tanggal dari proses konfirmasi pembayaran yang dilakun oleh admin. |
| tanggal_terima_pesanan | DATETIME | 2021-02-24 17:00:00 | Untuk mengetahui tanggal dari paket telah diterima oleh pembeli. |
| tanggal_selesai_transaksi | DATETIME | 2021-02-24 17:30:00 | Untuk mengetahui tanggal dari proses konfirmasi transaksi telah selesai yang dilakukan oleh admin. |
| bukti_transfer | STRING(255) | https://unsplash.com/photos/bC0uxTH7aS0 | Link gambar dari bukti transfer atau bukti pembayaran yang lain yang digunakan oleh admin untuk verifikasi. |
| status | INT | 1 | Status dari transactions `1` => `menunggu pembayaran`; `2` => `menunggu konfirmasi pembayaran oleh admin`; `3` => `menunggu pesanan sampai ke customer`; `4` => `pesanan sudah diterima oleh customer`; `5` => `pesanan selesai dikonfirmasi oleh admin`; `-1` => `pesanan gagal atau dibatalkan`. |

## Tabel Relasi
Penjelasan Tabel Relasi, merupakan tabel yang tidak bisa berdiri sendiri karena memerlukan attribute dari tabel lain untuk dapat menambahkan data baru. Contohnya tidak bisa menambahkan data pada Tabel Products jika tidak diberikan nilai `id` dari Tabel Categories.

### Products

Tabel Products

| Nama Field | Tipe Data | Contoh Nilai | Penjelasan |
| ---------- | --------- | ------------ | ---------- |
| id | INT | 1 | ID digunakan untuk membedakan antara data suatu Products dengan data product yang lain dan juga digunakan pada transaksi untuk mengetahui detail dari product yang masuk transaksi. |
| nama | STRING(90) | Buku Dasar Pemrograman Golang oleh Alterra Academy | Nama dari produk yang dijual untuk ditampilkan ke user. |
| stok | INT | 12 | Stok dari produk yang dijual, digunakan pada proses transaksi sehingga hanya dapat membeli produk dengan stok yang sesuai saja. |
| deskripsi | STRING(255) | Buku Dasar Pemrograman Golang ditulis oleh Tim Alterra Academy yang diterbitkan pada tahun 2021 ini meliputi materi algoritma dasar hingga penggunaan Framework Echo untuk membangun sebuah Backend berbasis RESTful API dengan jumlah halaman mencapai 420 halaman. | Deskripsi dari produk yang hendak dijual yang ditampilkan pada user.  
| harga | INT | 180000 | Harga dari produk yang dijual yang digunakan pada proses transaksi.
| categories_id | INT | 1 | ID dari categories produk untuk ditampilkan ke user pada saat proses pencarian atau sorting berdasarkan categories. |
| gambar | STRING(255) | https://unsplash.com/photos/bC0uxTH7aS0 | Link gambar dari produk yang hendak dijual. |

### Transaction Items

Tabel Transactions Items

| Nama Field | Tipe Data | Contoh Nilai | Penjelasan |
| ---------- | --------- | ------------ | ---------- |
| id | INT | 1 |  ID membedakan antara suatu transaction items dengan transaction items yang lain. |
| users_id | INT | 1 | users_id digunakan untuk mengetahui pemilik atau user dari suatu data transaction items. |
| products_id | INT | 1 | products_id digunakan untuk mengetahui detail dari products yang dibeli oleh user. |
| gambar | STRING(255) | https://unsplash.com/photos/bC0uxTH7aS0 | link gambar dari produk yang hendak dibeli supaya jika gambar pada master produk berubah, gambar pada transaction items tetap sesuai seperti saat user membeli produk tersebut. |
| status | INT | 0 | status menandakan status dari transaksi. `0` => `cart`; `1` => `mode transaksi mengikuti status pada tabel Transactions`. |
| harga | INT | 45000 | harga dari suatu products, disimpan di tabel Transaction Items supaya dapat menyimpan harga pada saat user membeli produk. |
| jumlah | INT | 2 | jumlah dari suatu products yang dibeli. |
| harga_total | INT | 90000 | total harga dari suatu transaction items yang didapatkan dari harga dikalikan dengan jumlah item yang dipesan. |
| transactions_id | INT | 1 | untuk mengetahui detail transactions dari suatu transactions items seperti kode invoice dan tanggal transaksi. |  

# Bank Ina Backend Test

# Stacks
- Golang, Sqlite (tidak perlu install aplikasi database sepert mysql atau postgres karena sekarang ini menggunakan database portable)

## Cara menjalankan project ini
- Go Version 1.21.6 
- Masuk ke terminal melalui root folder (Sejajar dengan file README.md) dan jalankan perintah 
    - `go mod tidy` (jika program baru pertama kali di download) 
    - `go run cmd/server/main.go`

## library oauth2 yang digunakan: 
    - https://github.com/go-oauth2/gin-server


## Postman Collections
    - Daftar endpoint bisa cek di file postman_collection.json
    - Hit endpoint oauth2/token untuk mendapatkan token (cek di postmen collection)

## Protected Tasks endpoint
    - Untuk bisa akses endpoint task harus memiliki token dulu dan di set di header dengan key: Authorization, value: Bearer token_yg_digenerate (bisa refer ke postman collection)

## Alur Kerja Program
    - Konsep MVC (model view controller) digunakan untuk membuat program ini dimana banyak framework populer menggunakan konsep ini seperti java springboot, php laravel ataupun nodejs adonisjs 
    - models: komponent yang digunakan untuk mengakses hal yang berkaitan dengan table di database
    - repositories: tempat pembuatan kontrak dan implementasi terhadap model 
    - services: tempat dimana segala hal yang bersifat penentuan alur logika bisnis mencakup proses penyimpandan data atau alur yang memerlukan kondisi tertentu bisa disimpan disini
    - controller: tempat dimana input diterima dan jembatan awal ketika di hit oleh router
    - routers: tempat untuk mapping alamat endpoint dengan controller
    - initializers: tempat untuk init semua komponen mencakup koneksi ke database, init services dan controller, karena golang butuh ini supaya semua alur dapat terkoneksi dan bekerja. Secara naluriah pada saat mengerjakan golang dengan menggunakan dependency injection dan itu terjadi di sini. 
    - Skeleton seperti ini dibuat supaya programmer dapat fokus ke komponen terkecil yang terasa lebih sederhana untuk melakukan suatu implementasi. 
<img width="1536" height="2752" alt="unnamed" src="https://github.com/user-attachments/assets/48b3c2b6-33d6-451b-880c-de9087151850" />

Panduan Belajar: Logika Dasar dan Angka dalam Bahasa Go

1. Pendahuluan: Memulai Langkah Pertama di Golang

Memasuki dunia pemrograman dengan bahasa Go (Golang) diawali dengan langkah yang sangat mendasar: membuat file Go pertama Anda. Berdasarkan materi "3 file go pertama", esensi dari memulai bukan sekadar mengetik, melainkan memahami bagaimana instruksi disusun dalam struktur yang dikenali oleh compiler. Langkah awal ini krusial karena di sinilah Anda meletakkan fondasi untuk memahami bagaimana sebuah program berinteraksi dengan data melalui variabel.

"Setiap aplikasi kelas dunia selalu dimulai dari sebuah file sederhana dan pemahaman yang presisi mengenai cara menyimpan data. Jangan ragu untuk bereksperimen, karena penguasaan fondasi adalah kunci kecepatan Anda di masa depan."

Setelah Anda berhasil menjalankan file pertama, langkah logis selanjutnya adalah memahami variabel—wadah di memori komputer yang kita gunakan untuk menyimpan dan mengelola informasi.

2. Fondasi Data: Variabel dan Tipe Data String

Dalam bahasa Go, variabel adalah elemen kunci untuk menyimpan informasi. Berdasarkan sintesis dari materi "mengenal variable", "variabel kosong", dan "penulisan singkat variable", Go menawarkan fleksibilitas deklarasi yang harus disesuaikan dengan kebutuhan struktur kode. Sebagai pengembang, Anda juga harus membiasakan penggunaan komentar (Video 6) untuk mendokumentasikan tujuan dari setiap variabel yang Anda buat.

Berikut adalah perbandingan cara mendeklarasikan variabel di Go:

Tipe Deklarasi	Contoh Penulisan	Kapan Digunakan
Biasa (Explicit)	var nama string = "Budi"	Digunakan di level package atau saat ingin mendefinisikan tipe data secara eksplisit demi kejelasan dokumen.
Tanpa Nilai Awal	var nama string	Digunakan saat variabel dideklarasikan di awal namun nilainya baru akan diisi kemudian (menginisialisasi zero value).
Singkat (Short Declaration)	nama := "Budi"	Hanya berlaku di dalam fungsi (local scope). Sangat efisien karena Go melakukan type inference secara otomatis.

Memahami tempat penyimpanan teks (string) adalah awal, namun aplikasi nyata sangat bergantung pada pengolahan data numerik untuk menjalankan logika bisnis.

3. Eksplorasi Tipe Data Angka

Berdasarkan materi "8 tipe data angka", Go menyediakan kontrol yang sangat spesifik terhadap penggunaan memori. Angka bukan sekadar nilai; pemilihan tipe angka yang tepat mencerminkan efisiensi aplikasi yang Anda bangun.

Karakteristik tipe data angka yang wajib dipahami pemula:

1. Integer (Bilangan Bulat): Digunakan untuk angka tanpa koma. Go menyediakan tipe seperti int8, int32, hingga int64. Penting untuk diketahui bahwa tipe int (tanpa angka di belakangnya) bersifat platform-dependent, yang berarti ukurannya akan mengikuti arsitektur prosesor Anda (32-bit atau 64-bit).
2. Floating Point (Bilangan Desimal): Digunakan untuk angka dengan presisi koma, seperti float32 dan float64.
3. Alias dan Spesifikasi: Pengembang harus memilih kapasitas yang sesuai dengan rentang angka yang akan diolah untuk mencegah memory overflow.

Setelah mengenal jenis angka, langkah logis berikutnya adalah mempelajari bagaimana memanipulasi angka-angka tersebut untuk menghasilkan informasi yang berguna.

4. Mengolah Logika dengan Operasi Matematika

Mengacu pada materi "9 operasi matematika di go", operator adalah alat utama untuk membangun logika aplikasi. Bukan sekadar hitungan dasar, operasi ini adalah mesin di balik logika seperti perhitungan diskon, penentuan koordinat, hingga sistem rating.

Operator dasar yang tersedia di Go meliputi:

* Penjumlahan (+)
* Pengurangan (-)
* Perkalian (*)
* Pembagian (/)
* Modulo (%) (Sisa bagi)

Expert Insight: Berhati-hatilah saat melakukan pembagian antar integer. Di Go, hasil pembagian integer akan selalu berupa integer (terjadi pemotongan desimal/truncation). Misalnya, 5 / 2 akan menghasilkan 2, bukan 2.5.

Contoh penulisan kode dalam Go:

package main

func main() {
    // Menghitung sisa stok setelah pembelian
    totalStok := 100
    jumlahBeli := 15
    sisaStok := totalStok - jumlahBeli
    
    // Perhitungan rata-rata (pentingnya konversi jika butuh desimal)
    rataRata := 10 / 3 // Hasilnya adalah 3
}


Terkadang kita membutuhkan nilai yang tidak boleh berubah sedikit pun selama aplikasi berjalan, yang membawa kita pada konsep konstanta.

5. Nilai Konstan: Keamanan dalam Ketetapan

Berdasarkan materi "10 nilai konstan", Konstanta adalah nilai yang tidak dapat diubah setelah dideklarasikan. Pengembang menggunakan konstanta untuk melindungi nilai-nilai kritikal agar tidak termodifikasi secara tidak sengaja oleh bagian kode lain.

Dibandingkan sekadar menggunakan angka ajaib (magic numbers), konstanta memberikan konteks yang jelas. Contoh penggunaan nyata adalah menentukan Application Port Number atau API Base URL.

Best Practice: Sangat direkomendasikan untuk menggunakan kata kunci const pada nilai yang bersifat statis untuk menjamin integritas data dan mempermudah pemeliharaan kode (maintenance).

Seringkali, input yang kita terima dari pengguna atau sistem eksternal datang dalam format teks (string), namun kita butuh mengolahnya secara matematis.

6. Konversi Data: Menghubungkan String dan Integer

Materi "11 convert string dan integer" menyoroti tantangan umum: data dari formulir atau API biasanya bertipe string. Tanpa konversi yang tepat, operasi matematika tidak mungkin dilakukan. Dalam Go, kita menggunakan paket bawaan strconv untuk menangani hal ini.

Salah satu fungsi yang paling sering digunakan adalah strconv.Atoi (Alpha to Integer). Karena Go sangat mengutamakan keamanan, fungsi ini tidak hanya mengembalikan hasil konversi, tetapi juga nilai error.

Daftar Periksa Konversi Data:

* [ ] Import Package: Pastikan Anda telah mengimpor package strconv.
* [ ] Validasi Input: Pastikan string hanya berisi karakter angka; "100a" akan menyebabkan kegagalan konversi.
* [ ] Error Handling: Wajib menangani nilai balik kedua (err) untuk memastikan program tidak crash jika konversi gagal.
* [ ] Tipe Data Tujuan: Pastikan hasil konversi disimpan dalam variabel dengan tipe data yang sesuai untuk operasi selanjutnya.

7. Kesimpulan dan Refleksi Belajar

Memahami variabel, tipe data angka, operasi matematika, konstanta, hingga teknik konversi merupakan satu kesatuan kepingan puzzle yang membentuk dasar logika pemrograman di bahasa Go. Elemen-elemen ini adalah fondasi yang mutlak dikuasai sebelum Anda melangkah ke topik yang lebih kompleks seperti pointers, structs, atau concurrency. Ingatlah bahwa kode yang baik bukan hanya kode yang berjalan, tetapi kode yang aman dan mudah dipahami oleh pengembang lain. Teruslah berlatih, karena kemahiran dalam logika dasar adalah investasi terbaik bagi setiap pengembang kurikulum dan praktisi pemrograman.

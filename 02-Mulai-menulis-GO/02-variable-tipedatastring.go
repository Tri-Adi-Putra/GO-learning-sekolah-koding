package main

import "fmt"

//tipe data kosong
var temp = "" 
var teks string

//penulisan model pertama
var namaDepan = "tri"
var namaBelakang = "putra"
var namaGabungan = namaDepan + " " + namaBelakang

//penulisan model ke dua
var nama, samping = "agnes", "levia"

func main() {

    sekolah := "universita Mercu" //short hand

	fmt.Println("saya Kuliah di ", sekolah)

    
    fmt.Println("Nama Depan saya " + " " + namaDepan)
    fmt.Println("perkenalkan nama saya " + " " + namaGabungan)
    
    // Perhatikan huruf P kapital pada Println
    fmt.Println("Nama Istri Saya " + " " + nama + " " + samping)
}

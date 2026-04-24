package main

import "fmt"

var namaDepan = "tri"

var namaBelakang = "putra"

var namaGabungan = namaDepan + " " + namaBelakang

func main() {

	var hasil int //perhitungan latihan modulo
	a := 32
	v := 10
	hasil = a % v

	//cara menaikan nilai angka dan menurunkan
	r := 10
	r++

	t := 21
	t--

	sekolah := "universita Mercu" //short hand

	//tipe data numeric/number
	umur := 50

	fmt.Println("umur kakek saya ", umur)
	fmt.Println("saya Kuliah di ", sekolah)
	fmt.Println("Nama Depan saya" + namaDepan)
	fmt.Println("perkenalkan nama saya" + namaGabungan)
	fmt.Println("perhitungan hasil modulo a dan v ", hasil)
	fmt.Println("angka 10 jika di naikan 1 nilai men jadi angka :", r)
}

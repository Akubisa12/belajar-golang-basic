package main

import "fmt"

func main(){

	//membuat constanta variabel(tidak bisa diubah)
	// const tinggi = 100
	//coba kita ubah
	// tinggi = 50 //tidak bisa diubah
	// fmt.Println("nilai tinggi =", tinggi)

	//operator
	// 3 macam operator

	// 1. Aritmatika
	// var jumlah = ((2+4)*2)/4
	// fmt.Println("hasil jumlah =", jumlah)

	// 2. Perbandingan
	// var jumlah = ((2+4)*2)/4
	// var uji = (jumlah == 3)
	// fmt.Printf("hasil jumlah = %d (%t) \n", jumlah, uji)

	// 3. Operator logika
	var kiri  = true
	var kanan = false
	kirikanan := kiri && kanan 
	kiriorkanan := kiri || kanan 
	notkiri := !kiri 
	fmt.Printf("kiri and kanan = %t \n", kirikanan)
	fmt.Printf("kiri or kanan = %t \n", kiriorkanan)
	fmt.Printf("not kiri = %t \n", notkiri)
}
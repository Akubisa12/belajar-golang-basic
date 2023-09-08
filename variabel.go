package main

import "fmt"

func main(){
	// 3 cara penulisan variabel di Go

	// 1.variabel beserta tipe data
	// var namadepan string = "rahul"

	// atau bisa juga seperti ini
	// var namablkg string 
	// namablkg = "subagio"

	// mencetak menggunakan printf
	// fmt.Printf("halo %s %s \n", namadepan, namablkg)

	// mencetak menggunakan println
	// fmt.Println("hallo " + namadepan + " " + namablkg)

	// 2. variabel tanpa tipe data

	// cara pertama
	// var namadepan = "rahul"

	//cara kedua
	// namablkg := "subagio"

	// mencetak menggunakan println
	// fmt.Println("hallo " + namadepan + " " + namablkg)

	//3. variabel underscore

	// cara pertama
	// _ = "sampah" // ini tidak bisa ditampilkan

	//cara kedua
	// namablkg, _ := "subagio", "sampah"
}
package main

import "fmt"

func main()  {

	//cara pertama membuat perulangan

	for i:= 0; i < 5; i++{ //mengulang sebanyak 5 kali
		fmt.Println("Angka",i)
	}

	// cara kedua membuat perulangan
	var i = 0
	for i < 5 {
		fmt.Println("Angka",i)
		i++
	}

	//cara ketiga membuat perulangan
	var i = 0
	for{
		fmt.Println("Angka",i)
		i++
		if i == 5{
			break // memberhentikan  program secara paksa
		}
	}
}
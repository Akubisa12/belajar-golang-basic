package main

import "fmt"

func main(){

	// cara pertama membuat array
	var buah [4]string 
	buah[0] = "Apel"
	buah[1] = "Jeruk"
	buah[2] = "Semangka"
	buah[3] = "Jambu"

	// output tanpa perulangan
	fmt.Println("jumlah data :",len(buah))// fungsi len ini gunanya untuk mengambil jumlah data yang ada didalam satu variabel
	fmt.Println("data arrray :",buah)

	// output dengan perulangan
	for i := 0; i < len(buah); i++ {
		fmt.Println("buah [",i,"]",buah[i])
	}

	//cara kedua membuat array
	var angka = [...]int{3,2,5,8,3} 
	fmt.Println("Jumlah data ",len(angka))
	fmt.Println("data array :", angka)

	// slice 
	// slice = bisa menbambahkan data ditengah tengah
	var buah[]string 
	buah = append(buah, "apel")
	buah = append(buah, "semangka")
	buah = append(buah, "jeruk")
	buah = append(buah, "mangga")

	fmt.Println("sebelum ditambah data")
	fmt.Println("jumlah data :",len(buah))
	fmt.Println("data array :",buah)

	buah = append(buah,"durian")
	buah = append(buah,"jambu")

	fmt.Println("sesudah ditambah data")
	fmt.Println("jumlah data :",len(buah))
	fmt.Println("data array :",buah)

	//menggunakan perulangan for 
	for i := 0; i < len(buah); i++ {
		fmt.Println("buah [",i,"] :",buah[i])
	}
}
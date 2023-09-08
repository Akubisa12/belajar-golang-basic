package main

import "fmt"

func main(){
	// contoh tipe data yang pertama(uint8 dan uint16)
	// memiliki batn 255

	var angka uint8 = 123
	fmt.Printf("angka = %d \n", angka)

	//kita coba lebihkan dari batas
	var angka uint8 = 256
	fmt.Printf("angka = %d \n", angka)

	// cara mengatasi nya
	// kita gunakan uint16 dengan batas maksimal 65535
	var angka uint16 = 256
	fmt.Printf("angka = %d \n", angka)

	//contoh tipe data kedua(int8 dan int16)
	var angka int8 = -123
	fmt.Printf("angka = %d \n", angka)

	// kita coba lebihkan batas(batas max= 128)
	var angka int8 = -129
	fmt.Printf("angka = %d \n", angka)

	// cara mengatasinya kita ganti dengan tipe data(int)
	var angka int = -129
	fmt.Printf("angka = %d \n", angka)

	//contoh tipe data ketiga(float)
	var desimal float32 = 2.534
	fmt.Println("angka desimal =", desimal)

	//menentukan berapa angka dibelakang koma
	var desimal float32 = 2.534
	fmt.Printf("angka desimal = %.5f", desimal) // .5f merupakan jumlah angka dibelakang koma

	// contoh tipe data keempat(bool)
	var sampai bool = true
	fmt.Printf("sudah? = %t", sampai) // %t untuk mencetak tipe data boolean
}
package main

import "fmt"

func main(){
	//membuat program ttg nilai(pengendalian)
	var nilai int 
	fmt.Print("Masukan nilai : ")
	 fmt.Scan(&nilai) //tanda & itu digunakan untuk memasukan variabel ke dalam inputan user
	if nilai == 10 {
	fmt.Println("sangat bagus")
	}else if nilai < 10 && nilai >= 5 {
	fmt.Println("bagus")
	}else if nilai < 5 && nilai > 1 {
	fmt.Println("buruk")
	}else {
	fmt.Println("coba lagi")
	}

	// seleksi kondisi menggunakan switch case
	var nilai int 
	fmt.Print("Masukan nilai : ")
	fmt.Scan(&nilai) //tanda & itu digunakan untuk memasukan variabel ke dalam inputan user
	switch nilai {
	case 10:
		fmt.Println("sangat bagus")
	case 8 :
		fmt.Println("bagus")
	case 5 :
		fmt.Println("buruk")
	default :
		fmt.Println("coba lagi")
	}
}
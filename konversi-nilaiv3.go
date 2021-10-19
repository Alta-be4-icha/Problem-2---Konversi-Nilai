package main

import "fmt"

func main() {
	var Nilai float64
	fmt.Print("Masukkan Nilai: ")
	fmt.Scanln(&Nilai)
	var studenscore int = 80
	if Nilai > 100 {
		fmt.Println("Nilai Yang Anda Masukkan Salah")
	} else {
		if Nilai >= float64(studenscore) {
			fmt.Println("A")
		} else if Nilai >= 65 && Nilai <= 75 {
			fmt.Println("B+")
		} else if Nilai >= 50 && Nilai <= 64 {
			fmt.Println("B")
		} else if Nilai > 35 && Nilai <= 49 {
			fmt.Println("C")
		} else if studenscore <= 34 {
			fmt.Println("D")
		}
	}
}

package main

import (
	"fmt" // Mengimpor package fmt untuk input/output
)

func main() {
	var angka1, angka2 float64 // Menyimpan dua angka bertipe desimal
	var operator string

	fmt.Println("=== GO Calculator ===")
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&angka1) // Membaca input dan menyimpannya ke variabel

	fmt.Print("Masukkan operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&angka2)

	var hasil float64

	switch operator { // Mengecek operator yang dimasukkan user
	case "+": // Menjalankan operasi sesuai pilihan user
		hasil = angka1 + angka2
	case "-":
		hasil = angka1 - angka2
	case "*":
		hasil = angka1 * angka2
	case "/":
		if angka2 != 0 { // Mengecek agar tidak terjadi pembagian dengan nol
			hasil = angka1 / angka2
		} else {
			fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan")
			return
		}
	default:
		fmt.Println("Operator tidak valid")
		return
	}

	fmt.Printf("Hasil: %g %s %g = %d\n", angka1, operator, angka2, int(hasil)) // Menampilkan hasil
}

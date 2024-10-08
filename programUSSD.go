package main

import "fmt"

func menu_transfer_pulsa() {
	var nomorHandphone string
	var nominalPulsa int
	var konfirmasi string

	fmt.Print("Masukkan nomor HP: ")
	fmt.Scan(&nomorHandphone)
	// memeriksa panjang nomor handphone, 10-12 digit
	if len(nomorHandphone) >= 10 && len(nomorHandphone) <= 12 {
		fmt.Print("Masukkan nominal pulsa yang ingin dikirim: ")
		fmt.Scan(&nominalPulsa)
		fmt.Printf("Kirim pulsa sebesar %d ke %s?\n", nominalPulsa, nomorHandphone)
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Print("Masukkan: ")
		fmt.Scan(&konfirmasi)
		if konfirmasi == "1" {
			fmt.Println("Pulsa berhasil dikirim!")
		} else if konfirmasi == "2" {
			fmt.Println("Pulsa gagal dikirim!")
		} else {
			fmt.Println("Opsi tidak valid!")
		}
	} else {
		fmt.Println("Nomor handphone tidak ditemukan")
	}
}

func menu_utama() {
	var opsi string

	fmt.Println("Selamat datang di menu USSD\nketik angka untuk mengakses menu")
	fmt.Println("1. Transfer Pulsa")
	fmt.Println("2. Keluar")
	fmt.Print("Masukkan: ")
	fmt.Scan(&opsi)

	switch opsi {
	case "1":
		menu_transfer_pulsa()
	case "2":
		break
	default:
		fmt.Println("angka tidak valid, mohon coba lagi")
		fmt.Scan(&opsi)
	}
}

func programUSSD() {
	var USSDcode string
	for true {
		fmt.Print("Masukkan kode USSD: ")
		fmt.Scan(&USSDcode)
		if USSDcode == "*858#" {
			menu_utama()
		} else {
			fmt.Println("USSD tidak valid, mohon coba lagi")
		}
	}
}

func main() {
	programUSSD()
}

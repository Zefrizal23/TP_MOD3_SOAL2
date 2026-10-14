package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan (N): ")
	fmt.Scan(&n)
	hariAwal := 4
	NDC := (hariAwal-1+n)%7 + 1
	Hari := [8]string{"", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Printf("%d hari setelah hari Kamis (4) adalah %s (%d)\n", n, Hari[NDC], NDC)
}
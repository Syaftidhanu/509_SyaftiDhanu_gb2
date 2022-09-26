package main

import (
	"fmt"
	"os"
	"strconv"
)

type Peserta struct {
	No        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {

	var Peserta = []Peserta{
		{No: 1, Nama: "Syafti", Alamat: "Bekasi", Pekerjaan: "Mahasiswa", Alasan: "Mendapatkan sertifikasi"},
		{No: 2, Nama: "yuga", Alamat: "Bekasi", Pekerjaan: "Mahasiswa", Alasan: "Mendapatkan sertifikasi"},
		{No: 3, Nama: "biba", Alamat: "Bekasi", Pekerjaan: "Mahasiswa", Alasan: "Mendapatkan sertifikasi"},
		{No: 4, Nama: "abul", Alamat: "Bekasi", Pekerjaan: "Mahasiswa", Alasan: "Mendapatkan sertifikasi"},
		{No: 5, Nama: "adul", Alamat: "Bekasi", Pekerjaan: "Mahasiswa", Alasan: "Mendapatkan sertifikasi"},
	}

	//}
	//fmt.Printf(Peserta[2])
	// for _, v := range Peserta {
	// 	fmt.Println("Nama :", v.Nama)
	// }

	// var absen int
	// fmt.Println("Lhat Data Peserta Ke - ")
	// fmt.Scan(&absen)
	var argsRaw = os.Args
	var args = argsRaw[1]
	no_absen, err := strconv.Atoi(args)
	_ = err

	if no_absen <= len(Peserta) {
		fmt.Println("Data Siswa:")
		fmt.Println("No Absen:", Peserta[no_absen-1].No)
		fmt.Println("Nama:", Peserta[no_absen-1].Nama)
		fmt.Println("Alamat:", Peserta[no_absen-1].Alamat)
		fmt.Println("Pekerjaan:", Peserta[no_absen-1].Pekerjaan)
		fmt.Println("Alasan:", Peserta[no_absen-1].Alasan)
	} else {
		fmt.Println("Student not exist!")
	}
}

// func .nama fungsi.(namastruct[]nama inisialisai struct) output
//func absensi(peserta[]peserta) peserta  {

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Struct untuk menyimpan data teman-teman kelas
type TemanKelas struct {
	ID         string `json:"id"`
	Code       string `json:"student_code"`
	Name       string `json:"student_name"`
	Address    string `json:"student_address"`
	Occupation string `json:"student_occupation"`
	Reason     string `json:"joining_reason"`
}

func main() {
	// Mengambil argumen dari command line
	args := os.Args

	// Pastikan argumen yang diharapkan ada
	if len(args) != 2 {
		fmt.Println("Cara penggunaan: go run biodata.go [student_code]")
		return
	}

	// Membaca file JSON
	data, err := ioutil.ReadFile("participants.json")
	if err != nil {
		fmt.Println("Gagal membaca file JSON:", err)
		return
	}

	// Parse data JSON ke dalam slice TemanKelas
	var temanKelas []TemanKelas
	if err := json.Unmarshal(data, &temanKelas); err != nil {
		fmt.Println("Gagal parsing data JSON:", err)
		return
	}

	// Mendapatkan student code dari argumen
	studentCode := args[1]

	// Cari teman berdasarkan student code
	studentCodeToFind := args[1] // Ganti dengan student code yang ingin Anda cari

	var foundStudent TemanKelas // TemanKelas adalah struktur data yang digunakan

	for _, teman := range temanKelas { // temanKelas adalah slice yang berisi data teman-teman
		if teman.StudentCode == studentCodeToFind {
			foundStudent = teman
			break
		}
	}

	if foundStudent.StudentCode == "" {
		fmt.Println("Tidak ada teman dengan student code", studentCodeToFind)
	} else {
		fmt.Println("Data teman kelas:")
		fmt.Println("Student Code:", foundStudent.StudentCode)
		fmt.Println("Nama:", foundStudent.Name)
		fmt.Println("Alamat:", foundStudent.Address)
		fmt.Println("Pekerjaan:", foundStudent.Occupation)
	}
}

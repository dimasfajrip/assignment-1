package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Struktur data yang digunakan untuk merepresentasikan informasi mahasiswa
type Student struct {
	ID          string `json:"id"`
	StudentCode string `json:"student_code"`
	Name        string `json:"student_name"`
	Address     string `json:"student_address"`
	Occupation  string `json:"student_occupation"`
	Reason      string `json:"joining_reason"`
}

// Struktur data yang digunakan untuk membaca file JSON yang berisi daftar mahasiswa
type Students struct {
	Participants []Student `json:"participants"`
}

func main() {
	// Memastikan bahwa ada tepat satu argumen pada baris perintah
	if len(os.Args) != 2 {
		fmt.Println("Harap masukkan kode mahasiswa setelah .go")
		return
	}

	// Membuka dan membaca isi file JSON yang bernama "participants.json"
	jsonFile, err := os.Open("participants.json")
	if err != nil {
		fmt.Println("Error ketika membuka file:", err)
		return
	}
	defer jsonFile.Close()

	// Membaca seluruh isi file JSON ke dalam variabel byteValue
	byteValue, _ := io.ReadAll(jsonFile)

	// Mendekode isi JSON ke dalam variabel struct students
	var students Students
	err = json.Unmarshal(byteValue, &students)

	// Menangani kesalahan jika terjadi error saat mendekode JSON
	if err != nil {
		fmt.Println("Error saat mendekode JSON:", err)
		return
	}

	// Mengambil kode mahasiswa dari argumen baris perintah
	codeMahasiswa := os.Args[1]

	// Mencari dan menampilkan data mahasiswa berdasarkan kode mahasiswa
	searchByCode(students.Participants, codeMahasiswa)
}

// Fungsi untuk mencari data mahasiswa berdasarkan kode mahasiswa
func searchByCode(students []Student, code string) {
	for _, student := range students {
		if student.StudentCode == code {
			fmt.Printf("\n")
			fmt.Printf("ID Mahasiswa    : %s\n", student.ID)
			fmt.Printf("Kode Mahasiswa  : %s\n", student.StudentCode)
			fmt.Printf("Nama            : %s\n", student.Name)
			fmt.Printf("Alamat          : %s\n", student.Address)
			fmt.Printf("Pekerjaan       : %s\n", student.Occupation)
			fmt.Printf("Alasan Bergabung: %s\n", student.Reason)
			fmt.Printf("\n")
			return
		}
	}
	fmt.Println("Mahasiswa dengan kode", code, "tidak ditemukan.")
}

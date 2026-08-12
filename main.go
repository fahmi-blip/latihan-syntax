package main

import "fmt"

type User struct {
	ID 			int			`json:"id"`
	Username 	string 		`json:"username"`
	IsActive 	bool 		`json:"is_active"`
}

func (u *User) Activate() { u.IsActive = true } 
func (u User) Info() string { return fmt.Sprintf("%s aktif: %v", u.Username, u.IsActive) } 


func swap(a,b *int) {
	tes := *a
	*a = *b
	*b = tes
}

func passByValue(a int){
	a = 10
}
func passByPointer(a *int){
	*a = 10
}

func updateSlice(s *[]string, newItem string){
	*s = append(*s, newItem)
}


func main() { 

// 2. Variabel dan Struktur Data

	// Variabel 
	nama := "Sari" 
	umur := 20 
	tinggi := 165.6 
	isActive := true 
	nilai := []int{10, 20, 30} 
	fmt.Println(nama,"\n", umur, "\n", tinggi, "\n" ,isActive,"\n", nilai)  

	// Map
	mahasiswa := make(map[string]int)

	//menambah data
	mahasiswa["Fahmi"] = 90
	mahasiswa["Sari"] = 85
	mahasiswa["Budi"] = 80

	//membaca dengan pengecekan keberadaan
	if n, ada := mahasiswa["Budi"]; ada { 
		fmt.Println("Budi:", n) 
	} else { 
		fmt.Println("Budi belum punya nilai") 
	}
	
	//menghapus Data
	delete(mahasiswa, "Budi")
	fmt.Println("Setelah dihapus: ", mahasiswa)
	
	//menelusuri seluruh data
	fmt.Println("Seluruh data mahasiswa:")

	for nama, nilai := range mahasiswa {
		fmt.Println("Nama:", nama, "| Nilai:", nilai)
	}


//3. Pointer

	//Fungsi swap
	nilai1, nilai2  := 10, 15
	
	fmt.Println("Sebelum diubah: ", nilai1, nilai2)
	swap(&nilai1, &nilai2)
	fmt.Println("Setelah diubah: ", nilai1, nilai2)

	x :=5
	passByValue(x)
	fmt.Println("pass by value:",x)

	passByPointer(&x)
	fmt.Println("pass by value:", x)

	
	//Fungsi updateSlice
	nilai3 := []string{"a", "b","c"} 
	fmt.Println("Sebelum menambahkan item ke slice:",nilai3)

	nilai3 = append(nilai3, "d") 
	fmt.Println("Setelah Menambahkan item ke slice:",nilai3)


//4. Struct Student


}
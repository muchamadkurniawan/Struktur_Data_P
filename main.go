package main

import "fmt"

var counter int = 0
var DBMember = [10]Member{}

func PrintCounter() {
	fmt.Println("counter : ", counter)
}

func InsertMember(anggota Member) {
	DBMember[counter].Id = anggota.Id
	DBMember[counter].Name.first_name = anggota.Name.first_name
	counter++
}

func ViewMember() {
	for i := 0; i < counter; i++ {
		fmt.Println("Id member : ", DBMember[i].Id)
		fmt.Println("Nama member : ", DBMember[i].Name.first_name)
		fmt.Println("__________")
	}
}

func main() {
	pelanggan1 := Member{
		Id:        1,
		Name:      Nama{"kurniawan", " ", "_"},
		Alamat:    "Surabaya",
		Point:     90.9,
		Create_at: "",
	}

	pelanggan2 := Member{
		Id:        2,
		Name:      Nama{"aan", " ", "_"},
		Alamat:    "malang",
		Point:     0.9,
		Create_at: "",
	}

	InsertMember(pelanggan1)
	InsertMember(pelanggan2)
	ViewMember()
}

package main

type Member struct {
	Id        int
	Name      Nama
	Alamat    string
	Point     float32
	Create_at string
}

type Nama struct {
	first_name  string
	middle_name string
	last_name   string
}

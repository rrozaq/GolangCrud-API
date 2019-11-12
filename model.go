package main

// membuat struct yang nantinya akan kita tampilkan di response
type Users struct {
	Id			int `form:"id" json:"id"`
	Npm			string `form:"npm" json:"npm"`
	Nama		string `form:"nama" json:"nama"`
	Jk		string `form:"jenis kelamin" json:"jk"`
}

// membuat response
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}
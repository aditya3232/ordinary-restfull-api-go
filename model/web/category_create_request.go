package web

/*
	- struct request & response perlu dibuat
	- karena tidak semua data dari database ditampilkan di response
	- dan tidak semua data bisa jadi request
	- jadi data model request & response perlu dibuat
*/

type CategoryCreateRequest struct {
	Name string
}

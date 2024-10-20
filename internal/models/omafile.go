package models

type Omafile struct {
	Id        int
	Brand     string
	Model     string
	Info      string
	Directory string
}

func NewOmafile(parametr ...string) Omafile {
	return Omafile{
		Brand:     parametr[0],
		Model:     parametr[1],
		Info:      parametr[2],
		Directory: parametr[3],
	}
}

func GetOmafile(id int, parametr ...string) Omafile {
	return Omafile{
		Id:        id,
		Brand:     parametr[0],
		Model:     parametr[1],
		Info:      parametr[2],
		Directory: parametr[3],
	}
}

package models

type Omafile struct {
	Id        int    `json:"id"`
	Brand     string `json:"Brand"`
	Model     string `json:"Model"`
	Info      string `json:"Desctiprion"`
	Directory string `json:"Directory"`
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

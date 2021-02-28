package model

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	BirthDate string `json:"birth_date"`
	Ktp string `json:"ktp"`
	Job string `json:"job"`
 	LastEducation string `json:"last_education"`
}

func (u User) TableName() string {
	return "mst_user"
}

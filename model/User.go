package model

type User struct {
	Id int `json:"id" gorm:"primaryKey unique not null"`
	Name string `json:"name" gorm:"not null"`
	BirthDate string `json:"birth_date" gorm:"not null"`
	Ktp string `json:"ktp" gorm:"not null size:16 unique"`
	Job string `json:"job" gorm:"not null"`
 	LastEducation string `json:"last_education" gorm:"not null"`
}

func (u User) TableName() string {
	return "mst_user"
}

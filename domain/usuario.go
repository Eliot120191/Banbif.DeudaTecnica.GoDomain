package domain

import (
	"gorm.io/gorm"
)

type USUARIO struct {
	ID     int    `json:"id" gorm:"type:int;not null;primary_key;autoIncrement"`
	Correo string `json:"correo" gorm:"type:varchar(200);"`
	/*Nombre    string `json:"nombre" gorm:"type:char(200);not null"`
	Apellidos string `json:"apellidos" gorm:"type:char(200);not null"`*/
}

func (u *USUARIO) BeforeCreate(tx *gorm.DB) (err error) {

	return
}

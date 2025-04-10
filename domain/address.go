package domain

import (
	"gorm.io/gorm"
)

type ADDRESS struct {
	ID         int64  `json:"id" gorm:"type:int;not null;primary_key;autoIncrement"`
	Id_Usuario string `json:"id_usuario" gorm:"type:int;not null"`
	Address    string `json:"address" gorm:"type:varchar(100);not null"`
	/*Nombre    string `json:"nombre" gorm:"type:char(200);not null"`
	Apellidos string `json:"apellidos" gorm:"type:char(200);not null"`*/
}

func (u *ADDRESS) BeforeCreate(tx *gorm.DB) (err error) {

	return
}

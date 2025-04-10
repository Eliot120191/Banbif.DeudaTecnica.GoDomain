package domain

import (
	"gorm.io/gorm"
)

type USUARIO struct {
	ID        int64  `json:"id" gorm:"type:int;not null;primary_key;autoIncrement"`
	Correo    string `json:"correo" gorm:"type:varchar(200);"`
	Nombre    string `json:"nombre" gorm:"type:varchar(200);not null"`
	Apellidos string `json:"apellidos" gorm:"type:varchar(200);not null"`
}

func (USUARIO) TableName() string {
	return "USUARIO"
}

func (u *USUARIO) BeforeCreate(tx *gorm.DB) (err error) {

	return
}

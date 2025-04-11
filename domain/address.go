package domain

import (
	"gorm.io/gorm"
)

type ADDRESS struct {
	BASE
	ID         int64  `json:"id" gorm:"type:int;not null;primary_key;autoIncrement"`
	Id_Usuario string `json:"id_usuario" gorm:"type:int;not null"`
	Address    string `json:"address" gorm:"type:varchar(100);not null"`
	Tx         string `json:"tx" gorm:"type:varchar(250);not null"`
	Tipo       string `json:"tipo" gorm:"type:char(1);not null"`
}

func (ADDRESS) TableName() string {
	return "ADDRESS"
}

func (u *ADDRESS) BeforeCreate(tx *gorm.DB) (err error) {

	return
}

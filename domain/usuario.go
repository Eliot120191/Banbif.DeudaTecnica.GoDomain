package domain

import (
	"gorm.io/gorm"
)

type USUARIO struct {
	BASE           BASE   `json:"base,omitempty"`
	Correo         string `json:"correo" gorm:"type:varchar(200);"`
	Nombre         string `json:"nombre" gorm:"type:varchar(200);not null"`
	Apellidos      string `json:"apellidos" gorm:"type:varchar(200);not null"`
	Cod_Rol        string `json:"cod_rol" gorm:"type:varchar(50);null"`
	Cod_Grupo      string `json:"cod_grupo" gorm:"type:varchar(50);null"`
	Code           string `json:"code" gorm:"type:varchar(1000);null"`
	Secret_Backend string `json:"secret_backend" gorm:"type:varchar(1000);null"`
}

func (USUARIO) TableName() string {
	return "USUARIO"
}

func (u *USUARIO) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

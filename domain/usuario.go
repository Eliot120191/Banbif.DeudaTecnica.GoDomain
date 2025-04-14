package domain

import (
	"time"

	"gorm.io/gorm"
)

type USUARIO struct {
	ID                   int64     `json:"id" gorm:"type:int;not null;primary_key;autoIncrement"`
	Es_Activo            bool      `json:"es_activo" gorm:"type:bool;not null;"`
	Usuario_Creacion     int64     `json:"usuario_creacion" gorm:"type:int;null;"`
	Fecha_Creacion       time.Time `json:"fecha_creacion" gorm:"type:datetime;not null;"`
	Usuario_Modificacion int64     `json:"usuario_modificacion" gorm:"type:int;null;"`
	Fecha_Modificacion   time.Time `json:"fecha_modificacion" gorm:"type:datetime;null;"`

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

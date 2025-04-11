package domain

import (
	"time"

	"gorm.io/gorm"
)

type BASE struct {
	gorm.Model
	ID                   int64     `json:"id" gorm:"type:int;not null;primary_key;autoIncrement"`
	Es_Activo            bool      `json:"es_activo" gorm:"type:datetime;not null;"`
	Usuario_Creacion     int64     `json:"usuario_creacion" gorm:"type:int;null;"`
	Fecha_Creacion       time.Time `json:"fecha_creacion" gorm:"type:datetime;not null;"`
	Usuario_Modificacion int64     `json:"usuario_modificacion" gorm:"type:int;null;"`
	Fecha_Modificacion   time.Time `json:"fecha_modificacion" gorm:"type:datetime;null;"`
}

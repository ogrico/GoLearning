package models

import (
	"time"
)

type CuentaBancaria struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	Cedula             string    `gorm:"size:20;not null" json:"cedula"`
	NombreCompleto     string    `gorm:"size:100;not null" json:"nombre_completo"`
	NumeroCuenta       string    `gorm:"size:20;not null" json:"numero_cuenta"`
	EntidadCuenta      string    `gorm:"size:50;not null" json:"entidad_cuenta"`
	SaldoDisponible    float64   `gorm:"type:decimal(15,2);not null" json:"saldo_disponible"`
	FechaActualizacion time.Time `gorm:"autoUpdateTime" json:"fecha_actualizacion"`
	FechaCreacion      time.Time `gorm:"autoCreateTime" json:"fecha_creacion"`
	EstadoCuenta       string    `gorm:"type:enum('activa', 'inactiva');not null" json:"estado_cuenta"`
}

func (CuentaBancaria) TableName() string {
	return "cuentasbancarias"
}

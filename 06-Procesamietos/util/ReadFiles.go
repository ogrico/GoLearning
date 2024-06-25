package util

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"procesamientos/models"
	"strconv"
	"strings"
	"time"
)

func ValidateExtension(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".txt", ".json", ".yaml", ".yml":
		return true
	default:
		return false
	}
}

func ReadFile(filePath string) ([]models.CuentaBancaria, error) {
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".txt":
		return fileTXT(filePath)
	case ".json":
		return fileJSON(filePath)
	default:
		return nil, fmt.Errorf("Archivo no valido: %s", ext)
	}
}

func fileJSON(filePath string) ([]models.CuentaBancaria, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data []models.CuentaBancaria
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func fileTXT(filePath string) ([]models.CuentaBancaria, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var cuentas []models.CuentaBancaria

	for _, record := range records {

		saldoDisponible, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			return nil, fmt.Errorf("error al convertir el saldo: %v", err)
		}
		fechaActualizacion, err := time.Parse("2006-01-02", record[5])
		if err != nil {
			return nil, fmt.Errorf("error al convertir  fecha actualizacion: %v", err)
		}
		fechaCreacion, err := time.Parse("2006-01-02", record[6])
		if err != nil {
			return nil, fmt.Errorf("error al convertir fecha creacion: %v", err)
		}

		cuenta := models.CuentaBancaria{
			Cedula:             record[0],
			NombreCompleto:     record[1],
			NumeroCuenta:       record[2],
			EntidadCuenta:      record[3],
			SaldoDisponible:    saldoDisponible,
			FechaActualizacion: fechaActualizacion,
			FechaCreacion:      fechaCreacion,
			EstadoCuenta:       record[7],
		}
		cuentas = append(cuentas, cuenta)
	}
	return cuentas, nil
}

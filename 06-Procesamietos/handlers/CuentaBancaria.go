package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"procesamientos/db"
	"procesamientos/models"
	"procesamientos/util"
	"strconv"
)

func GetCuentasBancarias(rw http.ResponseWriter, req *http.Request) {

	pageStr := req.URL.Query().Get("page")
	pageSizeStr := req.URL.Query().Get("page_size")

	// Convertir parámetros a enteros
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	var cuentas []models.CuentaBancaria
	result := db.Database().Limit(pageSize).Offset(offset).Find(&cuentas)
	if result.Error != nil {
		http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(cuentas)
}

func GetFile(rw http.ResponseWriter, req *http.Request) {
	// Se establece y valida el peso del archivo
	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(rw, "Peso no soportado", http.StatusBadRequest)
		return
	}

	// Obtener el archivo desde el formulario
	file, handler, err := req.FormFile("file")
	if err != nil {
		http.Error(rw, "Error al recuperar el archivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validar la extensión del archivo
	if !util.ValidateExtension(handler.Filename) {
		http.Error(rw, "Archivo no soportado", http.StatusBadRequest)
		return
	}

	cuentas, err := util.ReadFile(handler.Filename)
	if err != nil {
		fmt.Printf("Error al leer el archivo: %v\n", err)
		return
	}

	for _, cuenta := range cuentas {
		fmt.Printf("Cuenta Bancaria:\n")
		fmt.Printf("  Cedula: %s\n", cuenta.Cedula)
		fmt.Printf("  Nombre Completo: %s\n", cuenta.NombreCompleto)
		fmt.Printf("  Numero de Cuenta: %s\n", cuenta.NumeroCuenta)
		fmt.Printf("  Entidad de la Cuenta: %s\n", cuenta.EntidadCuenta)
		fmt.Printf("  Saldo Disponible: %.2f\n", cuenta.SaldoDisponible)
		fmt.Printf("  Fecha de Actualizacion: %s\n", cuenta.FechaActualizacion.Format("2006-01-02"))
		fmt.Printf("  Fecha de Creacion: %s\n", cuenta.FechaCreacion.Format("2006-01-02"))
		fmt.Printf("  Estado de la Cuenta: %s\n", cuenta.EstadoCuenta)
		fmt.Println()
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(cuentas)

}

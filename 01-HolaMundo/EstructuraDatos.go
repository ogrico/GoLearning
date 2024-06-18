package main

import "fmt"

func main() {
	var matriz = [...]int{1, 2, 3, 4, 5}
	fmt.Println(matriz)
	for _, value := range matriz {
		fmt.Println("Valor: ", value)
	}
	var matriz2 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matriz2)

	//Declaraciín e incializacion de una rebanada
	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}

	// Crear una rebanada apartir de otra
	diasRebanada := diasSemana[0:5]

	//Motrar tambien la longitud y la capacidad
	fmt.Println(diasRebanada)
	fmt.Println("Longitud: ", len(diasRebanada))
	fmt.Println("Capacidad: ", cap(diasRebanada))

	//Declaración e inicialización de una mapa
	colors := map[string]string{
		"rojo":   "#FF0000",
		"verder": "#00FF00",
		"azul":   "#0000FF",
	}

	fmt.Println(colors["verde"])

	type Persona struct {
		Nombre            string
		Edad              int
		CorreoElectronico string
	}
	var p Persona
	p.Nombre = "Juan"
	p.Edad = 30
	p.CorreoElectronico = "juan@example.com"
}

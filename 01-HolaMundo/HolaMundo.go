package main

import (
	"fmt"
	"runtime"
	"time"
)

func hora() {
	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("Mañana")
	} else if t.Hour() < 17 {
		fmt.Println("Tarde")
	} else {
		fmt.Println("Noche")
	}
}

func so() {
	os := runtime.GOOS
	switch os {
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	case "darwin":
		fmt.Println("Mac")
	default:
		fmt.Println("Otro OS")
	}
}

func main() {

	// Declaración de variables
	var (
		name, txt string
		age       int
	)

	name = "pedro"
	txt = "Prueba"
	age = 20

	// Constantes
	const Pi float64 = 3.1415926
	const (
		x = 100
		y = 0b1010
		z = 0o12
		w = 0xff
	)
	const (
		Lunes = iota + 1
		Martes
		Miercoles
		Jueves
		Viernes
		Sabado
		Domingo
	)
	fmt.Println("Hola Mundo desde Go")
	fmt.Println(name, txt, age)
	fmt.Println(x, y, z, w)
	fmt.Println(Lunes, Martes, Miercoles, Jueves, Viernes, Sabado, Domingo)

	// Comparación de números
	fmt.Println(1 == 1) // true
	fmt.Println(1 != 2) // true
	fmt.Println(2 < 3)  // true
	fmt.Println(3 > 4)  // false
	fmt.Println(4 <= 4) // true
	fmt.Println(5 >= 6) // false

	// Comparación de cadenas
	fmt.Println("hola" == "hola")  // true
	fmt.Println("hola" != "adios") // true
	fmt.Println("abc" < "def")     // true
	fmt.Println("ghi" > "fgh")     // true
	fmt.Println("hij" <= "hij")    // true
	fmt.Println("klm" >= "klmno")  // false

	// Comparación de booleanos
	fmt.Println(true == true)           // true
	fmt.Println(false != true)          // true
	fmt.Println(true && false == false) // true
	fmt.Println(true || false == true)  // true

	hora()
	so()

}

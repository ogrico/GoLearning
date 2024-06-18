package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func validateZero(divisor int) {
	if divisor == 0 {
		panic("Cannot divide by zero")
	}
}

func panico(a, b int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	validateZero(b)
	fmt.Println(a / b)
}

func division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	var (
		srt      = "123"
		num, err = strconv.Atoi(srt)
	)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Numero: ", num, "\n")
	fmt.Println(division(num, 3))
	// Defer
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString("Hello World")
	if err != nil {
		fmt.Println("Error: ", err, "\n")
		return
	}
	panico(100, 0)
	panico(10, 2)
	panico(40, 3)
	panico(70, 0)

	logTxt, e := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if e != nil {
		log.Fatal(e)
	}

	defer logTxt.Close()

	log.SetOutput(logTxt)
	log.Print("¡Oye, soy un Log!")
}

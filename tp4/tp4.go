package tp4

import (

	"fmt"
	"lab3-tps/tp4/calc"
)

func Tp4(){
	calc.SecretSwitch = true
	var num1 float64
	var num2 float64
	var op, usr string
	fmt.Printf("Hola! Bienvenido a GoCalc, por favor ingresa tu nombre: ")
	fmt.Scanln(&usr)
	fmt.Printf("Por favor ingresa un número: ")
	fmt.Scanln(&num1)
	fmt.Printf("Ingresa otro número: ")
	fmt.Scanln(&num2)
	fmt.Printf("Ahora ingresa la operación deseada (+, -, /, *): ")
	fmt.Scanln(&op)
	log := calc.Operate(num1, num2, usr, op)
	fmt.Printf("Results: \n%+v",log)

}





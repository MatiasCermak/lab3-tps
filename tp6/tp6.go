package tp6

import (
	"fmt"
	"lab3-tps/tp6/cCalc"
)

func Tp6(){
	ch := make(chan cCalc.OpLog)
	cCalc.SecretSwitch = true
	var num1 float64
	var num2 float64
	var usr string
	fmt.Printf("Hola! Bienvenido a GoCalc, por favor ingresa tu nombre: ")
	fmt.Scanln(&usr)
	fmt.Printf("Por favor ingresa un número: ")
	fmt.Scanln(&num1)
	fmt.Printf("Ingresa otro número: ")
	fmt.Scanln(&num2)
	cCalc.Operate(num1, num2, usr, ch)
	for i := 0; i < 4; i++ {
		fmt.Printf("%+v \n", <- ch)
	}

}





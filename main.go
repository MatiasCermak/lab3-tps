package main

import (
	"fmt"
	"lab3-tps/tp3"
	"lab3-tps/tp4"
	"lab3-tps/tp5"
	"lab3-tps/tp6"
	"lab3-tps/tp8"
)

func main(){
	fmt.Println("Trabajos Prácticos para la materia LABIII de la Universidad Blas Pascal.\nPor favor, ingrese un número del 1 al 5 para seleccionar el TP a ejecutar\n")
	var a int8
	fmt.Scanln(&a)
	switch a {
	case 1:
		tp3.Tp3()
		break
	case 2:
		tp4.Tp4()
		break
	case 3:
		tp5.Tp5()
		break
	case 4:
		tp6.Tp6()
		break
	case 5:
		tp8.Tp8()
		break
	default:
		println("No se ha seleccionado ningún numero")

	}
}
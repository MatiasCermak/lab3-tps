package cCalc

import "errors"

var SecretSwitch bool = false

type OpLog struct {
	Number1      float64
	Number2      float64
	Result       float64
	Username     string
	Operation    string
	SecretSwitch bool
	Error        error
}

func Add(n1, n2 float64, usr string, ch chan OpLog){
	if(SecretSwitch == true){
		n2+=5
	}
	ch <- OpLog{
		Number1:      n1,
		Number2:      n2,
		Result:       n1+n2,
		Username:     usr,
		Operation:    "Addition",
		SecretSwitch: SecretSwitch,
		Error: nil,
	}
}

func Subtract(n1, n2 float64, usr string, ch chan OpLog){
	if(SecretSwitch == true){
		n2+=5
	}
	ch <- OpLog{
		Number1:      n1,
		Number2:      n2,
		Result:       n1-n2,
		Username:     usr,
		Operation:    "Subtraction",
		SecretSwitch: SecretSwitch,
		Error: nil,
	}
}

func Divide(n1, n2 float64, usr string, ch chan OpLog){

	if(SecretSwitch == true){
		n1+=5
	}

	if n2 == 0 {
		ch <- OpLog{
			Number1:      n1,
			Number2:      n2,
			Result:       -1,
			Username:     usr,
			Operation:    "Division",
			SecretSwitch: SecretSwitch,
			Error:        errors.New("The divisor can't be 0"),
		}
		return
	}

	ch <- OpLog{
		Number1:      n1,
		Number2:      n2,
		Result:       n1/n2,
		Username:     usr,
		Operation:    "Division",
		SecretSwitch: SecretSwitch,
		Error: nil,
	}
}

func Multiply(n1, n2 float64, usr string, ch chan OpLog) {
	if(SecretSwitch == true){
		n2+=5
	}
	ch <- OpLog{
		Number1:      n1,
		Number2:      n2,
		Result:       n1*n2,
		Username:     usr,
		Operation:    "Multiplication",
		SecretSwitch: SecretSwitch,
		Error: nil,
	}
}

func Operate(n1, n2 float64, usr string, ch chan OpLog){
	go Add(n1, n2, usr, ch)
	go Subtract(n1, n2, usr, ch)
	go Multiply(n1, n2, usr, ch)
	go Divide(n1, n2, usr, ch)
}

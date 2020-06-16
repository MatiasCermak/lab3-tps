package cCalc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

)

//Using BDD

func TestDivide(t *testing.T) {
	//Given Two numbers 2, 5
	a := 2.0
	b := 5.0

	//When I divide 2/5
	ch := make(chan OpLog)
	go Divide(a,b,"TestUsr", ch)

	result := <- ch
	//Then The result should equal 2/5
	assert.Equal(t, 0.4, result.Result, fmt.Sprintf("The division is not correct."))
	//And There should be no error
	assert.Nil(t, result.Error, fmt.Sprintf("Error not nil."))
}

func TestAdd(t *testing.T){
	//Given Two numbers 1, 4
	a := 1.0
	b := 4.0

	//When I add both numbers
	ch := make(chan OpLog)
	go Add(a,b,"TestUsr", ch)
	result := <- ch

	//Then the result should equal 5
	assert.Equal(t, float64(5), result.Result, fmt.Sprintf("The sum is not correct."))
}

func TestDivide2(t *testing.T){
	//Given Two numbers 2, 0
	a := 2.0
	b := 0.0

	//When I divide 2/0
	ch := make(chan OpLog)
	go Divide(a,b,"TestUsr", ch)
	result := <- ch

	//Then The result should equal -1
	assert.Equal(t, float64(-1), result.Result, fmt.Sprintf("The division is not correct."))
	//And There should be an error
	assert.NotNil(t, result.Error, fmt.Sprintf("Error should not be nil."))
	//And The error message should read "The divisor can't be 0"
	assert.Equal(t, "The divisor can't be 0", result.Error.Error(), fmt.Sprintf("Error message should be \"The divisor can't be 0\""))
}

func TestSubtract(t *testing.T){
	//Given Two numbers 1, 4
	a := 1.0
	b := 4.0

	//When I subtract 1 - 4
	ch := make(chan OpLog)
	go Subtract(a,b,"TestUsr", ch)
	result := <- ch

	//Then the result should equal -3
	assert.Equal(t, float64(-3), result.Result, fmt.Sprintf("The subtraction is not correct."))
}

func TestMultiply(t *testing.T){
	//Given Two numbers 1, 4
	a := 5.0
	b := 4.0

	//When I multiply both numbers
	ch := make(chan OpLog)
	go Multiply(a,b,"TestUsr", ch)
	result := <- ch

	//Then the result should equal 20
	assert.Equal(t, float64(20), result.Result, fmt.Sprintf("The multiplication is not correct."))
}
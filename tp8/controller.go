package tp7

import (
	"github.com/gin-gonic/gin"
	"strconv"
)


func sum(c *gin.Context){
	a := c.Query("a")
	b := c.Query("b")
	if(a == "" || b == ""){
		c.String(400, "Missing Params a, b")
		return
	}
	A, err1:= strconv.ParseFloat(a, 64)
	B, err2:= strconv.ParseFloat(b, 64)
	if(err1 != nil || err2 !=nil){
		c.String(400, "Invalid Params a, b")
		return
	}
	c.String(200, "{\"result\": %v}", A+B)
}

func res(c *gin.Context){
	a := c.Query("a")
	b := c.Query("b")
	if(a == "" || b == ""){
		c.String(400, "Missing Params a, b")
		return
	}
	A, err1:= strconv.ParseFloat(a, 64)
	B, err2:= strconv.ParseFloat(b, 64)
	if(err1 != nil || err2 !=nil){
		c.String(400, "Invalid Params a, b")
		return
	}
	c.String(200, "{\"result\": %v}", A-B)
}

func muli(c *gin.Context){
	a := c.Query("a")
	b := c.Query("b")
	if(a == "" || b == ""){
		c.String(400, "Missing Params a, b")
		return
	}
	A, err1:= strconv.ParseFloat(a, 64)
	B, err2:= strconv.ParseFloat(b, 64)
	if(err1 != nil || err2 !=nil){
		c.String(400, "Invalid Params a, b")
		return
	}
	c.String(200, "{\"result\": %v}", A*B)
}

func div(c *gin.Context){
	a := c.Query("a")
	b := c.Query("b")
	if(a == "" || b == ""){
		c.String(400, "Missing Params a, b")
		return
	}
	A, err1:= strconv.ParseFloat(a, 64)
	B, err2:= strconv.ParseFloat(b, 64)
	if(err1 != nil || err2 !=nil){
		c.String(400, "Invalid Params a, b")
		return
	}
	if(B==0){
		c.String(400, "Invalid param b, cannot be 0 in div")
		return
	}
	c.String(200, "{\"result\": %v}", A/B)
}



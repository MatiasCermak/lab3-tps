package tp8

import (
	"github.com/gin-gonic/gin"
	)

func Tp8(){
	r := gin.Default()

	r.GET("/calc/sum", sum)
	r.GET("/calc/muli", muli)
	r.GET("/calc/res", res)
	r.GET("/calc/div", div)

	r.Run()
}
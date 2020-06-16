package tp5

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Item struct  {
	Name               string `json:"name"`
	CountryId          string `json:"country_id"`
	Currencies[] struct{
		Id     string
		Symbol string
	} `json:"currencies"`

}

func Tp5() {
	val, st := getReqToJson("https://api.mercadolibre.com/sites/MLA")
	fmt.Println(st)
	fmt.Printf("Value: %+v", val)
    r := gin.Default()
    r.GET("/answer", func(c *gin.Context) {
		c.JSON(st, val)
	} )

    r.Run()

}

func getReqToJson(url string) (Item, int) {
	req, erro := http.Get(url)
	if erro != nil {
		fmt.Println(erro)
	}
	if req.StatusCode != 200 {
		fmt.Println("Error: ", erro)
		return Item{}, req.StatusCode
	}
	defer req.Body.Close()
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		println("There has been an error reading the body of the response")
	}
	var res Item
	//fmt.Println(string(data))
	er := json.Unmarshal(data, &res)
	if er != nil {
		println("There has been an error with the Unmarshal of the json")
		fmt.Println(er)
	}

	return res, req.StatusCode
}


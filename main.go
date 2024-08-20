package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"golang.org/x/text/language"
)

func main() {
	fmt.Println("Starting vulnerable Go application...")

	// Using a vulnerable version of gin
	r := gin.Default()

	r.GET("/parse", func(c *gin.Context) {
		json := `{"name": {"first": "Janet", "last": "Prichard"}, "age": 47}`

		// Using a vulnerable version of gjson
		value := gjson.Get(json, "name.last")
		
		c.String(http.StatusOK, "Last name: "+value.String())
	})

	// Using a vulnerable version of golang.org/x/text
	tag := language.Make("en-US")
	fmt.Printf("Language tag: %s\n", tag)

	r.Run(":8080")
}
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Response from the api defined
type Response []struct {
	Q string `json:"q"`
	A string `json:"a"`
	H string `json:"h"`
}

// Struct to define quote
type Quote struct {
	ID    int    `json:"id"`
	Quote string `json:"title"`
}

// Define the constant url of the API
const url_api string = "https://zenquotes.io/api/random"

func RequestQuote() string {

	var result Response

	resp, err := http.Get(url_api)

	if err != nil {
		fmt.Println("No response from request")
	}
	body, _ := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		log.Fatalln(err)
	}

	// read json data into a Result struct
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	// return the field of the Struct -> quote = 'Q'
	return result[0].Q

}

func BuildQuotes() []Quote {

	var quotes []Quote

	for i := 1; i <= 3; i++ {

		var quote Quote
		quote.ID = i
		quote.Quote = RequestQuote()

		quotes = append(quotes, quote)

		// to avoid crashing the connection
		time.Sleep(1 * time.Second)

	}

	return quotes

}

// getAlbums responds with the list of all albums as JSON.
func getQuotes(c *gin.Context) {
	quotes := BuildQuotes()
	b, _ := json.Marshal(quotes)
	c.JSON(http.StatusOK, string(b))
}

func main() {

	quote := BuildQuotes()

	fmt.Println(quote)

	router := gin.Default()
	router.GET("/quote", getQuotes)

	router.Run("localhost:8080")
}

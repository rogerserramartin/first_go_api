package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
	"fmt"
)
*/

// field names MUST start with a capital letter, because that will turn that field into an EXPORTABLE field name
// capital letter is like public in java
type character struct {
	ID        string `json:"id"` // when we are serializing with JSON, convert ID into lowercase
	Name      string `json:"name"`
	Franchise string `json:"franchise"`
}

var characters = []character{
	{ID: "1", Name: "Alternis Dim", Franchise: "Bravely Default"},
	{ID: "2", Name: "Goku", Franchise: "Dragon Ball"},
	{ID: "3", Name: "Sakura", Franchise: "Card Captor Sakura"},
	{ID: "4", Name: "Jack", Franchise: "Stranger of Paradise"},
}

// get request
// we need to accept a context, a pointer to it
// a gin context, is essentially all the information about the request that allows us to return a response
func getCharacters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, characters) // STATUS OK is code 200
}

func createCharacter(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, characters) // STATUS OK is code 200
}

func main() {
	/*
		// this also works instead of var characters = [], but needs to be inside a method
		char := character{ID: "1", Name: "Monkey D Luffy", Franchise: "One Piece"}
		chars := []character{
			{ID: "1", Name: "Monkey D Luffy", Franchise: "One Piece"},
			{ID: "2", Name: "Monkey D Garp", Franchise: "One Piece"},
			{ID: "3", Name: "Monkey D Dragon", Franchise: "One Piece"},
		}
	*/
	router := gin.Default() // router comes from gin dependency
	// we can route a specific route to a function
	router.GET("/api/v1/characters", getCharacters)
	router.Run("localhost:8080")

}

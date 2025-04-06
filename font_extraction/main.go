// // main.go
// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()

// 	r.GET("/fonts", func(c *gin.Context) {
// 		fonts := []string{"Roboto", "Poppins", "Lato", "Open Sans", "Montserrat"}
// 		c.JSON(http.StatusOK, fonts)
// 	})

// 	r.Run(":8080")
// }

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type Font struct {
	Family string `json:"family"`
}

type GoogleFontsResponse struct {
	Items []Font `json:"items"`
}

func fontsHandler(c *gin.Context) {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	url := fmt.Sprintf("https://www.googleapis.com/webfonts/v1/webfonts?key=%s", apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("❌ Error fetching fonts:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch fonts"})
		return
	}
	defer resp.Body.Close()

	// 🔍 Debug: log raw response
	bodyBytes, _ := io.ReadAll(resp.Body)
	// fmt.Println("📦 Raw Response:", string(bodyBytes))

	// Now decode the JSON
	var data GoogleFontsResponse
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		fmt.Println("❌ JSON Decode Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode fonts"})
		return
	}

	var fontNames []string
	for _, font := range data.Items {
		fontNames = append(fontNames, font.Family)
	}

	// // ✅ Debug print the number of fonts and a few names
	// fmt.Printf("✅ Retrieved %d fonts\n", len(fontNames))
	// if len(fontNames) > 0 {
	// 	fmt.Println("Sample fonts:", fontNames[:5]) // show first 5
	// }

	c.JSON(http.StatusOK, fontNames)
}

// func main() {
// 	r := gin.Default()
// 	r.GET("/fonts", fontsHandler)
// 	r.Run(":8080")
// }

func main() {
	// Load env variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	port := os.Getenv("PORT")

	r := gin.Default()
	r.GET("/fonts", fontsHandler)
	r.Run(":" + port)
}

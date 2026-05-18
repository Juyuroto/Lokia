package services

import (
	"log"
	"os"

	"github.com/diverged/tavily-go/models"
	tavilygo "github.com/diverged/tavily-go"
)

func SearchWeb(query string) string {
	keys := []string {
		os.Getenv("TAVILY_KEY_1"),
		os.Getenv("TAVILY_KEY_2"),
		os.Getenv("TAVILY_KEY_3"),
	}

	for _, key := range keys {
		if key == "" {
			continue
		}

		client := tavilygo.NewClient(key)

		searchReq := models.SearchRequest{
           	Query:       query,
           	SearchDepth: "basic",
       	}

        response, err := tavilygo.Search(client, searchReq)
        if err != nil{
        	log.Printf("Clé épuisée ou erreur, tentative avec la suivante: %v", err)
         	continue
        }

        var content string
        for _, result := range response.Results {
            content += result.Content + "\n"
        }
        return content

	}

 	log.Println("Toutes les clés Tavily sont épuisées")
    return ""
}
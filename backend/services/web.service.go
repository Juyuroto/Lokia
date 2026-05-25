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
	}
	
 	log.Println("Toutes les clés Tavily sont épuisées")
    return ""
    
}
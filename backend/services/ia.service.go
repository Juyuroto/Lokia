package services

import (
	"context"
	"log"
	"fmt"
	
	"github.com/ollama/ollama/api"
)

func LocalIA() {
	client, err := api.ClientFromEnvironment()
		if err != nil {
			log.Fatal(err)
		}

	ctx := context.Background()

	req := &api.GenerateRequest{
		Model:  "llama3.2-vision:11b",
		Prompt: "Say hello to me",
	}
	
	var response string
	err = client.Generate(ctx, req, func(resp api.GenerateResponse) error {
		response += resp.Response
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
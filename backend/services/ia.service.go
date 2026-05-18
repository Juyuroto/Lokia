package services

import (
	"context"
	"fmt"
	
	"github.com/ollama/ollama/api"
)

const model = "llama3.2-vision:11b"

var IAClient *api.Client

func LoadIA(){
	client, err := api.ClientFromEnvironment()
	if err != nil {
		panic(err)
	}
	IAClient = client

	fmt.Println("IA chargée avec le modèle :", model)
}

func AskIA(prompt string, history []api.Message) (string, error){

	if IAClient == nil {
	    return "", fmt.Errorf("IA non initialisée")
	}

	ctx := context.Background()

	history = append(history, api.Message{
	    Role:    "user",
	    Content: prompt,
	})

	req := &api.ChatRequest{
		Model:  model,
		Messages: history,
	}

	var result string
	err := IAClient.Chat(ctx, req, func(resp api.ChatResponse) error {
		result += resp.Message.Content
		return nil
	})

	return result, err
}
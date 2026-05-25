package services

import (
	"context"
	"fmt"
	
	"github.com/ollama/ollama/api"
)

const model = "llama3.2-vision:11b"

var IAClient *api.Client

func LoadIA(){

}

func AskIA(prompt string, history []api.Message) (string, error){

}
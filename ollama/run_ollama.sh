#!/bin/bash
echo "Starting Ollama server..."
ollama serve & sleep 5
echo "Pulling models"
ollama pull nomic-embed-text:latest && ollama pull llama3.2-vision:11b && wait
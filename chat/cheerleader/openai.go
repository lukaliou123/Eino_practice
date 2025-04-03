package main

import (
	"context"
	"github.com/cloudwego/eino/components/model"
	"log"
	"os"

	"github.com/cloudwego/eino-ext/components/model/openai"
)

func createOpenAIChatModel(ctx context.Context) model.ChatModel {
	key := os.Getenv("OPENAI_API_KEY")
	modelName := os.Getenv("OPENAI_MODEL_NAME")
	baseURL := os.Getenv("OPENAI_BASE_URL")
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		BaseURL: baseURL,
		Model:   modelName,
		APIKey:  key,
	})
	if err != nil {
		log.Fatalf("create openai chat model failed, err=%v", err)
	}
	return chatModel
}

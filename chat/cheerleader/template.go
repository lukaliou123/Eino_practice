package main

import (
	"context"
	"log"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func createTemplate() prompt.ChatTemplate {
	// 创建模板，使用 FString 格式
	ChatTemplate := prompt.FromMessages(schema.FString,
		// 系统消息模板
		schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。你的目标是帮助程序员保持积极乐观的心态，提供技术建议的同时也要关注他们的心理健康。"),

		// 插入需要的对话历史（新对话的话这里不填）
		schema.MessagesPlaceholder("chat_history", true),
		// 用户消息模板
		schema.UserMessage("问题: {question}"),
	)
	return ChatTemplate
}

func createMessagesFromTemplate(role string, style string, question string, chatHistory []*schema.Message) []*schema.Message {
	template := createTemplate()

	// 使用模板生成消息
	message, err := template.Format(context.Background(), map[string]any{
		"role":         "程序员鼓励师",
		"style":        "积极、温暖且专业",
		"question":     "我今天很沮丧，一直和家里人起矛盾，怎么办？",
		"chat_history": chatHistory,
	})
	if err != nil {
		log.Fatalf("format template failed: %v", err)
	}
	return message
}

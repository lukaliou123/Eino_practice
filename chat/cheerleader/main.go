package main

import (
	"context"
	"github.com/cloudwego/eino/schema"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	ctx := context.Background()
	// 使用template生成一个messages
	log.Printf("===create messages===\n")
	chatHistory := []*schema.Message{
		schema.UserMessage("你好"),
		schema.AssistantMessage("嘿，我是你的鼓励师，每个人都会遇到家庭矛盾，不是吗？", nil),
		schema.UserMessage("我尝试了，但是他们不听我的。"),
		schema.AssistantMessage("不要沮丧，你可以和家里人好好沟通，听听他们的想法，找到解决问题的方法。", nil),
	}
	messages := createMessagesFromTemplate("家庭矛盾鼓励师", "积极、温暖且专业", "我今天很沮丧，一直和家里人起矛盾，怎么办？", chatHistory)
	log.Printf("messages: %v\n\n", messages)
	// 创建llm，这里选择ollama
	log.Printf("===create llm===\n")
	cm := createOllamaChatModel(ctx)
	log.Printf("===llm generate===\n")
	result := generate(ctx, cm, messages)
	log.Printf("result: %+v\n\n", result)

	log.Printf("===llm stream generate===\n")
	streamResult := stream(ctx, cm, messages)
	reportStream(streamResult)
}

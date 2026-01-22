package gemini

import (
	"context"
	"fmt"
	"os"

	genai "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GenerateTreino(prompt string) (string, error) {
	ctx := context.Background()
    apiKey := os.Getenv("GEMINI_API_KEY")

    // Adicionando o Endpoint da v1beta explicitamente
    client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
if err != nil {
    return "", err
}
defer client.Close()

    
    
/*     // Opcional: Adicione um log aqui para debugar se necessário
    // fmt.Println("Chamando modelo Gemini 1.5 Flash...")

iter := client.ListModels(ctx)
for {
    m, err := iter.Next()
    if err == iterator.Done {
        break
    }
    if err != nil {
        fmt.Println("Erro ao listar modelos:", err)
        break
    }
    fmt.Println("Modelo disponível:", m.Name)
} */
model := client.GenerativeModel("gemini-2.5-flash")

    resp, err := model.GenerateContent(ctx, genai.Text(prompt))
    if err != nil {
        return "", fmt.Errorf("erro ao gerar conteúdo: %v", err)
	}
	// 🔐 Validações obrigatórias
	if resp == nil || len(resp.Candidates) == 0 {
		return "", fmt.Errorf("gemini não retornou resposta")
	}

	parts := resp.Candidates[0].Content.Parts
	if len(parts) == 0 {
		return "", fmt.Errorf("gemini retornou resposta vazia")
	}

	text, ok := parts[0].(genai.Text)
	if !ok {
		return "", fmt.Errorf("resposta da gemini não é texto")
	}

	return string(text), nil
}

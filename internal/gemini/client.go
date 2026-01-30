package gemini

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func GenerateTreino(prompt string) (string, error) {
	ctx := context.Background()

	// 1. Carrega as configurações do arquivo JSON baixado do Google Cloud
	data, err := os.ReadFile("client_secret.json")
	if err != nil {
		log.Fatal(err)
	}

	// 2. Define o escopo necessário para o Gemini (Generative Language API)
	config, err := google.ConfigFromJSON(data, "https://www.googleapis.com/auth/generative-language")
	if err != nil {
		log.Fatal(err)
	}

	// 3. Obter o Token (Simulação de fluxo Desktop)
	// Em uma API Web, você redirecionaria o usuário para authURL
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Acesse o link para autorizar: \n%v\n", authURL)

	var authCode string
	fmt.Print("Digite o código de autorização: ")
	fmt.Scan(&authCode)

	token, err := config.Exchange(ctx, authCode)
	if err != nil {
		log.Fatal(err)
	}

	// 4. Inicializa o cliente Gemini usando o Token OAuth2 em vez de API Key
	client, err := genai.NewClient(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

model := client.GenerativeModel("gemini-.5-flash")

    resp, err := model.GenerateContent(ctx, genai.Text(prompt))
    if err != nil {
        return "", fmt.Errorf("erro ao gerar conteúdo: %v", err)
	}

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

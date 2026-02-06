package DB

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type FirebaseDB struct {
	App       *firebase.App
	Auth      *auth.Client
	Firestore *firestore.Client
}

func ConnectDB() (*FirebaseDB, error) {
    ctx := context.Background()

    // 1. Defina explicitamente o ID do seu projeto (veja no console do Firebase)
    projectID := "gen-lang-client-0189356877" // Verifique se este é o ID correto do seu projeto 'goback'

    conf := &firebase.Config{ProjectID: projectID}
    
    // 2. Se você estiver usando API Key para simplificar:
    opt := option.WithAPIKey("GOOGLE_APPLICATION_CREDENTIALS") 
    // Ou continue usando o JSON se preferir, mas garanta que o JSON seja desse projeto
    
    app, err := firebase.NewApp(ctx, conf, opt)
    if err != nil {
        return nil, err
    }

    // 3. AQUI ESTÁ O TRUQUE: O Firestore por padrão busca o banco "(default)".
    // Como o seu se chama "goback", precisamos inicializar o cliente manualmente ou renomear.
    // Se quiser usar o cliente do SDK do Firebase:
    firestoreClient, err := firestore.NewClientWithDatabase(ctx, projectID, "goback", opt)
    if err != nil {
        return nil, fmt.Errorf("erro ao conectar no banco goback: %v", err)
    }

    authClient, err := app.Auth(ctx)
    if err != nil {
        return nil, err
    }

    fmt.Println("🔥 Conectado com sucesso ao banco 'goback'!")

    return &FirebaseDB{
        App:       app,
        Auth:      authClient,
        Firestore: firestoreClient,
    }, nil
}
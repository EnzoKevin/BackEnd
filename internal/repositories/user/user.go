package user

import (
	model "BACKEND/internal/models"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cloud.google.com/go/firestore"

	"time"
)

type UserRepo struct {
	db *firestore.Client
}

func NewUserRepository(db *firestore.Client) *UserRepo {
	return &UserRepo{db: db}
}



func (r *UserRepo) GetAll() ([]model.User, error) {
	ctx := context.Background()

	docs, err := r.db.Collection("users").Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var users []model.User

	fmt.Printf("🔍 Documentos encontrados no Firestore: %d\n", len(docs)) // Log de debug
	for _, doc := range docs {
		var u model.User
		if err := doc.DataTo(&u); err != nil {
			return nil, err
		}

		u.ID = doc.Ref.ID
		users = append(users, u)
	}

	return users, nil
}



func (r *UserRepo) Add(user model.User) (string, error) {
	ctx := context.Background()

	doc, _, err := r.db.Collection("users").Add(ctx, map[string]interface{}{
		"name": user.Name,
		"email":     user.Email,
		"password":  user.Password,
		"height":    user.Height,
		"weight":    user.Weight,
		"btype":    user.BType,
		"target":    user.Target,
		
	
	})

	if err != nil {
		return "", err
	}

	return doc.ID, nil
}



func (r *UserRepo) AddTreino(user model.CreateTreino) (string, error) {
    ctx := context.Background()

    if user.ID == "" {
        return "", fmt.Errorf("ID do usuário é obrigatório para salvar o treino")
    }
    
    docRef, _, err := r.db.Collection("users").
        Doc(user.ID).
        Collection("Train").
        Add(ctx, map[string]interface{}{
		"ID": user.ID,
		"Time": time.Now(),
		"Segunda":     user.Segunda,
		"Terca":  user.Terca,
		"Quarta":    user.Quarta,
		"Quinta":    user.Quinta,
		"Sexta":    user.Sexta,
		"Sabado":    user.Sabado,
		"Domingo":    user.Domingo,
		
	
	})

    if err != nil {
        return "", fmt.Errorf("falha ao salvar no banco goback: %v", err)
    }

    return docRef.ID, nil
}

func (r *UserRepo) GetTreino(userID string) (*model.CreateTreino, error) {
    ctx := context.Background()
    
    docRef := r.db.Collection("users").
        Doc(userID).
        Collection("Train").OrderBy("Time", firestore.Desc).Limit(1).Documents(ctx)

	defer docRef.Stop()

    doc, err := docRef.Next()
    if err != nil {
        return nil, fmt.Errorf("erro ao obter documento: %v", err)
    }

		var u model.CreateTreino
	
    if err := doc.DataTo(&u); err != nil {
        return nil, fmt.Errorf("erro ao converter dados: %v", err)
    }
		u.ID = doc.Ref.ID

    return &u, nil
}

func (r *UserRepo) GetByID(id string) (*model.User, error) {
	ctx := context.Background()

	doc, err := r.db.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil
		}
		return nil, err
	}

	var u model.User
	if err := doc.DataTo(&u); err != nil {
		return nil, err
	}

	u.ID = doc.Ref.ID
	return &u, nil
}



func (r *UserRepo) DeleteUser(id string) error {
	ctx := context.Background()

	_, err := r.db.Collection("users").Doc(id).Delete(ctx)
	return err
}



func (r *UserRepo) EmailExists(email string) (bool, error) {
	ctx := context.Background()

	iter := r.db.
		Collection("users").
		Where("email", "==", email).
		Limit(1).
		Documents(ctx)

	docs, err := iter.GetAll()
	if err != nil {
		return false, err
	}

	return len(docs) > 0, nil
}

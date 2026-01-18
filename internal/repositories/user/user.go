package user

import (
	"context"

	model "BACKEND/internal/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cloud.google.com/go/firestore"
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
	})

	if err != nil {
		return "", err
	}

	return doc.ID, nil
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

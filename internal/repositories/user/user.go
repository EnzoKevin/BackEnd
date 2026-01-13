package user

/*
import (
	model "BACKEND/internal/models"
)

type Users struct {
	users []model.User
}

func New() *Users {
	return &Users{
		users: make([]model.User, 0),
	}
}

func (u *Users) GetAll() []model.User {
	return u.users
}

func (u *Users) GetByID(id string) *model.User {
	for _, v := range u.users {
		if v.ID.String() == id {
			return &v
		}
	}
	return nil
}

func (u *Users) DeleteUser(id string) bool {
	for i, v := range u.users {
		if v.ID.String() == id {
			u.users = append(u.users[:i], u.users[i+1:]...)
			return true
		}
	}
	return false
}

func (u *Users) Add(newUser model.User) {
	u.users = append(u.users, newUser)
}

func (u *Users) EmailExists(email string) bool {
	for _, v := range u.users {
		if v.Email == email {
			return true
		}
	}
	return false
} */

import (
	model "BACKEND/internal/models"
	"database/sql"
)

type UserRepo struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) *UserRepo {
	return &UserRepo{connection: connection}
}

func New() *UserRepo {
	return &UserRepo{connection: nil}
}


func (pr *UserRepo) GetAll() ([]model.User, error) {
	query := "SELECT id, user_name, email FROM users"
	rows, err := pr.connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (pr *UserRepo) Add(user model.User) (int, error) {
	var id int
	query := `
	INSERT INTO users (user_name, email, password)
	OUTPUT INSERTED.id
	VALUES (@p1, @p2, @p3)
	`
	err := pr.connection.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Password,
	).Scan(&id)

	return id, err
}

func (pr *UserRepo) GetByID(id string) (*model.User, error) {
	query := "SELECT id, user_name, email FROM users WHERE id = @1"

	var u model.User
	err := pr.connection.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (pr *UserRepo) DeleteUser(id string) bool {
	query := "DELETE FROM users WHERE id = @1"
	_, err := pr.connection.Exec(query, id)

	if err != nil {
		return false
	}

	return true
}

func (pr *UserRepo) EmailExists(email string) (bool, error) {
	query := "SELECT COUNT(1) FROM users WHERE email = @1"
	var count int
	err := pr.connection.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
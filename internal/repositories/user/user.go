package user

/* import (
	model "BACKEND/internal/models"
)

type Users struct {
	users []model.User
}



func (u Users) GetAll() []model.User {
	return u.users
}



func (u *Users) Add(NewUser model.User) {
	u.users = append(u.users, NewUser)
}
*/

import (
	model "BACKEND/internal/models"
	"database/sql"
	"fmt"
)

type userRepo struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) *userRepo {
	return &userRepo{connection: connection}
}


func (u *userRepo) EmailExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)"
	err := u.connection.QueryRow(query, email).Scan(&exists)
	return exists, err
}



func (pr *userRepo) GetAll() ([]model.User, error) {

	query := "SELECT id, user_name, Email FROM user"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.User{}, err
	}

	var userList []model.User
	var userObj model.User

	for rows.Next(){
		err = rows.Scan(
			&userObj.ID,
			&userObj.Name,
			&userObj.Email)
		if err != nil {
			fmt.Println(err)
		return []model.User{}, err
		}

		userList = append(userList, userObj)
	}

	rows.Close()

	return userList, nil
}

func (u *userRepo) Add(user model.User) (int, error) {
	var id int
	query := `
	INSERT INTO users (user_name, email, password)
	VALUES ($1, $2, $3)
	RETURNING id
	`
	err := u.connection.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Password,
	).Scan(&id)
	return id, err
}


func (pr *userRepo) GetUserById(id_user int) (*model.User, error) {
	query, err := pr.connection.Prepare("SELECT * FROM user WHERE ID = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var user model.User

	err = query.QueryRow(id_user).Scan(
		&user.ID,
		&user.Name,
		&user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	query.Close()

	return &user, nil
}
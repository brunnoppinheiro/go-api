package repository

import (
	"database/sql"
	"fmt"
	"go-api/cmd/model"
)

// Interface que será usada no usecase
type UserRepositoryInterface interface {
	GetUsers() ([]model.User, error)
}

// Struct concreta com letra minúscula para encapsulamento
type userRepository struct {
	connection *sql.DB
}

// Função para instanciar o repositório, retornando a interface
func NewUserRepository(connection *sql.DB) UserRepositoryInterface {
	return &userRepository{
		connection: connection,
	}
}

// Implementação do método da interface
func (ur *userRepository) GetUsers() ([]model.User, error) {
	query := "SELECT user_id, user_name, user_age, user_email, user_created, user_phone, user_status FROM users"
	rows, err := ur.connection.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.User{}, err
	}
	defer rows.Close()

	var userlist []model.User

	for rows.Next() {
		var userobj model.User
		err := rows.Scan(
			&userobj.UserID,
			&userobj.UserName,
			&userobj.UserAge,
			&userobj.UserEmail,
			&userobj.UserCreated,
			&userobj.UserPhone,
			&userobj.UserStatus,
		)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return []model.User{}, err
		}

		userlist = append(userlist, userobj)
	}

	return userlist, nil
}

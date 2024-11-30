package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-embreo-assessment/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	IsUsernameExist(username string) (bool, error)
	IsEmailExist(email string) (bool, error)
	// GetUserByUsername(username string) (*models.User, error)
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) IsUsernameExist(username string) (bool, error) {
	query := `
		SELECT id, username, email, password FROM users where username = $1
	`

	user := &models.User{}

	err := r.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *userRepository) IsEmailExist(email string) (bool, error) {
	query := `
	SELECT id, username, email, password FROM users where email = $1
`

	user := &models.User{}

	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *userRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users(email, username, password) VALUES ($1, $2, $3) RETURNING id
	`

	err := r.DB.QueryRow(
		query,
		user.Email,
		user.Username,
		user.Password,
	).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

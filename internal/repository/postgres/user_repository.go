package postgres

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserRepository defines the methods for user-related database operations.
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// GetUsers fetches all users from the database.
func (r *UserRepository) GetUsers() ([]User, error) {
	rows, err := r.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetUsersByName(name string) ([]User, error) {
	rows, err := r.DB.Query("SELECT id, name, email FROM users WHERE name = $1", name)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetUserById(id int) (User, error) {
	row, err := r.DB.Query("SELECT id, name, email FROM users WHERE id = $1", id)
	if err != nil {
		return User{}, fmt.Errorf("failed to fetch user: %w", err)
	}
	defer row.Close()

	var user User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return User{}, fmt.Errorf("failed to scan user: %w", err)
		}
	}
	return user, nil

}

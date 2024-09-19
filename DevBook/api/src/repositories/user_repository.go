package repositories

import (
	"database/sql"
	"devbook_api/src/models"
	"fmt"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (repository UserRepository) FindAll(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)
	rows, err := repository.Db.Query("SELECT id, name, nick, email, created_at FROM user WHERE name LIKE ? OR nick LIKE ?", nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository UserRepository) FindById(id uint64) (models.User, error) {
	row, err := repository.Db.Query("SELECT id, name, nick, email, created_at FROM user WHERE id = ?", id)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository UserRepository) Create(user models.User) (uint64, error) {
	statement, err := repository.Db.Prepare("INSERT INTO user (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Printf("Error creating statement: %t", err)
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		fmt.Printf("Error executing statement: %t", err)
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Error reading result: %t", err)
		return 0, err
	}

	return uint64(lastId), nil
}

func (repository UserRepository) Update(userID uint64, user models.User) (uint64, error) {
	statement, err := repository.Db.Prepare("UPDATE user SET name = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, userID)
	if err != nil {
		return 0, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return uint64(affectedRows), nil
}

func (repository UserRepository) Delete(id uint64) (uint64, error) {
	statement, err := repository.Db.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(id)
	if err != nil {
		return 0, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return uint64(affectedRows), nil
}

func (repository UserRepository) FindByEmail(email string) (models.User, error) {
	row, err := repository.Db.Query("SELECT id, password, name FROM user WHERE email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Password, &user.Name); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository UserRepository) Follow(followerID, followedID uint64) error {
	statement, err := repository.Db.Prepare("INSERT IGNORE INTO follower (user_id, follower_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(followedID, followerID); err != nil {
		return err
	}

	return nil
}

func (repository UserRepository) Unfollow(followerID, followedID uint64) error {
	statement, err := repository.Db.Prepare("DELETE FROM follower WHERE user_id = ? AND follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(followedID, followerID); err != nil {
		return err
	}

	return nil
}

func (repository UserRepository) FindFollowers(userID uint64) ([]models.User, error) {
	rows, err := repository.Db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.created_at
		FROM user u
		INNER JOIN follower f ON u.id = f.follower_id
		WHERE f.user_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository UserRepository) FindFollowing(userID uint64) ([]models.User, error) {
	rows, err := repository.Db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.created_at
		FROM user u
		INNER JOIN follower f ON u.id = f.user_id
		WHERE f.follower_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository UserRepository) GetCurrentUserPassword(userID uint64) (string, error) {
	row, err := repository.Db.Query("SELECT password FROM user WHERE id = ?", userID)
	if err != nil {
		return "", err
	}
	defer row.Close()

	var password string
	if row.Next() {
		if err := row.Scan(&password); err != nil {
			return "", err
		}
	}

	return password, nil
}

func (repository UserRepository) UpdatePassword(userID uint64, password string) error {
	statement, err := repository.Db.Prepare("UPDATE user SET password = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, userID); err != nil {
		return err
	}

	return nil
}

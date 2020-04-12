package user_repo

import (
	"context"
	"database/sql"
	"github.com/rishi-suresh-keshav/user-service/models"
	"github.com/rishi-suresh-keshav/user-service/repository"
	"log"
)

func NewMysqlUserRepo(Conn *sql.DB) repository.UserRepo {
	return &mysqlUserRepo{Conn: Conn}
}

type mysqlUserRepo struct {
	Conn *sql.DB
}

func (repo *mysqlUserRepo) GetUserById(context context.Context, id int64) (*models.User, error) {
	query := "Select id, first_name, last_name, email From users Where id=?"

	rows, err := repo.Conn.QueryContext(context, query, id)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	println(rows)

	user := models.User{}
	for rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email)

		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (repo *mysqlUserRepo) CreateUser(context context.Context, user *models.User) (int64, error) {
	query := "Insert users SET first_name=?, last_name=?, email=?"

	stmt, err := repo.Conn.PrepareContext(context, query)
	if err != nil {
		panic(err)
		return -1, err
	}

	response, err := stmt.ExecContext(context, user.FirstName, user.LastName, user.Email)
	defer stmt.Close()
	if err != nil {
		return -1, err
	}

	return response.LastInsertId()
}

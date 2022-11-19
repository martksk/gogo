package repository

import (
	"database/sql"
	"fmt"
)

type userRepositoryDB struct {
	db *sql.DB
}

func NewUserRepositoryDB(db *sql.DB) userRepositoryDB {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) CreateUser(email string, password string, secret string) (*User, error){
	fmt.Printf("CreateUser")
	return nil, nil
}

func (r userRepositoryDB) CheckUser(emai string) (*User, error) {
	fmt.Println("CheckUser")
	return nil, nil
}
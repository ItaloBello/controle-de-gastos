package model

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

type User struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name" binding:"required"`
	Email     string    `db:"email" json:"email" binding:"required"`
	HashPass  string    `db:"hash_pass" json:"hashPass" binding:"required"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(u.HashPass)) //cria a senha criptografada em hash de md5
	//aqui usa o hash.sum(nil) para retornar o hash sem alterar nada e depois faz um encoding para passar de numerico para string
	u.HashPass = hex.EncodeToString(hash.Sum(nil))
}

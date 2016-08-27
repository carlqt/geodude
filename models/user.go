package models

import (
  "fmt"
  "golang.org/x/crypto/bcrypt"
  "encoding/json"
)

type User struct {
  ID int `json:"id"`
  Email string `json:"email"`
  Username string `json:"username"`
  Password string `json:"password"`
  PasswordConfirmation string `json:",omitempty"`
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  ContactNumber string `json:"contact"`
  Role string `json:"role"`
}

func (u *User) Validate() error {
  switch {
  case u.Password != u.PasswordConfirmation:
    return fmt.Errorf("Password does not match confirmation")
  }

  return nil
}

func (u *User) Create() error {
  hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
  stmnt, err := db.Prepare("INSERT INTO users(username, email, password, role) VALUES(?, ?, ?, ?)")

  if err == nil {
    result, _ := stmnt.Exec(u.Username, u.Email, hashedPassword, u.Role)
    id, _ := result.LastInsertId()
    u.ID = int(id)
    return nil
  } else {
    return err
  }
}

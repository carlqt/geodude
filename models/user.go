package models

import (
  "fmt"
  "golang.org/x/crypto/bcrypt"
)

type User struct {
  ID int
  Email string
  Username string
  Password string
  PasswordConfirmation string
  FirstName string
  LastName string
  ContactNumber string
  Role string
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

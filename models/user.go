package models

import (
  "fmt"
  "golang.org/x/crypto/bcrypt"
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
  err := db.QueryRow("INSERT INTO users(username, email, password, role) VALUES($1, $2, $3, $4) returning id", 
    u.Username, u.Email, hashedPassword, u.Role).Scan(&u.ID)


  if err == nil {
    return nil
  } else {
    return err
  }
}

// func (u *User) Authenticate() error{
//   var hashedPassword []byte

//   rows, err := db.Query("SELECT id, password FROM users WHERE username = $1", u.Username)
//   defer rows.Close()

//   if err != nil {
//     return err
//   }

//   for rows.Next() {
//     err := rows.Scan(&u.ID, &hashedPassword)

//     if err != nil {
//       revel.ERROR.Println(err)
//     }
//   }

//   err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(u.Password))

//   if err != nil {
//     return false
//   } else {
//     return true
//   }
// }

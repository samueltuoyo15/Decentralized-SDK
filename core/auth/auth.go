package auth

import (
  "errors"
  "golang.org/x/crypto/bcrypt"
)

var users = make(map[string]string)

func SignUp(email, password string) error {
  if _, exists := users[email]; exists{
    return errors.New("user already exist")
  }
  
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    return err
  }
  
  users[email] = string(hashedPassword)
  
  return nil
}

func Login(email, password string) error {
  
  hashedPassword, exists := users[email]
  if !exists{
    return errors.New("user does not exist")
  }
  
  return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
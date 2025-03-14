package auth 

import (
  "testing"
  )
  
func TestSignUp(t *testing.T){
  err := SignUp("user@example.com", "password123")
  
  if err != nil {
    t.Errorf("Sign up failed: %v", err)
  }
  
  err = SignUp("user@example.com", "password123")
  
  if err == nil {
    t.Error("Signnup should fail for duplicate users")
  }
}

func TestLogin(t *testing.T){
  err := SignUp("user@example.com", "password123")
  
  if err != nil {
    t.Fatalf("Sign up failed: %v", err)
  }
  
  err = Login("user@example.com", "password123")
  
  if err != nil {
    t.Errorf("Login up failed: %v", err)
  }
  
  err = Login("user@example.com", "testwrongpassword")
  
  if err == nil {
    t.Errorf("Login should fail for incorrect password")
  }
}
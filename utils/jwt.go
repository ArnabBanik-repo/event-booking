package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var sk string = os.Getenv("SECRET_KEY")

func GenerateToken(id string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
    "exp": time.Now().Add(24*time.Hour).Unix(),
	})

  return t.SignedString([]byte(sk))
}

func VerifyToken(t string) (string, error) {
  if t == ""{
    return "", errors.New("No cookie provided")
  } 

  pt, err := jwt.Parse(t, func (token *jwt.Token) (interface{}, error) {
    
    _, ok := token.Method.(*jwt.SigningMethodHMAC)
    if !ok {
      return nil, errors.New("Invalid signing method")
    }

    return []byte(sk), nil
  })

  if err != nil {
    return "", err
  }

  if !pt.Valid {
    return "", errors.New("Invalid jwt")
  }

  claims, ok := pt.Claims.(jwt.MapClaims)
  if !ok {
    return "", errors.New("Invalid jwt claims")
  }

  e := claims["exp"].(float64)
  if float64(time.Now().Unix()) > e {
    return "", errors.New("jwt expired. Login again")
  }

  return claims["id"].(string), nil
}

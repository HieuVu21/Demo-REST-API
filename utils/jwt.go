package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "2110"

func GenerateToken(email string, userID int64) (string, error) {
	fmt.Print(userID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}
func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil 
	})
	if err != nil {
		return 0, errors.New("cant parse token")
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId,nil
}

// package utils

// import (
// 	"errors"
// 	"time"
// 	"strconv"

// 	"github.com/golang-jwt/jwt/v5"
// )

// const secretKey = "2110"

// // GenerateToken creates a new JWT token with userID and email
// func GenerateToken(email string, userID int64) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"email":  email,
// 		"userId": userID,
// 		"exp":    time.Now().Add(time.Hour * 2).Unix(),
// 	})
// 	return token.SignedString([]byte(secretKey))
// }

// // VerifyToken parses and validates a JWT token, returning the userID if valid
// func VerifyToken(token string) (int64, error) {
// 	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 		_, ok := token.Method.(*jwt.SigningMethodHMAC)
// 		if !ok {
// 			return nil, errors.New("unexpected signing method")
// 		}
// 		return []byte(secretKey), nil
// 	})
// 	if err != nil {
// 		return 0, errors.New("can't parse token")
// 	}

// 	// Check if token is valid
// 	if !parsedToken.Valid {
// 		return 0, errors.New("invalid token")
// 	}

// 	// Extract claims
// 	claims, ok := parsedToken.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return 0, errors.New("invalid token claims")
// 	}

// 	// Safely extract userId (handle both float64 and string cases)
// 	userIdInterface := claims["userId"]
// 	switch v := userIdInterface.(type) {
// 	case float64:
// 		return int64(v), nil
// 	case string:
// 		userId, err := strconv.ParseInt(v, 10, 64)
// 		if err != nil {
// 			return 0, errors.New("invalid userId format")
// 		}
// 		return userId, nil
// 	default:
// 		return 0, errors.New("unexpected userId type")
// 	}
// }

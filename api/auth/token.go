package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//		CREATE TOKEN USERS
func CreateToken(user_id uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expayed sebelum 1 jam
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

//		EXTRACT TOKEN
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// //		CREATE TOKEN ADMINS
// func CreateTokenAdmins(admins_id uint32) (string, error) {
// 	claims := jwt.MapClaims{}
// 	claims["authorized"] = true
// 	claims["admins_id"] = admins_id
// 	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expayed sebelum 1 jam
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(os.Getenv("API_SECRET")))

// // }

// //		APABILA TOKEN VALID
// func TokenValid(r *http.Request) error {
// 	tokenString := ExtractToken(r)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("API_SECRET")), nil
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		Pretty(claims)
// 	}
// 	return nil
// }

//		extract token admin
func ExtractTokenIDAdmin(r *http.Request) (uint32, error) {

	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["admin_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
}

//		EXTRACT TOKEN ID
func ExtractTokenID(r *http.Request) (uint32, error) {

	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id, admins_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
}

// 	Cukup tampilkan klaim licely di terminal
func Pretty(data interface{}) {
	e, err := json.MarshalIndent(data, "", "")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(e))
}

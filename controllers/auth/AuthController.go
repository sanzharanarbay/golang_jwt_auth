package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"jwt_auth_golang/models"
	u "jwt_auth_golang/utils"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
)

type JwtToken struct {
	AccessToken string `json:"access-token"`
}

var jwt_secret = os.Getenv("jwt_secret")

func Login(w http.ResponseWriter, req *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	defer req.Body.Close()
	username := user.Username
	password := user.Password

	/* Another way to grab the form inputs from the request
	req.ParseForm()
	username := req.FormValue("Username")
	password := req.FormValue("Password")
	*/
	result := models.GetUsername(username)
	if result == nil {
		u.Respond(w, u.Message(false, "Your credentials do not match our records"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))

	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Your credentials do not match our records"))
		return
	}
	// access token ttl
	ttl := 2 * time.Minute
	accessTokenExpire := os.Getenv("access_token_expire")
	min, err := strconv.Atoi(accessTokenExpire)
	if err != nil {
		log.Println(err)
	}
	if accessTokenExpire != "" {
		ttl = time.Duration(min) * time.Minute
	}
	CreateToken(w, username, password, ttl)
}

func CreateToken(w http.ResponseWriter, username string, password string, ttl time.Duration) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      ttl,
	})

	tokenString, error := token.SignedString([]byte(jwt_secret))
	if error != nil {
		fmt.Println(error)
	}
	resp := u.Message(true, "success")
	resp["data"] = JwtToken{AccessToken: tokenString}
	u.Respond(w, resp)
	return
}


package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
	"jwt_auth_golang/handlers"
	"jwt_auth_golang/models"
	"jwt_auth_golang/modules/auth"
	"jwt_auth_golang/services"
	u "jwt_auth_golang/utils"
	"net/http"
	"os"
	"strconv"
)

var redisClient = services.InitRedis()
var rd = auth.NewAuth(redisClient)
var tk = auth.NewToken()
var service = handlers.NewProfile(rd, tk)


func Login(c *gin.Context) {
	user := &models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	username := user.Username
	password := user.Password

	/* Another way to grab the form inputs from the request
	req.ParseForm()
	username := req.FormValue("Username")
	password := req.FormValue("Password")
	*/
	result := models.GetUsername(username)
	if result == nil {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	userId := strconv.FormatUint(uint64(result.ID), 10)

	ts, err := service.Tk.CreateToken(userId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := service.Rd.CreateAuth(userId, ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	resp := u.Message(true, "success")
	resp["data"] = tokens
	u.Respond(c, http.StatusOK, resp)
	return

}

func Logout(c *gin.Context) {

	metadata, _ := service.Tk.ExtractTokenMetadata(c.Request)
	if metadata != nil {
		deleteErr := service.Rd.DeleteTokens(metadata)
		if deleteErr != nil {
			c.JSON(http.StatusBadRequest, deleteErr.Error())
			return
		}
	}
	c.JSON(http.StatusOK, "Successfully logged out")
	return
}

func Refresh(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}
	refreshToken := mapToken["refresh_token"]

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}
		userId, roleOk := claims["user_id"].(string)
		if  roleOk == false {
			c.JSON(http.StatusUnprocessableEntity, "unauthorized")
			return
		}
		//Delete the previous Refresh Token
		delErr := service.Rd.DeleteRefresh(refreshUuid)
		if delErr != nil { //if any goes wrong
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		//Create new pairs of refresh and access tokens
		ts, createErr := service.Tk.CreateToken(userId)
		if  createErr != nil {
			c.JSON(http.StatusForbidden, createErr.Error())
			return
		}
		//save the tokens metadata to redis
		saveErr := service.Rd.CreateAuth(userId, ts)
		if saveErr != nil {
			c.JSON(http.StatusForbidden, saveErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		resp := u.Message(true, "success")
		resp["data"] = tokens
		c.JSON(http.StatusCreated, resp)
	} else {
		c.JSON(http.StatusUnauthorized, "refresh expired")
	}
}





package jwt_auth

import (
	"fmt"
	"net/http"

	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	usercase UserUsecase
}

func (c AuthController) Login(ctx *gin.Context) {

	var loginReq LoginReq

	if err := ctx.BindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid input")
		return
	}
	username, password := c.usercase.FindByUsername(loginReq.Username)
	fmt.Println(username, password)

	// authenticate user (compare username dan password)
	if loginReq.Username == username && loginReq.Password == password {
		// generate JWT token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = loginReq.Username
		claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "gagal generate token"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unregistered user"})
	}
}

func NewUserController(u UserUsecase) *AuthController {
	controller := AuthController{
		usercase: u,
	}

	return &controller
}

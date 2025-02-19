package handlers

import (
	"fmt"
	"jwtgen/constants"
	"jwtgen/database"
	"jwtgen/initializers"
	"jwtgen/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "can't convert to hash",
		})
		return
	}

	user := models.User{
		Email:    body.Email,
		Password: string(hash),
	}
	result := database.CreateUser(user)
	if result != nil {
		fmt.Println("errr :: ", result)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
	})
}

func LogIn(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}
	result, err := database.AuthorizeUser(body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password not matched",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user": body.Email,
		"exp":  time.Now().Add(time.Hour).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	// privateKey, _ := initializers.LoadKeys()
	tokenString, err := token.SignedString(initializers.PrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create the token",
		})
	}
	//set cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, time.Now().Hour(), "/", "", false, false)
	c.Redirect(http.StatusMovedPermanently, constants.ModelRedirect)

	c.JSON(http.StatusOK, gin.H{
		"message": "log in successful",
	})
}

func GetUserData(c *gin.Context) {
	var body struct {
		Email string `json:"email"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
	}
	result, err := database.GetUser(body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Details": result,
	})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in",
	})
}

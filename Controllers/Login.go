package Controllers

import (
	"userauth_with_crud/Models"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

//GetUserByUsername ... Get the user by username
func LoginUserByUsername(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)

	var user_found Models.User
	err := Models.GetUserByUsername(&user_found, user.Username)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if comparePasswords(user_found.Password, []byte(user.Password)) {
		atClaims := jwt.MapClaims{}
		atClaims["authorized"] = true
		atClaims["username"] = user.Username
		atClaims["level"] = user.Level
		atClaims["exp"] = time.Now().Add(time.Hour * 24 * 2).Unix()
		at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
		token, err := at.SignedString([]byte("secret"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
		}
		c.JSON(http.StatusOK, gin.H{
			"username": user.Username,
			"level":    user.Level,
			"token":    token,
		})
	} else {
		c.JSON(http.StatusForbidden, "Wrong Password")
	}

}

package Routes

import (
	"userauth_with_crud/Controllers"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", auth, Controllers.GetUsers)
		grp1.POST("user", Controllers.CreateUser)
		grp1.POST("userlogin", Controllers.LoginUserByUsername)
		grp1.GET("user/:id", auth, Controllers.GetUserByID)
		grp1.PUT("user/:id", auth, Controllers.UpdateUser)
		grp1.DELETE("user/:id", auth, Controllers.DeleteUser)
	}
	return r
}

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		// fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}

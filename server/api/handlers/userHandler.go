package handlers

import (
	"encoding/json"
	models "go-gin-boilerplate/server/api/models"
	services "go-gin-boilerplate/server/api/services"
	"go-gin-boilerplate/server/utils/confighelper"
	jwthelper "go-gin-boilerplate/server/utils/jwthelper"
	"go-gin-boilerplate/server/utils/logginghelper"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Save user to DB
func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{}
		byteArray, err := bindData(c)
		if err != nil {
			logginghelper.LogError("ERROR : Error occurred while binding token  ---> ", err)
			c.JSON(http.StatusInternalServerError, "")
		}
		err = json.Unmarshal(byteArray, &user)
		if err != nil {
			logginghelper.LogError("ERROR : Error occurred while unmarshalling  ---> ", err)
			c.JSON(http.StatusInternalServerError, "")
		}
		ctx := c.Request.Context()
		err = services.AddUser(ctx, user)
		if err != nil {
			logginghelper.LogError("ERROR : Error occurred in service  ---> ", err)
			c.JSON(http.StatusInternalServerError, "")
		}
		c.JSON(http.StatusOK, "User added successfully!!!")
	}
}

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userLogin := models.Login{}
		byteArray, err := bindData(c)
		if err != nil {
			logginghelper.LogError("ERROR : Error occurred while binding token  ---> ", err)
			c.JSON(http.StatusInternalServerError, "")
		}
		err = json.Unmarshal(byteArray, &userLogin)
		if err != nil {
			logginghelper.LogError("ERROR : Error occurred while unmarshalling  ---> ", err)
			c.JSON(http.StatusInternalServerError, "")
		}
		userDetails, isValidUser, err := services.ValidateUser(userLogin)
		if err != nil {
			logginghelper.LogError("ERROR : Invalid user!  ---> ", err)
			c.JSON(http.StatusInternalServerError, "")
		}

		if isValidUser {
			token, err := jwthelper.GenerateToken(userDetails.UserName, userDetails.Password, 24*time.Hour)
			if err != nil {
				logginghelper.LogError("ERROR : Error while generating token:  ---> ", err)
				c.JSON(http.StatusInternalServerError, "")
			} else {
				c.Header("Authorization", token)
				c.JSON(http.StatusOK, userDetails)
				return
			}

		} else {
			c.String(203, "Enter correct credentials!!!")
		}

	}
}

// convert the encoded payload to json
func bindData(c *gin.Context) ([]byte, error) {
	var byteArray []byte
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logginghelper.LogError("ERROR : Error occurred while reading request  ---> ", err)
		return byteArray, err
	}
	tokenString := string(body)
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(confighelper.GetConfig("PrivateKey")), nil
	})
	if err != nil {
		logginghelper.LogError("ERROR : Error occurred while decoding token  ---> ", err)
		return byteArray, err
	}

	// Marshal the map into a JSON string.
	interfaceData, err := json.Marshal(claims)
	if err != nil {
		logginghelper.LogError("ERROR : Error occurred while unmarshalling claims  ---> ", err)
		return byteArray, err
	}

	return interfaceData, nil
}

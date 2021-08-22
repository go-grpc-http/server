package api

import (
	"context"
	dao "go-gin-boilerplate/server/api/dao"
	models "go-gin-boilerplate/server/api/models"
	"go-gin-boilerplate/server/utils/logginghelper"
	"strconv"
	"strings"
)

// AddUser service will send the user to DAO
func AddUser(ctx context.Context, user models.User) error {
	userDetails := GenerateCredentials(user)
	// Mongo DB
	err := dao.SaveUserToDB(ctx, userDetails)
	if err != nil {
		logginghelper.LogError("ERROR : Error occurred while AddUser  ---> ", err)
		return err
	}

	//MYSQL
	// SaveUserToMYSQLDB(ctx, user)
	return nil
}

func GenerateCredentials(user models.User) models.User {
	// strconv.FormatInt(int64(i2), 10)
	user.UserName = strings.ToLower(user.FirstName) + strings.ToLower(user.LastName[:1]) + digits(user.Mobile)
	user.Password = strings.ToLower(user.LastName) + strings.ToLower(user.FirstName[:1]) + digits(user.Mobile)
	return user
}

func GetLoginCredentials(user models.User) models.Login {
	return models.Login{
		UserName: user.UserName,
		Password: user.Password,
	}
}

func ValidateUser(userRequest models.Login) (models.User, bool, error) {
	userLogin, userFound, err := dao.GetUserFromDB(userRequest)
	if userFound {
		return userLogin, true, nil
	}
	return userLogin, false, err
}

func digits(num int) (firstThreeDigits string) {
	firstThree := num / 10000000
	firstThreeDigits = strconv.Itoa(firstThree)
	return
}

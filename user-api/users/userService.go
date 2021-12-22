package users

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FelixAnna/web-service-dlw/user-api/users/entity"
	"github.com/FelixAnna/web-service-dlw/user-api/users/repository"
	"github.com/gin-gonic/gin"
)

var repo repository.UserRepo

func init() {
	repo = &repository.UserRepoDynamoDB{}
}

func GetAllUsers(c *gin.Context) {
	users, err := repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := repo.GetByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserById(c *gin.Context) {
	strId := c.Param("userId")
	user, err := repo.GetById(strId)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserBirthdayById(c *gin.Context) {
	userId := c.Param("userId")
	birthday := c.Query("birthday")
	err := repo.UpdateBirthday(userId, birthday)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("User birthday updated, userId: %v.", userId))
}

func UpdateUserAddressById(c *gin.Context) {
	userId := c.Param("userId")
	var addresses []entity.Address
	if err := c.BindJSON(&addresses); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err := repo.UpdateAddress(userId, addresses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("User address updated, userId: %v.", userId))
}

func AddUser(c *gin.Context) {
	var new_user entity.User
	if err := c.BindJSON(&new_user); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	id, err := repo.Add(&new_user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("User %v created!", *id))
}

func RemoveUser(c *gin.Context) {
	userId := c.Param("userId")
	err := repo.Delete(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("User %v deleted!", userId))
}

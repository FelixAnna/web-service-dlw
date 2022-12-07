package users

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FelixAnna/web-service-dlw/user-api/users/entity"
	"github.com/FelixAnna/web-service-dlw/user-api/users/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserApi), "*"))

type UserApi struct {
	Repo repository.UserRepo
}

func (api *UserApi) GetAllUsers(c *gin.Context) {
	users, err := api.Repo.GetAll()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (api *UserApi) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := api.Repo.GetByEmail(email)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (api *UserApi) GetUserById(c *gin.Context) {
	strId := c.Param("userId")
	user, err := api.Repo.GetById(strId)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (api *UserApi) UpdateUserBirthdayById(c *gin.Context) {
	userId := c.Param("userId")
	birthday := c.Query("birthday")
	err := api.Repo.UpdateBirthday(userId, birthday)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("User birthday updated, userId: %v.", userId))
}

func (api *UserApi) UpdateUserAddressById(c *gin.Context) {
	userId := c.Param("userId")
	var addresses []entity.Address
	if err := c.BindJSON(&addresses); err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := api.Repo.UpdateAddress(userId, addresses)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("User address updated, userId: %v.", userId))
}

func (api *UserApi) AddUser(c *gin.Context) {
	var new_user entity.User
	if err := c.BindJSON(&new_user); err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	id, err := api.Repo.Add(&new_user)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("User %v created!", *id))
}

func (api *UserApi) RemoveUser(c *gin.Context) {
	userId := c.Param("userId")
	err := api.Repo.Delete(userId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("User %v deleted!", userId))
}

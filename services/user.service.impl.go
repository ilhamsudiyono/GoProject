package services

import (
	"ProjectFirst/config"
	"ProjectFirst/dao"
	"ProjectFirst/models"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var userDao dao.UserDao = dao.UserDaoImpl{}

type UserServiceImp struct{}

func (UserServiceImp) CreateUser(user *models.Users) (u *models.Users, e error) {
	defer config.CatchError(&e)

	result, err := bcrypt.GenerateFromPassword([]byte(user.Pwd), 4)

	if err == nil {
		user.Pwd = string(result)
		var createDt = models.Timestamp(time.Now())
		user.CreatedDate = &createDt
		return userDao.CreateUser(user)
	}
	return nil, err
}

func (UserServiceImp) GetUserById(id string) (u models.Users, e error) {
	defer config.CatchError(&e)
	u, err := userDao.GetUserById(id)

	u.Pwd = ""
	return u, err
}

func (UserServiceImp) Login(username string, pwd string) (u models.Users, e error) {
	defer config.CatchError(&e)

	result, err := userDao.GetUserByUsername(username)

	if err == nil {
		var err = bcrypt.CompareHashAndPassword([]byte(result.Pwd), []byte(pwd))
		if err == nil {
			result.Pwd = ""
			v, _ := config.GenerateToken(username)
			result.Token = &v
			return result, nil
		}

	}
	return models.Users{}, errors.New("Username/Password incorrect")
}

package services

import (
	"errors"
	"strings"

	"goscrum/goscrum/server/db"
	"goscrum/goscrum/server/graph/model"
	"goscrum/goscrum/server/util"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	dbClient *gorm.DB
}

func NewUserService(dbClient *gorm.DB) UserService {
	return UserService{dbClient: dbClient}
}

var ErrRoleNotValid = errors.New("role is invalid")

func (u UserService) GetUsers(limit, offset int, search string) ([]*model.User, error) {
	var dbUsers []db.User
	query := u.dbClient.
		Limit(limit).
		Offset(offset)

	if strings.TrimSpace(search) != "" {
		searchQuery := "%" + strings.TrimSpace(search) + "%"
		query = query.Where("first_name LIKE ? OR last_name LIKE ? OR email LIKE ?",
			searchQuery, searchQuery, searchQuery)
	}

	err := query.
		Find(&dbUsers).
		Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}

	var users []*model.User
	err = util.Copy(&users, &dbUsers)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u UserService) GetUsersCount(search string) (int, error) {
	var count int
	query := u.dbClient.Model(&db.User{})
	if strings.TrimSpace(search) != "" {
		searchQuery := "%" + strings.TrimSpace(search) + "%"
		query = query.Where("first_name LIKE ? OR last_name LIKE ? OR email LIKE ?",
			searchQuery, searchQuery, searchQuery)
	}
	err := query.Count(&count).Error
	return count, err
}

func (u UserService) GetUser(id string) (*model.User, error) {
	var dbUser db.User
	err := u.dbClient.Where("id = ?", id).First(&dbUser).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	var user model.User
	err = util.Copy(&user, &dbUser)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

package usecase

import "module31/internal/entity"

type (
	Usecase interface {
		CreateUser(*entity.User) (string, error)
		DeleteUser(string) (string, error)
		UpdateUser(string, int) error
		GetFriends(string) ([]string, error)
		MakeFriends(string, string) (string, string, error)
		GetUsers() []*entity.User
	}

	Repository interface {
		CreateUser(*entity.User) (string, error)
		DeleteUser(string) (string, error)
		UpdateAge(string, int) error
		GetFriends(string) ([]string, error)
		MakeFriends(string, string) (string, string, error)
		GetUsers() []*entity.User
	}
)

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) *usecase {
	return &usecase{
		repository: repository,
	}
}

//CreateUser accepts user struck, sends to the repository and return user id
func (u *usecase) CreateUser(user *entity.User) (string, error) {
	uid, err := u.repository.CreateUser(user)
	return uid, err
}

//DeleteUser accepts user id, sends to the repository and return username
func (u *usecase) DeleteUser(id string) (string, error) {
	name, err := u.repository.DeleteUser(id)
	return name, err
}

//GetUsers sends to the repository and return slice of all users
func (u *usecase) GetUsers() []*entity.User {
	allUsers := u.repository.GetUsers()
	return allUsers
}

//UpdateUser accepts user id and new age,sends to the repository
func (u *usecase) UpdateUser(id string, newAge int) error {
	err := u.repository.UpdateAge(id, newAge)
	return err
}

//MakeFriends accepts target and source id,sends to the repository and returns users names
func (u *usecase) MakeFriends(target string, source string) (string, string, error) {
	name1, name2, err := u.repository.MakeFriends(target, source)
	return name1, name2, err
}

//GetFriends accepts user id,sends to the repository and return slice of friends names
func (u *usecase) GetFriends(userId string) ([]string, error) {
	allUsers, err := u.repository.GetFriends(userId)
	return allUsers, err
}

package usecase

import "module31/internal/entity"

type (
	Usecase interface {
		CreateUser(*entity.User) (int, error)
		DeleteUser(int) (string, error)
		UpdateUser(int, int) error
		//GetFriends(int) ([]int, error)
		MakeFriends(int, int) (string, string, error)
		GetUsers() map[int]*entity.User
	}

	Repository interface {
		CreateUser(*entity.User) (int, error)
		DeleteUser(int) (string, error)
		UpdateAge(int, int) error
		//GetFriends(int) ([]string, error)
		MakeFriends(int, int) (string, string, error)
		GetUsers() map[int]*entity.User
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

func (u *usecase) CreateUser(user *entity.User) (int, error) {
	uid, error := u.repository.CreateUser(user)
	return uid, error
}
func (u *usecase) DeleteUser(id int) (string, error) {
	name, error := u.repository.DeleteUser(id)
	return name, error
}
func (u *usecase) GetUsers() map[int]*entity.User {
	allUsers := u.repository.GetUsers()
	return allUsers
}
func (u *usecase) UpdateUser(id int, newAge int) error {
	err := u.repository.UpdateAge(id, newAge)
	return err
}
func (u *usecase) MakeFriends(target int, source int) (string, string, error) {
	name1, name2, err := u.repository.MakeFriends(target, source)
	return name1, name2, err
}

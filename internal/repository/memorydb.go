package repository

import (
	"module31/internal/entity"
	"sync"
)

type memorydb struct {
	sync.Mutex
	index     int
	usersById map[int]*entity.User
}

func NewMemorydb() (*memorydb, error) {
	return &memorydb{
		usersById: make(map[int]*entity.User),
	}, nil
}

func (r *memorydb) CreateUser(user *entity.User) (int, error) {
	r.Lock()
	defer r.Unlock()
	r.index++
	user.Id = r.index
	r.usersById[user.Id] = user
	return user.Id, nil
}
func (r *memorydb) DeleteUser(id int) (string, error) {
	r.Lock()
	defer r.Unlock()
	for i, u := range r.usersById {
		for j, f := range u.Friends {
			if f == r.usersById[id].Name {
				r.usersById[i].Friends = append(u.Friends[:j], u.Friends[j+1:]...)
			}
		}
	}
	deletedName := r.usersById[id].Name
	delete(r.usersById, id)
	return deletedName, nil
}
func (r *memorydb) GetUsers() map[int]*entity.User {
	r.Lock()
	defer r.Unlock()

	return r.usersById
}
func (r *memorydb) UpdateAge(id int, newAge int) error {
	r.Lock()
	defer r.Unlock()
	r.usersById[id].Age = newAge
	return nil
}
func (r *memorydb) MakeFriends(target int, source int) (string, string, error) {
	r.Lock()
	defer r.Unlock()
	r.usersById[target].Friends = append(r.usersById[target].Friends, r.usersById[source].Name)
	r.usersById[source].Friends = append(r.usersById[source].Friends, r.usersById[target].Name)
	return r.usersById[target].Name, r.usersById[source].Name, nil
}

package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
func readFile(m *memorydb) *memorydb {
	rawDataIn, _ := ioutil.ReadFile("userStorage.json")
	_ = json.Unmarshal(rawDataIn, &m.usersById)
	for i := range m.usersById {
		for j := range m.usersById {
			if m.usersById[i].Id < m.usersById[j].Id {
				m.index = m.usersById[j].Id
			}
		}
	}
	return m
}
func writeFile(m *memorydb) {
	rawDataOut, err := json.MarshalIndent(m.usersById, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}
	err = ioutil.WriteFile("userStorage.json", rawDataOut, 0644)
	if err != nil {
		log.Fatal("Cannot write updated file:", err)
	}
}

func (r *memorydb) CreateUser(user *entity.User) (int, error) {
	r.Lock()
	defer r.Unlock()
	readFile(r)

	r.index++
	user.Id = r.index
	r.usersById[user.Id] = user
	writeFile(r)
	return user.Id, nil
}
func (r *memorydb) DeleteUser(id int) (string, error) {
	r.Lock()
	defer r.Unlock()
	readFile(r)
	for i, u := range r.usersById {
		for j, f := range u.Friends {
			if f == r.usersById[id].Name {
				r.usersById[i].Friends = append(u.Friends[:j], u.Friends[j+1:]...)
			}
		}
	}
	deletedName := r.usersById[id].Name
	delete(r.usersById, id)
	writeFile(r)
	return deletedName, nil
}
func (r *memorydb) GetUsers() map[int]*entity.User {
	r.Lock()
	defer r.Unlock()
	readFile(r)
	return r.usersById
}
func (r *memorydb) UpdateAge(id int, newAge int) error {
	r.Lock()
	defer r.Unlock()
	readFile(r)
	r.usersById[id].Age = newAge
	writeFile(r)
	return nil
}
func (r *memorydb) MakeFriends(target int, source int) (string, string, error) {
	r.Lock()
	defer r.Unlock()
	readFile(r)
	r.usersById[target].Friends = append(r.usersById[target].Friends, r.usersById[source].Name)
	r.usersById[source].Friends = append(r.usersById[source].Friends, r.usersById[target].Name)
	writeFile(r)
	return r.usersById[target].Name, r.usersById[source].Name, nil
}
func (r *memorydb) GetFriends(userId int) ([]string, error) {
	r.Lock()
	defer r.Unlock()
	readFile(r)
	friends := r.usersById[userId].Friends
	return friends, nil
}

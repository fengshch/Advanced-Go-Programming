package _1_before

import (
	"database/sql"
	"net/http"
	"strconv"
)

type CreateUserEndpoint struct {
	db *sql.DB
}

func (s *CreateUserEndpoint) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	user, err := s.extractUser(req)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.validate(user)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := s.saveToDB(user)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	resp.WriteHeader(http.StatusCreated)
	resp.Header().Set("location", "/user/"+strconv.Itoa(id))
}

// extract user from HTTP request
func (s *CreateUserEndpoint) extractUser(request *http.Request) (*User, error) {
	// implementation removed
	return &User{}, nil
}

// validate the supplied user is complete and correct
func (s *CreateUserEndpoint) validate(user *User) error {
	// implementation removed
	return nil
}

// save the supplied user to the database
func (s *CreateUserEndpoint) saveToDB(user *User) (int, error) {
	// implementation removed
	return 0, nil
}

type User struct {
	Name string
}

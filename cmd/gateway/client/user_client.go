package client

import (
	"example.com/models"
	"fmt"
	"net/http"
)

type UserClient struct {
	Url string
}

func GetUserClient() *UserClient {
	return &UserClient{
		Url: "localhost:8083",
	}
}

func (u *UserClient) GetUsers() ([]models.User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users", u.Url), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(req)
	return nil, nil
}

package client

import (
	"encoding/json"
	"example.com/models"
	"fmt"
	"net/http"
)

type UserClient struct {
	Url string
}

func GetUserClient() *UserClient {
	return &UserClient{
		Url: "http://localhost:8083",
	}
}

func (u *UserClient) GetUsers() ([]models.User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users", u.Url), nil)
	var users []models.User
	resp, err := http.DefaultClient.Do(req)
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

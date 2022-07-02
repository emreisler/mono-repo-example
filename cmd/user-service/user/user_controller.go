package user

import (
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := getUsers()
	if err != nil {
		return err
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		return err
	}
	return nil
}

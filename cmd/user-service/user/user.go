package user

import (
	"example.com/models"
)

//getUsers Getting users from DB - mock data
func getUsers() ([]models.User, error) {
	return []models.User{
		models.User{1, "User1", "Emre", "ISLER", 25.6},
		models.User{2, "User2", "Michael", "JORDAN", 27.3},
	}, nil
}

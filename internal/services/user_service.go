package services

import "my_project_go/internal/models"

func GetAllUsers() []models.User {
	return []models.User{
		{ID: 1, Name: "Alice", Age: 25},
		{ID: 2, Name: "Bob", Age: 30},
	}
}

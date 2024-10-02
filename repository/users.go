package repository

import (
	"go-auth/models"
)

var users = map[string]models.User{
	"admin":     {Username: "admin", Password: "admin123", Role: "ADM"},
	"professor": {Username: "professor", Password: "professor123", Role: "Professor"},
	"candidato": {Username: "candidato", Password: "candidato123", Role: "Candidato"},
	"empresa":   {Username: "empresa", Password: "empresa123", Role: "Empresa"},
}

func GetUsers() map[string]models.User {
	return users
}

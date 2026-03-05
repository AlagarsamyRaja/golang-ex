package service

import (
	"fmt"
	"sqlConnection/models"
)

func Insert() string {
	var user models.UserData
	var name, role string

	fmt.Println("Enter name:")
	fmt.Scanf("%s\n", &name)
	user.Name = name

	fmt.Println("Enter role:")
	fmt.Scanf("%s\n", &role)
	user.Role = role

	insertRow := fmt.Sprintf("INSERT INTO employee(name,role) VALUES ('%s','%s')", user.Name, user.Role)

	return insertRow
}

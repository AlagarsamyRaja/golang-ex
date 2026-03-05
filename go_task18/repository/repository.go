package repository

import (
	"fmt"
	"sqlConnection/pkg"
	"sqlConnection/service"
)

func DisplayRow() ([]string, error) {

	var tableRows []string

	rows, err := pkg.DB.Query("SELECT * FROM employee")
	if err != nil {
		return nil, fmt.Errorf("Check the query")
	}

	var id int
	var name, role string

	fmt.Print("Id\tName\tRole\n")
	for rows.Next() {
		rows.Scan(&id, &name, &role)
		tableRows = append(tableRows, fmt.Sprintf("%d\t%s\t%s", id, name, role))
	}

	return tableRows, nil
}

func InsertRow() error {

	insert := service.Insert()
	_, err := pkg.DB.Query(insert)
	if err != nil {
		return fmt.Errorf("Rows not inserted %s", err)
	}

	fmt.Println("1 row inserted")
	return nil

}

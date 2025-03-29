package repository

import (
	"context"
	"log"

	"github.com/usernamenenad/ftn-acr/models/database"
	models "github.com/usernamenenad/ftn-acr/models/pg"
)

func GetAllClassrooms(ctx context.Context) []models.Classroom {
	conn, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var classrooms []models.Classroom
	query := "SELECT * FROM classrooms;"
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var classroom models.Classroom
		err = rows.Scan(&classroom.Id, &classroom.ClassroomName, &classroom.ClassroomBuilding)
		if err != nil {
			log.Fatal(err)
		}
		classrooms = append(classrooms, classroom)
	}

	return classrooms
}

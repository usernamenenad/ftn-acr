package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/usernamenenad/ftn-acr/models/database"
	models "github.com/usernamenenad/ftn-acr/models/pg"
)

func GetAllPrograms(ctx context.Context) []models.Program {
	conn, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var programs []models.Program
	query := "SELECT * FROM programs;"
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var program models.Program
		err = rows.Scan(&program.Id, &program.ProgramName, &program.ProgramShortName)
		if err != nil {
			log.Fatal(err)
		}
		programs = append(programs, program)
	}

	return programs
}

func GetProgramByName(ctx context.Context, name string) models.Program {
	conn, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var program models.Program
	query := fmt.Sprintf("SELECT * FROM programs WHERE program_name = '%s';", name)
	if err = conn.QueryRowContext(ctx, query).Scan(&program.Id, &program.ProgramName, &program.ProgramShortName); err != nil {
		log.Fatal(err)
	}

	return program
}

func GetProgramByShortName(ctx context.Context, shortName string) models.Program {
	conn, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var program models.Program
	query := fmt.Sprintf("SELECT * FROM programs WHERE program_short_name = '%s';", shortName)
	if err = conn.QueryRowContext(ctx, query).Scan(&program.Id, &program.ProgramName, &program.ProgramShortName); err != nil {
		log.Fatal(err)
	}

	return program
}

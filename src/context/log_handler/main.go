package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ddd/src/context/log_handler/app/command"
	commandM "github.com/ddd/src/context/match_reporting/app/command"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"

	"github.com/ddd/src/context/log_handler/infra/service"
	serviceM "github.com/ddd/src/context/match_reporting/infra/service"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	selectLogFileGr, ctx := errgroup.WithContext(context.Background())

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	app := service.NewApplication(ctx, selectLogFileGr, db)
	appM := serviceM.NewApplication(ctx)

	defer db.Close()

	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	pathFile := wd + "/crosscutting/building_blocks/infra/tmp/qgames.log"
	selectLogFileCommand := command.SelectLogFileCommand{Path: pathFile}
	resultLogFile, err := app.Commands.SelectLogFile.Handle(ctx, selectLogFileCommand)
	if err != nil {
		panic(err.Error())
	}

	createHumanLogFileCommand := command.CreateHumanLogFileCommand{Content: resultLogFile}
	resultHumanLogFile, err := app.Commands.CreateHumanLogFile.Handle(ctx, createHumanLogFileCommand)
	if err != nil {
		panic(err.Error())
	}

	rawData := func() [][]string {
		var rawData [][]string
		for _, row := range resultHumanLogFile {
			rawData = append(rawData, []string{
				row.GetPlayerWhoKilled(),
				row.GetPlayerWhoDied(),
				row.GetMeanOfDeath(),
			})
		}
		return rawData
	}

	findPlayersKilledCommand := commandM.FindPlayersKilledCommand{Data: rawData()}
	resultPlayersKilled, err := appM.Commands.FindPlayersKilled.Handle(ctx, findPlayersKilledCommand)
	fmt.Print(resultPlayersKilled)

	if err := selectLogFileG.Wait(); err != nil {
		fmt.Println("Received error:", err)
	}

}

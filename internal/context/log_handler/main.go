package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ddd/internal/context/log_handler/app/command"
	commandM "github.com/ddd/internal/context/match_reporting/app/command"
	"golang.org/x/sync/errgroup"

	"github.com/ddd/internal/context/log_handler/infra/service"
	serviceM "github.com/ddd/internal/context/match_reporting/infra/service"
)

func main() {

	selectLogFileGr, ctx := errgroup.WithContext(context.Background())

	app := service.NewApplication(ctx, selectLogFileGr)
	appM := serviceM.NewApplication(ctx)

	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	pathFile := wd + "/pkg/building_blocks/infra/tmp/qgames.log"
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
	if err == nil {
		fmt.Println(resultPlayersKilled)
	}

	if err := selectLogFileGr.Wait(); err != nil {
		fmt.Println("Received error:", err)
	}

}

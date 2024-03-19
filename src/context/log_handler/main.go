package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/ddd/src/context/log_handler/app/command"
	commandM "github.com/ddd/src/context/match_reporting/app/command"

	"github.com/ddd/src/context/log_handler/infra/service"
	serviceM "github.com/ddd/src/context/match_reporting/infra/service"
)

func main() {

	ctx := context.Background()
	selectLogFileWg := sync.WaitGroup{}
	app := service.NewApplication(ctx, &selectLogFileWg)
	appM := serviceM.NewApplication(ctx)
	defer selectLogFileWg.Wait()

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

}

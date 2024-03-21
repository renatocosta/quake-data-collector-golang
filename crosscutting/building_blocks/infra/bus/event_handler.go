package bus

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ddd/crosscutting/building_blocks/domain"
	"github.com/ddd/src/context/log_handler/infra/ports/persistence"
	"golang.org/x/sync/errgroup"
)

type AdditionalDependencies struct {
	LogFileRepo persistence.LogFileRepository
}

func NewAdditionalDependencies(db *sql.DB) AdditionalDependencies {
	return AdditionalDependencies{LogFileRepo: persistence.NewLogFileRepository(db)}
}

type EventHandlerFunc func(context.Context, domain.Event, AdditionalDependencies) error

// UserRegisteredHandler handles the user registered event
func HandleEvent(ctx context.Context, g *errgroup.Group, eventChan <-chan domain.Event, additionalDependencies AdditionalDependencies, handlers map[string][]EventHandlerFunc) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Event Handler channel cancelled")
			return
		case event := <-eventChan:

			if handlersFunc, ok := handlers[event.Type]; ok {

				for _, handlerFunc := range handlersFunc {
					g.Go(func() error {
						return handlerFunc(ctx, event, additionalDependencies)
					})
				}

				continue
			}

			panic(fmt.Sprintf("Unknown event type: %s", event.Type))

		}

	}

}

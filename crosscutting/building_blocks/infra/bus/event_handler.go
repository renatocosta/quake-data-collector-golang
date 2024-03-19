package bus

import (
	"context"
	"fmt"
	"sync"

	"github.com/ddd/crosscutting/building_blocks/domain"
)

// UserRegisteredHandler handles the user registered event
func HandleEvent(ctx context.Context, wg *sync.WaitGroup, eventChan <-chan domain.Event, handlers map[string][]func(context.Context, domain.Event)) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Event Handler channel cancelled")
			return
		case event := <-eventChan:

			if handlersFunc, ok := handlers[event.Type]; ok {

				for _, handlerFunc := range handlersFunc {
					wg.Done()
					handlerFunc(ctx, event)
				}

				continue
			}

			panic(fmt.Sprintf("Unknown event type: %s", event.Type))

		}

	}
}

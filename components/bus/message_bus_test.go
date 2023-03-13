package bus_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/bus"
	"github.com/ThreeDotsLabs/watermill/message"
)

func TestNewEventProcessor(t *testing.T) {
	generateTopic := func(messageType string) string {
		return messageType
	}
	subscriberConstructor := func(handlerName string) (subscriber message.Subscriber, e error) {
		return nil, nil
	}

	router, err := message.NewRouter(message.RouterConfig{}, watermill.NopLogger{})
	require.NoError(t, err)

	messageBus, err := bus.NewMessageBus(router, bus.JSONMarshaler{}, generateTopic, subscriberConstructor, watermill.NopLogger{})
	require.NoError(t, err)

	err = bus.AddHandler(messageBus,
		"testHandler",
		func(ctx context.Context, event ExampleEvent) error {
			fmt.Println("Handling", event.ID)
			return nil
		},
	)
	require.NoError(t, err)
}

type ExampleEvent struct {
	ID string
}
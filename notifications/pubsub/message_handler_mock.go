// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package pubsub

import (
	"sync"
)

var (
	lockMessageHandlerMockHandleMessage sync.RWMutex
)

// MessageHandlerMock is a mock implementation of MessageHandler.
//
//     func TestSomethingThatUsesMessageHandler(t *testing.T) {
//
//         // make and configure a mocked MessageHandler
//         mockedMessageHandler := &MessageHandlerMock{
//             HandleMessageFunc: func(topic string, data []byte) error {
// 	               panic("TODO: mock out the HandleMessage method")
//             },
//         }
//
//         // TODO: use mockedMessageHandler in code that requires MessageHandler
//         //       and then make assertions.
//
//     }
type MessageHandlerMock struct {
	// HandleMessageFunc mocks the HandleMessage method.
	HandleMessageFunc func(topic string, data []byte) error

	// calls tracks calls to the methods.
	calls struct {
		// HandleMessage holds details about calls to the HandleMessage method.
		HandleMessage []struct {
			// Topic is the topic argument value.
			Topic string
			// Data is the data argument value.
			Data []byte
		}
	}
}

// HandleMessage calls HandleMessageFunc.
func (mock *MessageHandlerMock) HandleMessage(topic string, data []byte) error {
	if mock.HandleMessageFunc == nil {
		panic("moq: MessageHandlerMock.HandleMessageFunc is nil but MessageHandler.HandleMessage was just called")
	}
	callInfo := struct {
		Topic string
		Data  []byte
	}{
		Topic: topic,
		Data:  data,
	}
	lockMessageHandlerMockHandleMessage.Lock()
	mock.calls.HandleMessage = append(mock.calls.HandleMessage, callInfo)
	lockMessageHandlerMockHandleMessage.Unlock()
	return mock.HandleMessageFunc(topic, data)
}

// HandleMessageCalls gets all the calls that were made to HandleMessage.
// Check the length with:
//     len(mockedMessageHandler.HandleMessageCalls())
func (mock *MessageHandlerMock) HandleMessageCalls() []struct {
	Topic string
	Data  []byte
} {
	var calls []struct {
		Topic string
		Data  []byte
	}
	lockMessageHandlerMockHandleMessage.RLock()
	calls = mock.calls.HandleMessage
	lockMessageHandlerMockHandleMessage.RUnlock()
	return calls
}
package services

import (
	"testing"
	mocks "testing-demo/mocks/services"
)

func Test_MessageProcessor_Sends_Message(t *testing.T) {
	//arrange
	messageService := &mocks.MessageService{}
	messageService.On("Send", "Hello World").Return(true)

	messageProcessor := NewMessageProcessor(messageService)

	//act
	result := messageProcessor.Process("Hello World")

	//assert
	messageService.AssertExpectations(t)
	if !result {
		t.Errorf("Expected true but got false")
	}
}

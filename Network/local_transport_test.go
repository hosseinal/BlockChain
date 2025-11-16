package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalTransport_SendMessage(t *testing.T) {
	lt1 := NewLocalTransport("addr1")
	lt2 := NewLocalTransport("addr2")

	err := lt1.Connect(lt2)
	if err != nil {
		t.Fatalf("Failed to connect transports: %v", err)
	}

	err = lt2.Connect(lt1)
	if err != nil {
		t.Fatalf("Failed to connect transports: %v", err)
	}

	message := []byte("Hello, World!")
	err = lt1.SendMessage(lt2.addr, message)

	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	receivedRPC := <-lt2.Consume()
	if string(receivedRPC.Payload) != string(message) {
		t.Fatalf("Expected message %s, got %s", message, receivedRPC.Payload)
	}

	// checked the reviced RPC is similar to the sent one with equal assertion
	assert.Equal(t, receivedRPC.Payload, message)
}

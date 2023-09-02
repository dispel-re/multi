package backend

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientAuthenticationRequest(t *testing.T) {
	// Arrange
	packet := []byte{
		255, 41, // Command code
		19, 0, // Packet length
		2, 0, 0, 0, // Unknown
		112, 97, 115, 115, 0, // Password
		108, 111, 103, 105, 110, 0, // Username
	}

	// Act
	req := ClientAuthenticationRequest(packet[4:])
	unknown, username, password, err := req.Parse()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, uint32(2), unknown)
	assert.Equal(t, "pass", password)
	assert.Equal(t, "login", username)
}

package backend

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccountRequest(t *testing.T) {
	// Arrange
	packet := []byte{
		255, 42, // Command code
		22, 0, // Packet length
		33, 78, 0, 0, // CD-key
		112, 97, 115, 115, 119, 111, 114, 100, 0, // Password
		117, 115, 101, 114, 0, // User name
		0, 0, 49, 207, 69, 0, // Unknown
	}

	// Act
	req := CreateNewAccountRequest(packet[4:])
	cdKey, username, password, unknown, err := req.Parse()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, uint32(20001), cdKey)
	assert.Equal(t, "password", password)
	assert.Equal(t, "user", username)
	assert.True(t, bytes.Equal([]byte{0, 0, 49, 207, 69, 0}, unknown))
}

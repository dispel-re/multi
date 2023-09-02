package backend

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UpdateCharacterSpellsRequest(t *testing.T) {
	// Arrange
	packet := []byte{
		255, 73, // Command code
		58, 0, // Packet length
		117, 115, 101, 114, 0, // User name
		99, 104, 97, 114, 97, 99, 116, 101, 114, 0, // Character name
		2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, // Spells array
	}

	// Act
	req := UpdateCharacterSpellsRequest(packet[4:])
	username, characterName, spells, err := req.Parse()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "user", username)
	assert.Equal(t, "character", characterName)
	assert.Equal(t,
		[]byte{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		spells,
	)
}

func TestUpdateCharacterSpellsRequest_Parse(t *testing.T) {
	t.Run("valid payload", func(t *testing.T) {
		// Arrange
		data := append(
			[]byte("user\x00character\x00"),
			[]byte{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0}...,
		)

		// Act
		req := UpdateCharacterSpellsRequest(data)
		username, characterName, spells, err := req.Parse()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "user", username)
		assert.Equal(t, "character", characterName)
		assert.Equal(t, []byte{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0}, spells)
	})

	t.Run("invalid spells length", func(t *testing.T) {
		// Arrange
		data := []byte("user\x00character\x00badspells")

		// Act
		req := UpdateCharacterSpellsRequest(data)
		username, characterName, spells, err := req.Parse()

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid length")
		assert.Empty(t, username)
		assert.Empty(t, characterName)
		assert.Empty(t, spells)
	})

	t.Run("missing null byte", func(t *testing.T) {
		// Arrange
		data := []byte("usercharacter\x00spells")

		// Act
		req := UpdateCharacterSpellsRequest(data)
		username, characterName, spells, err := req.Parse()

		// Assert
		assert.Error(t, err)
		assert.Empty(t, username)
		assert.Empty(t, characterName)
		assert.Empty(t, spells)
	})
}

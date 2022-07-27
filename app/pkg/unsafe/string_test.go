package unsafe

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBytesToString(t *testing.T) {
	const length = 100
	s := make([]byte, length)
	_, err := rand.Read(s)
	require.NoError(t, err, "read random bytes")
	//
	str := string(s)
	result := BytesToString(s)
	//
	assert.Equal(t, str, result, "compare old bytes with new")
}

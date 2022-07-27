package unsafe

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringToBytes(t *testing.T) {
	const length = 100
	s := make([]byte, length)
	//
	_, err := rand.Read(s)
	require.NoError(t, err, "read random bytes")
	//
	str := string(s)
	result := StringToBytes(&str)
	//
	assert.Equal(t, s, result, "compare old bytes with new")
	assert.Equal(t, cap(s), cap(result), "compare capacity")
}

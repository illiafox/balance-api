package unsafe

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringToBytes(t *testing.T) {
	const length = 1
	s := make([]byte, length)

	_, err := rand.Read(s)
	require.NoError(t, err, "read random bytes")

	str := string(s)
	result := StringToBytes(&str)

	require.Equal(t, s, result, "compare old bytes with new")
	require.Equal(t, cap(s), cap(result), "compare capacity")
}

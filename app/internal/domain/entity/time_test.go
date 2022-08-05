package entity

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTime(t *testing.T) {
	// for marshalling and unmarshalling test
	type parse struct {
		Time Time `json:"time"`
	}

	//
	var now = parse{
		Time: Time{time.Now()},
	}

	data, err := json.Marshal(now)
	require.NoError(t, err, "marshal json")

	var parsed parse
	require.NoError(t, json.Unmarshal(data, &parsed), "unmarshal json")

	require.Equal(t,
		// We can't compare Unix nano/microseconds due to rounding inaccuracy
		// UnixNano() returns nanoseconds since the Unix epoch
		now.Time.Unix(), parsed.Time.Unix(),
		"compare unix time",
	)
}

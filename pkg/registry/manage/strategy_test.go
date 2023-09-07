package manage

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitchLogging(t *testing.T) {
	jf := []byte(`{"feature": "heartbeat", "switch": 1}`)
	err := json.Unmarshal(jf, &SetLoggingSwitchRequest{})
	assert.Nil(t, err)
}
package manage

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitchLogging(t *testing.T) {
	jf := []byte(`{"feature": "heartbeat", "switch": 1}`)
	resp :=  &SetLoggingSwitchRequest{}
	err := json.Unmarshal(jf, resp)
	assert.Nil(t, err)
	t.Log(resp)
}
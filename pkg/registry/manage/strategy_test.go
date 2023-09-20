package manage

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitchLogging(t *testing.T) {
	jf := []byte(`{"feature": "heartbeat", "switch": 0}`)
	resp := &SetLoggingSwitchRequest{}
	err := json.Unmarshal(jf, resp)
	assert.Nil(t, err)
	t.Log(resp)
}

func TestGetConnections(t *testing.T) {
	jf := []byte(`{"Type": 1}`)
	resp := &GetConnectionsRequest{}
	err := json.Unmarshal(jf, resp)
	assert.Nil(t, err)
	t.Log(resp)
}

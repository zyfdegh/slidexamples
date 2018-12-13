package serialization

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYDTextRespUnmarshal(t *testing.T) {
	body := []byte(`{
		"taskId": "079560a6c9f34783bdce47e168510038",
        "action": 2,
        "labels": [
            {
                "label": 100,
                "level": 2,
                "details": {
                    "hint": [
                        "xxx","ooo"
                    ]
                }
            }
        ]
	}`)

	var r YDTextResp
	err := json.Unmarshal(body, &r)
	assert.NoError(t, err)
}

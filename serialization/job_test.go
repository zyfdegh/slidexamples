package serialization

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalCustomizedTrainingRequestJson(t *testing.T) {
	requestJson := `
	{
		"type": "training",
		"priority": 1000.0,
		"remark": "user defefined description",
		"training": {
		"type": "build",
		"workspace": {
			"create": {
				"datasetIds": [ "59f1759e5af792be013532b8", "59fc343c5af792332528e567" ]
			},
			"annotation": "pascolvoc"
		},
		"resources": {
			"ps": {
				"request": {
					"cpu": "200m",
					"mem": "512Mi",
					"gpu": "1",
					"replica": 3
				},
				"limit": {
					"cpu": "4",
					"mem": "4Gi",
					"gpu": "8"
				}
			},
			"worker": {
				"request": {
					"cpu": "1.5",
					"mem": "2Gi"
				},
				"limit": {
					"cpu": "2",
					"mem": "4Gi"
				}
			}
		},
		"buildConfig": {
			"from": "tensorflow/tensorflow:latest-py3",
			"tag":  "gcr.io/my-project/custom:latest",
			"file":	"main.py",
			"entrypoint": "main.py",
			"registryServer" : "asia.gcr.io/linker-aurora"
		}
	},
	"activeDeadlineSeconds": 300,
	"scheduledAt": "2020-01-01T10:00:20.021Z"
	}
	`
	req := JobRequest{}
	err := jsonpb.UnmarshalString(requestJson, &req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", req)

	// field assertions

	training := req.GetTraining()
	assert.NotNil(t, training)

	workspace := training.GetWorkspace()
	assert.NotNil(t, workspace)
	assert.NotNil(t, workspace.GetCreate())
	assert.Equal(t, []string{"59f1759e5af792be013532b8", "59fc343c5af792332528e567"}, workspace.GetCreate().GetDatasetIds())

	resources := training.GetResources()
	assert.NotNil(t, resources)
	assert.Equal(t, "200m", resources["ps"].Request.GetCpu())
	assert.Equal(t, "512Mi", resources["ps"].Request.GetMem())
	assert.Equal(t, "1", resources["ps"].Request.GetGpu())
	assert.Equal(t, uint32(3), resources["ps"].Request.GetReplica())

	assert.Equal(t, "4", resources["ps"].Limit.GetCpu())
	assert.Equal(t, "4Gi", resources["ps"].Limit.GetMem())
	assert.Equal(t, "8", resources["ps"].Limit.GetGpu())

	assert.Equal(t, "1.5", resources["worker"].Request.GetCpu())
	assert.Equal(t, "2Gi", resources["worker"].Request.GetMem())

	assert.Equal(t, "2", resources["worker"].Limit.GetCpu())
	assert.Equal(t, "4Gi", resources["worker"].Limit.GetMem())

	assert.Equal(t, uint64(300), req.GetActiveDeadlineSeconds())
	assert.Equal(t, "2020-01-01T10:00:20.021Z", req.GetScheduledAt())
}

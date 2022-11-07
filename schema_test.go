package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_requestToSchema(t *testing.T) {
	// read testdata into memory
	simpleReq, _ := os.ReadFile("testdata/simple.json")
	simpleGo, _ := os.ReadFile("testdata/simple.go")
	batchJobQueueReq, _ := os.ReadFile("testdata/batch_job_queue.json")
	batchJobQueueGo, _ := os.ReadFile("testdata/batch_job_queue.go")
	ecsClusterReq, _ := os.ReadFile("testdata/ecs_cluster.json")
	ecsClusterGo, _ := os.ReadFile("testdata/ecs_cluster.go")
	ecsServiceReq, _ := os.ReadFile("testdata/ecs_service.json")
	ecsServiceGo, _ := os.ReadFile("testdata/ecs_service.go")
	medialiveChannelReq, _ := os.ReadFile("testdata/medialive_channel.json")
	medialiveChannelGo, _ := os.ReadFile("testdata/medialive_channel.go")
	medialiveMultiplexReq, _ := os.ReadFile("testdata/medialive_multiplex.json")
	medialiveMultiplexGo, _ := os.ReadFile("testdata/medialive_multiplex.go")

	tests := []struct {
		name    string
		b       []byte
		want    []byte
		wantErr bool
	}{
		{"simple", simpleReq, simpleGo, false},
		{"batch job queue", batchJobQueueReq, batchJobQueueGo, false},
		{"ecs cluster", ecsClusterReq, ecsClusterGo, false},
		{"ecs service", ecsServiceReq, ecsServiceGo, false},
		{"medialive channel", medialiveChannelReq, medialiveChannelGo, false},
		{"medialive multiplex", medialiveMultiplexReq, medialiveMultiplexGo, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := requestToSchema(tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("requestToSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestToSchema() = %s, want %s", got, tt.want)
			}
		})
	}
}

func Test_toSnakeCase(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"already snake", "already_snake", "already_snake"},
		{"one word", "Thing", "thing"},
		{"multiple words", "AnotherThing", "another_thing"},
		{"multiple caps start", "HTTPRequest", "http_request"},
		{"multiple caps end", "RequestID", "request_id"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toSnakeCase(tt.s); got != tt.want {
				t.Errorf("toSnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

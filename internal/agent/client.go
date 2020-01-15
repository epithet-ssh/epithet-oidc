package agent

import (
	"fmt"

	"google.golang.org/grpc"
)

//go:generate protoc -I ../../proto agent.proto --go_out=plugins=grpc:.

// NewClient creates  anew RPC client for an Agent
func NewClient(path string) (AgentServiceClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("unix:%s", path), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := NewAgentServiceClient(conn)
	return client, nil
}

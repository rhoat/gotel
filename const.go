package gotel

import (
	"fmt"
	"strings"
)

type Destination int8

const (
	STDOUT Destination = iota + 1
	HTTP
	GRPC
)

var destinationToString = map[Destination]string{
	STDOUT: "stdout",
	HTTP:   "http",
	GRPC:   "grpc",
}

// Destination maps strings to their Destination values.
var stringToDestination = map[string]Destination{
	"stdout": STDOUT,
	"http":   HTTP,
	"grpc":   GRPC,
}

// String implements the fmt.Stringer interface.
func (d *Destination) String() string {
	return destinationToString[*d]
}

func StringToDestination(destination string) (*Destination, error) {
	if val, ok := stringToDestination[strings.ToLower(destination)]; ok {
		return &val, nil
	}
	return nil, fmt.Errorf("invalid OtelDestination: %s", destination)
}

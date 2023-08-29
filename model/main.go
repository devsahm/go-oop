package model

import (
	"fmt"
	"strings"
)

type ipAddress struct {
	ip    string
	parts []string
}

func NewIp(i string) (ipAddress, error) {
	parts := strings.Split(string(i), ".")

	if len(parts) < 4 {
		return ipAddress{}, fmt.Errorf("invalid")
	}

	return ipAddress{i, parts}, nil
}

func GetParts(i ipAddress) []string {
	return i.parts
}

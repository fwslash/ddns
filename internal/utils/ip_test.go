package utils

import (
	"errors"
	"testing"
  "net"

	"github.com/stretchr/testify/assert"
)

func TestGetIps_Success(t *testing.T) {
	// Override resolveIP to avoid real DNS
	resolveIP = func(domain string) ([]net.IP, error) {
		return []net.IP{net.ParseIP("5.6.7.8")}, nil
	}

	mockClient := &MockClient{ResponseBody: "1.2.3.4"}

	ips, err := GetIps("example.com", mockClient)

	assert.NoError(t, err)
	assert.Equal(t, "1.2.3.4", ips.Current)
	assert.Equal(t, "5.6.7.8", ips.Registered)
}

func TestGetIps_CurrentIpFails(t *testing.T) {
	mockClient := &MockClient{Err: errors.New("network error")}

	_, err := GetIps("example.com", mockClient)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "network error")
}

func TestGetIps_RegisteredIpFails(t *testing.T) {
	resolveIP = func(domain string) ([]net.IP, error) {
		return nil, errors.New("dns failure")
	}

	mockClient := &MockClient{ResponseBody: "1.2.3.4"}

	_, err := GetIps("example.com", mockClient)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "dns failure")
}


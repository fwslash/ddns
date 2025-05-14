package utils

import (
	"errors"
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestCheckAndUpdate_WhenIpsDiffer_ShouldReturnUpdated(t *testing.T) {
	mock := &MockClient{ResponseBody: "update-success"}

	config := &DynDnsConfig{
		DynDnsDomain:         "example.com",
		DynDnsUpdateEndpoint: "https://dns.example.com/update",
		DynDnsToken:          "token123",
		CurrentIp:            "1.2.3.4",
		RegisteredIp:         "5.6.7.8",
		Client:               mock,
	}

	result, err := CheckAndUpdate(config)

	assert.NoError(t, err)
	assert.Equal(t, Updated, result)
}

func TestCheckAndUpdate_WhenIpsAreSame_ShouldReturnNoChange(t *testing.T) {
	mock := &MockClient{ResponseBody: "should-not-be-called"}

	config := &DynDnsConfig{
		DynDnsDomain:         "example.com",
		DynDnsUpdateEndpoint: "https://dns.example.com/update",
		DynDnsToken:          "token123",
		CurrentIp:            "1.2.3.4",
		RegisteredIp:         "1.2.3.4",
		Client:               mock,
	}

	result, err := CheckAndUpdate(config)

	assert.NoError(t, err)
	assert.Equal(t, NoChange, result)
}

func TestCheckAndUpdate_WhenUpdateFails_ShouldReturnFailed(t *testing.T) {
	mock := &MockClient{
		Err: errors.New("simulated network error"),
	}

	config := &DynDnsConfig{
		DynDnsDomain:         "example.com",
		DynDnsUpdateEndpoint: "https://dns.example.com/update",
		DynDnsToken:          "token123",
		CurrentIp:            "1.2.3.4",
		RegisteredIp:         "5.6.7.8", 
    Client:               mock,
	}

	result, err := CheckAndUpdate(config)

	assert.Error(t, err)
	assert.Equal(t, Failed, result)
}


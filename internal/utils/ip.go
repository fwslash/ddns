package utils

import (
	"fmt"
	"io"
	"net"
)

var resolveIP = net.LookupIP

type Ips struct {
	Current, Registered string
}


func GetIps(domain string, client HttpClient) (*Ips, error) {
	currentIp, err := getCurrentIp(client)
	if err != nil {
		return nil, err
	}

	registeredIp, err := getRegisteredIp(domain)
	if err != nil {
    fmt.Println(err)
		return nil, err
	}

	return &Ips{
		Current:    currentIp,
		Registered: registeredIp,
	}, nil
}

func getCurrentIp(client HttpClient) (string, error) {
	resp, err := client.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func getRegisteredIp(domain string) (string, error) {
	ip, err := resolveIP(domain)
	if err != nil {
		return "", err
	}

	return ip[0].String(), err
}

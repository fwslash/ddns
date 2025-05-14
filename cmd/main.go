package main

import (
	"context"
	"fmt"
	"github.com/fwslash/ddns/internal/utils"
	"os"
	"os/signal"
	"syscall"
	"time"
  "log"
)

func main() {
	client := utils.Client()
	dynDnsDomain := os.Getenv("DYNDNS_NAME")

  ips, err := utils.GetIps(dynDnsDomain, client)
  if err != nil {
    log.Fatalf("Failed to get IPs: %v", err)
  }

	config := &utils.DynDnsConfig{
		DynDnsDomain:         dynDnsDomain,
		DynDnsUpdateEndpoint: os.Getenv("DYNDNS_UPDATE_ENDPOINT"),
		DynDnsToken:          os.Getenv("DYNDNS_TOKEN"),
		CurrentIp:            ips.Current,
		RegisteredIp:         ips.Registered,
		Client:               client,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("Received shutdown signal, exiting...")
		cancel()
	}()

	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	utils.CheckAndUpdate(config)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context canceled, shutting down loop.")
			return
		case <-ticker.C:
			utils.CheckAndUpdate(config)
		}
	}
}

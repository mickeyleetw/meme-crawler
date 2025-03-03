package core

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

// HTTPClientPool returns a configured http.Client with custom timeout and transport settings
func HTTPClientPool() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 30 * time.Second,
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				fmt.Printf("🔌 Connecting to: %s\n", addr)
				dialer := &net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}
				return dialer.DialContext(ctx, network, addr)
			},
		},
	}
}

// HTTPClientPool returns a configured http.Client with custom timeout and transport settin

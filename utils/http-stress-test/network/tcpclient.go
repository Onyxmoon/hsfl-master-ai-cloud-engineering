package network

import (
	"fmt"
	"net"
	"net/url"
)

type TCPClient struct{}

func NewTcpClient() *TCPClient {
	return &TCPClient{}
}

func (c *TCPClient) Send(targetURL string) error {
	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		return err
	}
	target := parsedURL.Host
	path := parsedURL.Path
	if path == "" {
		path = "/"
	}

	conn, err := net.Dial("tcp", target)
	if err != nil {
		return err
	}
	request := fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\n\r\n", path, target)

	_, err = conn.Write([]byte(request))
	if err != nil {
		return err
	}

	return nil
}

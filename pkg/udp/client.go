package udp

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Client interface {
	Send(b []byte) error
}

type client struct {
	mux     sync.Mutex
	conn    *net.UDPConn
	address string
}

func (c *client) Send(b []byte) error {
	if c.conn == nil {
		return nil
	}

	c.mux.Lock()
	//defer time.Sleep(1 * time.Millisecond)
	defer c.mux.Unlock()

	_, err := c.conn.Write(b)
	if err != nil {
		go c.tryConnect()

		return err
	}

	return nil
}

func (c *client) Close() error {
	if c.conn == nil {
		return nil
	}

	return c.conn.Close()
}

func NewClient(address string) (*client, error) {
	c := &client{
		conn:    nil,
		address: address,
	}

	go c.tryConnect()

	return c, nil
}

// TODO: mutex on c.conn
func (c *client) tryConnect() {
	for {
		err := c.connect()
		if err != nil {
			time.Sleep(3 * time.Second)
			continue
		}

		break
	}
}

func (c *client) connect() error {
	addr, err := net.ResolveUDPAddr("udp", c.address)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}

	c.conn = conn
	fmt.Println("resolved")
	return nil
}

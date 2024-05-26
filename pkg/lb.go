package lb

import (
	"io"
	"log"
	"net"
)

type LoadBalancer struct {
	Config        *LBConfig
	healthChecker *HealthChecker
	selector      *Selector
}

func NewLoadBalancer(config *LBConfig) *LoadBalancer {
	return &LoadBalancer{
		Config:        config,
		selector:      NewSelector(),
		healthChecker: NewHealthChecker(config.Servers, config.HealthCheckRoute, config.HealthCheckInterval),
	}
}

func (s *LoadBalancer) Start() {
	ln, err := net.Listen("tcp", s.Config.Address)
	if err != nil {
		log.Panicf("Could not start server: %+v", err)
	}

	s.healthChecker.Run()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error handling connection: %+v", err)
		}

		go s.handleConnection(conn)
	}
}

func (s *LoadBalancer) handleConnection(conn net.Conn) {
	buf := make([]byte, s.Config.BufferSize)
	var data []byte

	// Read request data
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == net.ErrClosed {
				log.Println("Connection closed from client")
			} else if err == io.EOF {
				log.Println("End of data from client")
			} else {
				log.Printf("Error reading data: %+v", err)
			}
			break
		}
		data = append(data, buf[:n]...)
		if n < int(s.Config.BufferSize) {
			break
		}
	}

	// Forward request data to one of the application servers
	res, err := s.routeRequest(data)
	if err != nil {
		log.Printf("Got error while receiving response: %+v", err)
	}

	// Send response back to client
	conn.Write(res)
}

func (s *LoadBalancer) routeRequest(data []byte) ([]byte, error) {
	serverAddress := s.selectServer()
	log.Printf("Forwarding request to: %s", serverAddress)

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		return nil, err
	}
	_, err = conn.Write(data)
	if err != nil {
		return nil, err
	}

	var response []byte
	buf := make([]byte, s.Config.BufferSize)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == net.ErrClosed || err == io.EOF {
				log.Printf("connection closed by the server")
			} else {
				log.Printf("got error while getting response from server: %+v", err)
			}
			break
		}

		response = append(response, buf[:n]...)
		if n < int(s.Config.BufferSize) {
			break
		}
	}
	return response, nil
}

func (s *LoadBalancer) selectServer() string {
	servers := s.healthChecker.GetHealthyServers()
	idx := s.selector.Next() % uint64(len(servers))
	return servers[idx]
}

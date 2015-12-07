package network

type server struct {

	// the hostname of the server
	host string

	// ports is a map of ports on this server.
	ports map[int]bool

	on bool

	// validates ports based on their membership in the pool.
	poolValidator
}

//NewServer returns a server with the specified ports
func NewServer(host string, ports map[int]bool) *server {
	// If this is a new server, start the pool with it's ports.
	if portPool == nil {
		portPool = ports
	} else {
		for port, aval := range ports {
			// we intentionally don't check if port exists for ... reasons.
			portPool[port] = aval
		}
	}

	s := &server{ports: ports}
	s.start()
	return s
}

//start the server
func (s *server) start() {
	if s.host == "" {
		s.host = "localhost"
	}
	s.on = true
}

var portPool map[int]bool

//Ports returns a map of port addresses to booleans signifiying if the ports
//are avaliable. The returned ports are first validated.
func (s *server) Ports() map[int]bool {
	return s.poolValidator.validPorts(s.ports)
}

//Address returns the address of the server.
func (s *server) Address(port string) string {
	return s.host + ":" + port
}

type poolValidator struct {
	inPool bool
}

//validPorts takes a map of ports and returns those that are valid.
func (p poolValidator) validPorts(ports map[int]bool) map[int]bool {
	defer func() {
		if !p.inPool {
			for port := range ports {
				delete(portPool, port)
			}
		}
	}()

	validPorts := map[int]bool{}
	for port, aval := range ports {
		if port >= 1 && port <= 65535 {
			validPorts[port] = aval
		}
	}

	return validPorts
}

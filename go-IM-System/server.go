package main

type Server struct {
	ip   string
	port int
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		ip:   ip,
		port: port,
	}
	return server
}

func (server Server) Start() {
	// socket listen

	// accept

	// do handler

	// close listen socker
}

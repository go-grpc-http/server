package server

func (s *server) WithDualRegisterer(regs ...Registry) {
	// TODO: need to add flags
	for index := range regs {
		s.registeredGrpcHandlers = append(s.registeredGrpcHandlers, regs[index])
		s.registeredHttpHandlers = append(s.registeredHttpHandlers, regs[index])
	}
}

func (s *server) WithGrpcRegisterer(regs ...GrpcRegisterer) {
	s.registeredGrpcHandlers = append(s.registeredGrpcHandlers, regs...)
}

func (s *server) WithHttpRegisterer(regs ...HttpRegisterer) {
	s.registeredHttpHandlers = append(s.registeredHttpHandlers, regs...)
}

func (s *server) WithUnaryServerInterceptor(regs ...HttpRegisterer) {
	s.registeredHttpHandlers = append(s.registeredHttpHandlers, regs...)
}

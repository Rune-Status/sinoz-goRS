package login

import "math/rand"

type Service struct {
	// TODO
}

func (s *Service) GenerateSessionKey() int64 {
	return rand.Int63()
}
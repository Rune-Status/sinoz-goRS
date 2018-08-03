package game

type Service struct {
	world *World
}

func NewService() *Service {
	return &Service{world: NewWorld()}
}

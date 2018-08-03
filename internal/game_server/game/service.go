package game

type Service struct {
	Assets Assets
	World  *World
}

func NewService(assets Assets) *Service {
	return &Service{
		Assets: assets,
		World: NewWorld(),
	}
}

package adding

type service struct{}

// Service is our adding service interface
type Service interface {
	AddEggMoves(...EggMove)
}

// Repository is our respository interface
type Repository interface {
	AddEggMove()
}

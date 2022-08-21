package listing

const (
	// Weak represents a type a pokemon is weak against
	Weak = iota - 1
	// NoEffect means 0 damage
	NoEffect
	// Effective means it does damage
	Effective
	// SuperEffective means it does a lot of damage
	SuperEffective
)

// TypeEffectiveNess represents our table
type TypeEffectiveNess struct {
	ID            int
	PokemonTypeID int
	TypeName      string
	PokemonName   string
	Against       string
	Score         int
}

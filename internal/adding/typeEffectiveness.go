package adding

const (
	// SuperWeak is a type that really hurts a pokemo
	SuperWeak = iota - 1
	// Weak is a type that hurts a pokemon
	Weak
	// NoEffect has no effect against a pokmeon
	NoEffect
	// Effective is a type that a pokemon is effective against
	Effective
	// SuperEffective is a type that a pokemon is
	SuperEffective
)

// TypeEffectiveNess represents a table entry
type TypeEffectiveNess struct {
	ID            int
	PokemonTypeID int
	Against       string
	Score         int
}

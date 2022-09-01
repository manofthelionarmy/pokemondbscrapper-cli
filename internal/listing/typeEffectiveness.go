package listing

const (
	// SuperWeak is a type that really hurts a pokemo
	SuperWeak = iota - 1

	// NotVeryEffective means an attack does damage, but not a lot
	NotVeryEffective

	// Weak is a type that hurts a pokemon and does 1x damage
	Weak

	// NoEffect has no effect against a pokmeon and does zero damage
	NoEffect

	// Effective is a type that a pokemon is effective against
	Effective
	// SuperEffective is a type that a pokemon is
	SuperEffective
)

const (
	// Normal is our normal type
	Normal = "normal"
	// Fighting is our fighting type
	Fighting = "fighting"
	// Rock is our rock type
	Rock = "rock"
	// Fire is our fire type
	Fire = "fire"
	// Poison is our poison type
	Poison = "poison"
	// Ghost is our ghost type
	Ghost = "ghost"
	// Water is our water type
	Water = "water"
	// Ground is our ground type
	Ground = "ground"
	// Dark is our dark type
	Dark = "dark"
	// Grass is our grass type
	Grass = "grass"
	// Flying is our flying type
	Flying = "flying"
	// Dragon is our dragon type
	Dragon = "dragon"
	// Electric is our electric type
	Electric = "electric"
	// Psychic is our psychic type
	Psychic = "psychic"
	// Steel is our steel type
	Steel = "steel"
	// Ice is our ice type
	Ice = "ice"
	// Bug is our bug type
	Bug = "bug"
	// Fairy is our fairy type
	Fairy = "fairy"
)

// TypeEffectiveNess represents a table entry
type TypeEffectiveNess struct {
	ID           int    `field:"type_effectiveness_id"`
	TypeName     string `field:"type_name"`
	AgainstType  string `field:"against_tpye"`
	AttackScore  int    `field:"attack_score"`
	DefenseScore int    `field:"defense_score"`
}

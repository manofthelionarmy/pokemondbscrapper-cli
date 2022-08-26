package sqlite

import "testing"

func TestSqlite(t *testing.T) {
	for scenario, f := range map[string]func(t *testing.T){
		"adding entries into type effectiveness": testTypeEffectivess,
	} {
		t.Run(scenario, f)
	}
}

func testTypeEffectivess(t *testing.T) {
	db := NewBuilder().WithDataSource("pokemon.db").Build()
	db.TypeEffectiveNess()
}

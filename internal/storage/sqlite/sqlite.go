package sqlite

import "database/sql"

// Sqlite is our repo implementation
type Sqlite struct {
	db *sql.DB
}

// AddEggMove adds an eggmove entry
func (sqlite *Sqlite) AddEggMove() {

}

// AddPokemon adds Pokemon entry
func (sqlite *Sqlite) AddPokemon() {

}

// AddType adds type entry
func (sqlite *Sqlite) AddType() {

}

// AddTypeEffectivenessEntry adds an entry to the type effectiveness table
func (sqlite *Sqlite) AddTypeEffectivenessEntry() {

}

package sqlcontent

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/manofthelionarmy/pokemondbscrapper/pkg/listing"
	"github.com/stretchr/testify/require"
)

func TestSQLContent(t *testing.T) {
	for scenario, f := range map[string]func(*testing.T){
		"can create a table":        testCreateTable,
		"can store pokemon listing": testStorePokemonListing,
		"can store move listing":    testStoreMoveListing,
		"can store eggmove listing": testStoreEggMove,
		"can store moveset listing": testListingMoveset,
	} {
		t.Run(scenario, f)
	}
}

func testCreateTable(t *testing.T) {
	temp, _ := ioutil.TempFile("", "test")
	f, err := os.OpenFile(temp.Name(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	require.NoError(t, err)
	defer os.Remove(temp.Name())
	defer temp.Close()

	sc := NewSQLContent(f)
	err = sc.CreateTable("dummy")
	require.NoError(t, err)

	err = sc.Flush()
	require.NoError(t, err)
	fi, err := sc.SQLFile.Stat()

	require.NoError(t, err)
	require.Greater(t, fi.Size(), int64(0))
}

func testStorePokemonListing(t *testing.T) {
	pokemon := listing.AllPokemon()
	temp, _ := ioutil.TempFile("", "test")
	f, err := os.OpenFile(temp.Name(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	require.NoError(t, err)
	defer os.Remove(temp.Name())
	defer temp.Close()

	sc := NewSQLContent(f)
	err = sc.InsertInto("dummy", pokemon)
	require.NoError(t, err)

	err = sc.Flush()
	require.NoError(t, err)

	require.NoErrorf(t, err, "%+v", pokemon)
	fi, err := sc.SQLFile.Stat()
	require.NoError(t, err)

	require.NotZero(t, fi.Size())
}

func testStoreMoveListing(t *testing.T) {
	moves := listing.AllPossibleMoves()
	temp, _ := ioutil.TempFile("", "test")
	f, err := os.OpenFile(temp.Name(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	require.NoError(t, err)
	defer os.Remove(temp.Name())
	defer temp.Close()

	sc := NewSQLContent(f)
	err = sc.InsertInto("moves", moves)
	require.NoError(t, err)

	err = sc.Flush()
	require.NoError(t, err)

	fi, err := sc.SQLFile.Stat()
	require.NotZero(t, fi.Size())
}

func testStoreEggMove(t *testing.T) {
	// TODO: benchmark test, we have to do multiple fetches of eggmoves for every pokemon
	// Find out if there's a better way to get a pokemon's eggmoves...
	// O (P * E)
	// Could Represent egg move per pokemon as a graph: root vertex is egg move while neighbors are pokemon
	// How many root vertexs will we have? The number of root vertexes we'll have is O(E), therfore
	// finding all of the Pokemon with an eggmoves is O(V+E) and therofore O( E * (V + E)) (exploring all neighbors for each disjoint graph...)
	eggMoves := listing.PokemonEggMoves("bulbasaur")
	temp, _ := ioutil.TempFile("", "test")
	f, err := os.OpenFile(temp.Name(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	require.NoError(t, err)
	defer os.Remove(temp.Name())
	defer temp.Close()

	sc := NewSQLContent(f)
	err = sc.InsertInto("eggMoves", eggMoves, Options{PokemonName: "Bulbasaur"})
	require.NoError(t, err)

	err = sc.Flush()
	require.NoError(t, err)

	fi, err := sc.SQLFile.Stat()
	require.NotZero(t, fi.Size())
}

func testListingMoveset(t *testing.T) {
	moveset := listing.PokemonMoveset("bulbasaur")
	temp, _ := ioutil.TempFile("", "test")
	f, err := os.OpenFile(temp.Name(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	require.NoError(t, err)
	defer os.Remove(temp.Name())
	defer temp.Close()

	sc := NewSQLContent(f)
	err = sc.InsertInto("movest", moveset, Options{PokemonName: "Bulbasaur"})
	require.NoError(t, err)

	err = sc.Flush()
	require.NoError(t, err)
	fi, err := sc.SQLFile.Stat()
	require.NotZero(t, fi.Size())
}

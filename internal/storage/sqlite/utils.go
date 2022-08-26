package sqlite

import (
	"database/sql"

	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/adding"
)

func appendGrassType(typeEffectivenessChart []adding.TypeEffectiveNess) []adding.TypeEffectiveNess {
	typeEffectivenessChart = append(typeEffectivenessChart, []adding.TypeEffectiveNess{
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Normal,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Weak,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Fire,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.SuperWeak,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Water,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Electric,
			AttackScore:  adding.Effective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Grass,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Ice,
			AttackScore:  adding.Effective,
			DefenseScore: adding.SuperWeak,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Poison,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.SuperWeak,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Ground,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Flying,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.SuperWeak,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Psychic,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Bug,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.SuperWeak,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Rock,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Ghost,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Dragon,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Dark,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Steel,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Fairy,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Grass,
			AgainstType:  adding.Fighting,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
	}...)
	return typeEffectivenessChart
}

func appendNormalType(typeEffectivenessChart []adding.TypeEffectiveNess) []adding.TypeEffectiveNess {
	typeEffectivenessChart = append(typeEffectivenessChart, []adding.TypeEffectiveNess{
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Normal,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Fire,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Water,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Electric,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Grass,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Ice,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Fighting,
			AttackScore:  adding.Effective,
			DefenseScore: adding.SuperWeak,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Poison,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Ground,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Flying,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Psychic,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Bug,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Rock,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Ghost,
			AttackScore:  adding.NoEffect,
			DefenseScore: adding.NoEffect,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Dragon,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Dark,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Steel,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Normal,
			AgainstType:  adding.Fairy,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
	}...)
	return typeEffectivenessChart
}

func appendFireType(typeEffectivenessChart []adding.TypeEffectiveNess) []adding.TypeEffectiveNess {
	typeEffectivenessChart = append(typeEffectivenessChart, []adding.TypeEffectiveNess{
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Normal,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Fire,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Water,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.SuperWeak,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Electric,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Grass,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Ice,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Fighting,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Poison,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Ground,
			AttackScore:  adding.Effective,
			DefenseScore: adding.SuperWeak,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Flying,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Psychic,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Bug,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Rock,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.SuperWeak,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Ghost,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Dragon,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Dark,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Steel,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Fire,
			AgainstType:  adding.Fairy,
			AttackScore:  adding.Effective,
			DefenseScore: adding.NotVeryEffective,
		},
	}...)
	return typeEffectivenessChart
}

func checkTypeEntriesAdded(db *sql.DB, typeName string) bool {

	checkEntriesAlreadyAdded, _ := db.Prepare("SELECT COUNT(*) FROM type_effectiveness WHERE type_name=$1;")

	result, err := checkEntriesAlreadyAdded.Query(typeName)
	if err != nil {
		panic(err)
	}

	var count int
	for result.Next() {
		err := result.Scan(&count)
		if err != nil {
			panic(err)
		}
	}
	// there are 18 types in total, so there should be 18 entries per type
	if count < 18 {
		return false
	}
	// for result.Next() {
	// 	typeEffectiveness := adding.TypeEffectiveNess{}
	// 	s := reflect.ValueOf(&typeEffectiveness).Elem()
	// 	numColumns := s.NumField()
	// 	columns := make([]interface{}, numColumns)
	// 	for i := 0; i < numColumns; i++ {
	// 		field := s.Field(i)
	// 		columns[i] = field.Addr().Interface() // get the address of each field
	// 	}
	// 	// I love this trick!!!
	// 	result.Scan(columns...) // populate each one
	// 	fmt.Println(typeEffectiveness)
	// }
	return true
}

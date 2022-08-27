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

func appendWaterType(typeEffectivenessChart []adding.TypeEffectiveNess) []adding.TypeEffectiveNess {
	typeEffectivenessChart = append(typeEffectivenessChart, []adding.TypeEffectiveNess{
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Normal,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Fire,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Water,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Electric,
			AttackScore:  adding.Effective,
			DefenseScore: adding.SuperWeak,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Grass,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.SuperWeak,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Ice,
			AttackScore:  adding.Effective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Fighting,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Poison,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Ground,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Flying,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Psychic,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Bug,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Rock,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Ghost,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Dragon,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Dark,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Steel,
			AttackScore:  adding.Effective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Water,
			AgainstType:  adding.Fairy,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
	}...)
	return typeEffectivenessChart
}

func appendElectricType(typeEffectivenessChart []adding.TypeEffectiveNess) []adding.TypeEffectiveNess {
	typeEffectivenessChart = append(typeEffectivenessChart, []adding.TypeEffectiveNess{
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Normal,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Fire,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Water,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Electric,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Grass,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Ice,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Fighting,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Poison,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Ground,
			AttackScore:  adding.NoEffect,
			DefenseScore: adding.SuperWeak,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Flying,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Psychic,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Bug,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Rock,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Ghost,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Dragon,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Dark,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Steel,
			AttackScore:  adding.Effective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Electric,
			AgainstType:  adding.Fairy,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
	}...)
	return typeEffectivenessChart
}

func appendFightingType(typeEffectivenessChart []adding.TypeEffectiveNess) []adding.TypeEffectiveNess {
	typeEffectivenessChart = append(typeEffectivenessChart, []adding.TypeEffectiveNess{
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Normal,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Fire,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.SuperWeak,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Water,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Electric,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Grass,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Ice,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Fighting,
			AttackScore:  adding.Effective,
			DefenseScore: adding.SuperEffective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Poison,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Ground,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Flying,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Psychic,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Bug,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Rock,
			AttackScore:  adding.Effective,
			DefenseScore: adding.SuperEffective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Ghost,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Dragon,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Dark,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Steel,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.SuperWeak,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Fairy,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
	}...)
	return typeEffectivenessChart
}

func appendIceType(typeEffectivenessChart []adding.TypeEffectiveNess) []adding.TypeEffectiveNess {
	typeEffectivenessChart = append(typeEffectivenessChart, []adding.TypeEffectiveNess{
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Normal,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Fire,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.SuperWeak,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Water,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Electric,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Grass,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Ice,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.NotVeryEffective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Fighting,
			AttackScore:  adding.Effective,
			DefenseScore: adding.SuperEffective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Poison,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Ground,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Flying,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Psychic,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Bug,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Rock,
			AttackScore:  adding.Effective,
			DefenseScore: adding.SuperEffective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Ghost,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Dragon,
			AttackScore:  adding.SuperEffective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Dark,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Steel,
			AttackScore:  adding.NotVeryEffective,
			DefenseScore: adding.SuperWeak,
		},
		{
			TypeName:     adding.Ice,
			AgainstType:  adding.Fairy,
			AttackScore:  adding.Effective,
			DefenseScore: adding.Effective,
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

func populateTypeEffectiveness(db *sql.DB, typeName string) {
	// Check if the type has already been added
	if checkTypeEntriesAdded(db, typeName) {
		return
	}

	var f func([]adding.TypeEffectiveNess) []adding.TypeEffectiveNess
	switch typeName {
	case adding.Normal:
		f = appendNormalType
	case adding.Fire:
		f = appendFireType
	case adding.Water:
		f = appendWaterType
	case adding.Electric:
		f = appendElectricType
	case adding.Grass:
		f = appendGrassType
	case adding.Ice:
		f = appendIceType
	case adding.Fighting:
		f = appendIceType
	}

	typeEffectivenessChart := make([]adding.TypeEffectiveNess, 0)
	typeEffectivenessChart = f(typeEffectivenessChart)
	for _, entry := range typeEffectivenessChart {
		stmt, err := db.Prepare(`INSERT INTO type_effectiveness (type_name, against_type, attack_score, defense_score)
			VALUES ($1, $2, $3, $4);
		`)
		if err != nil {
			panic(err)
		}
		if _, err := stmt.Exec(entry.TypeName, entry.AgainstType, entry.AttackScore, entry.DefenseScore); err != nil {
			panic(err)
		}
	}
}

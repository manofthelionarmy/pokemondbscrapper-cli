package sqlite

import (
	"database/sql"
)

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

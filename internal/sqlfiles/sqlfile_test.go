package sqlfiles

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSQLFile(t *testing.T) {
	for scenario, f := range map[string]func(*testing.T){
		"testNewSQLFile": testNewSQLFile,
	} {
		t.Run(scenario, f)
	}
}

func testNewSQLFile(t *testing.T) {
	tempFile, _ := os.CreateTemp("temp", "")
	defer os.Remove("temp")
	defer tempFile.Close()
	require.NotNil(t, NewSQLFILE(tempFile))
}

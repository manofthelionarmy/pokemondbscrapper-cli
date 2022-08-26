package webscraper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuilder(t *testing.T) {
	for scenario, f := range map[string]func(*testing.T){
		"test builder": testBuilder,
	} {
		t.Run(scenario, f)
	}
}

func testBuilder(t *testing.T) {
	builder := NewBuilder()
	require.NotNil(t, builder)

	builder.WithURL("test")
	require.Equal(t, builder.webScaper.url, "test")
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/adding"
	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/listing"
	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/storage/sqlite"
	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/storage/webscraper"
	"github.com/spf13/cobra"
)

// pokemonTypesCmd represents the pokemonTypes command
var pokemonTypesCmd = &cobra.Command{
	Use:   "pokemonTypes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		scraper := webscraper.NewBuilder().WithURL("https://pokemondb.net").Build()
		db := sqlite.NewBuilder().WithDataSource("pokemon.db").Build()
		listingSvc := listing.NewService(scraper)

		addingSvc := adding.NewService(db)
		pokemon := listingSvc.AllPokemon()
		addingSvc.PokemonType(pokemon)
	},
}

func init() {
	rootCmd.AddCommand(pokemonTypesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pokemonTypesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pokemonTypesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

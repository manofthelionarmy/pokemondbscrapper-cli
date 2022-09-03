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

// movesetCmd represents the moveset command
var movesetCmd = &cobra.Command{
	Use:   "moveset",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		webScraper := webscraper.NewBuilder().
			WithURL("https://pokemondb.net").
			Build()

		listingSvc := listing.NewService(webScraper)

		db := sqlite.NewBuilder().
			WithDataSource("pokemon.db").
			Build()

		// adding service adds entries to the db
		addingSvc := adding.NewService(db)

		pokemon := listingSvc.AllPokemon()

		for _, p := range pokemon {
			moves := listingSvc.Moveset(p.Name)
			addingSvc.Moveset(p.PokedexNo, moves)
		}
	},
}

func init() {
	rootCmd.AddCommand(movesetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// movesetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// movesetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

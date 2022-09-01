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

// typeEffectivenessCmd represents the typeEffectiveness command
var typeEffectivenessCmd = &cobra.Command{
	Use:   "typeEffectiveness",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := sqlite.NewBuilder().
			WithDataSource("pokemon.db").
			Build()

		scraper := webscraper.NewBuilder().
			WithURL("https://pokemondb.net").
			Build()
		listingSvc := listing.NewService(scraper)

		typeEffectivness := listingSvc.TypeEffectiveNess()
		addingSvc := adding.NewService(db)
		addingSvc.TypeEffectiveNess(typeEffectivness)
	},
}

func init() {
	rootCmd.AddCommand(typeEffectivenessCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// typeEffectivenessCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// typeEffectivenessCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

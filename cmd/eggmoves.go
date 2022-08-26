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

// eggmovesCmd represents the eggmoves command
var eggmovesCmd = &cobra.Command{
	Use:   "eggmoves",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		webScraper := webscraper.NewBuilder().WithURL("").Build()
		listingService := listing.NewService(webScraper)
		eggMoves := listingService.EggMoves("bulbasuar")

		db := sqlite.NewBuilder().
			WithDataSource("").
			Build()

		addingService := adding.NewService(db)
		addingService.EggMoves(0, eggMoves)
	},
}

func init() {
	rootCmd.AddCommand(eggmovesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// eggmovesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// eggmovesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

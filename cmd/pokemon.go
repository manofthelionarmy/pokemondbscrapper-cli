/*
Copyright © 2022 Armando Ramon Leon Jr <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/adding"
	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/listing"
	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/storage/sqlite"
	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/storage/webscraper"
	"github.com/spf13/cobra"
)

// pokemonCmd represents the pokemon command
var pokemonCmd = &cobra.Command{
	Use:   "pokemon",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		scraper := webscraper.NewBuilder().WithURL("https://pokemondb.net").Build()
		listingSvc := listing.NewService(scraper)

		datasource, err := cmd.Flags().GetString("db")
		if err != nil {
			log.Fatal(err)
		}

		if len(datasource) == 0 {
			log.Fatal("empty value passed into db")
		}

		db := sqlite.NewBuilder().
			WithDataSource(datasource).
			Build()

		addingSvc := adding.NewService(db)
		pokemon := listingSvc.AllPokemon()
		addingSvc.Pokemon(pokemon)
	},
}

func init() {
	rootCmd.AddCommand(pokemonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pokemonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pokemonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

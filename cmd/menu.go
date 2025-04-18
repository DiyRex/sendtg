package cmd

import (
	"github.com/spf13/cobra"
	"sendtg/internal/ui"
)

var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "Launch interactive menu",
	Run: func(cmd *cobra.Command, args []string) {
		ui.StartMenu()
	},
}

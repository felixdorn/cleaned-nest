package cli

import (
	"fmt"

	"github.com/me/nest/common"
	"github.com/spf13/cobra"
)

func DiagnoseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "diagnose",
		Short: "diagnose your configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			diagnosis := common.DiagnoseConfiguration()

			fmt.Printf("\n  \033[1m\033[38;2;255;60;0mErrors:\033[0m\n")

			if len(diagnosis.Errors) == 0 {
				fmt.Println("  \033[1m\033[38;2;120;120;120m- no errors\033[0m")
			} else {
				for _, err := range diagnosis.Errors {
					fmt.Printf("  \033[1m\033[38;2;255;255;255m- %s\033[0m\n", err.Title)
					if err.Error != nil {
						fmt.Printf("  \033[1m\033[38;2;125;125;125m  %s\033[0m\n", err.Error.Error())
					}
				}
			}

			fmt.Printf("\n  \033[1m\033[38;2;250;175;0mRecommendations:\033[0m\n")

			if len(diagnosis.Recommendations) == 0 {
				fmt.Println("  \033[1m\033[38;2;120;120;120m- no recommendations\033[0m")
			} else {
				for _, warn := range diagnosis.Recommendations {
					fmt.Printf("  \033[1m\033[38;2;255;255;255m- %s\033[0m\n", warn.Title)
					if warn.Details != "" {
						fmt.Printf("  \033[1m\033[38;2;125;125;125m  %s\033[0m\n", warn.Details)
					}
				}
			}

			return nil
		},
	}

	return cmd
}

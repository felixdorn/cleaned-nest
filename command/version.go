package command

import (
	"fmt"

	"github.com/redwebcreation/nest/global"
	"github.com/spf13/cobra"
)

func runVersionCommand(_ *cobra.Command, _ []string) error {
	fmt.Printf("nest@%s\n", global.Version)
	return nil
}

func VersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "print nest's version",
		RunE:  runVersionCommand,
	}

	return cmd
}

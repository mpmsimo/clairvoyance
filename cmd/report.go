package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"clairvoyance/app/reporting"
)

var reportCmd = &cobra.Command{
	Use:   "Clairvoyance",
	Short: "Terraform state drift detection and reporting.",
    Long: `Usage:
            clairvoyance report.`,
            Run: func(cmd *cobra.Command, args []string) {
            	reporting.SendReport()
			},
}
func init() {
	fmt.Println("cmd/report/go running.")
	rootCmd.AddCommand(reportCmd)
}

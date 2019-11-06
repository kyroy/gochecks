package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/kyroy/gochecks/pkg/gotest"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	testCmd.AddCommand(testPRCmd)
	testPRCmd.Flags().StringVar(&directory, "dir", ".", "path to the test directory")
	// TODO ideas:
	// - coverage: enforces a minimum coverage
	// - package-coverage: enforces a minimum coverage per package
	//   option to only warn/notify if lower
}

var testPRCmd = &cobra.Command{
	Use:   "pr",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		// write test results to file for later steps
		f, err := os.Open(TestFile)
		if err != nil {
			return fmt.Errorf("failed to open file '%s': %v", TestFile, err)
		}
		defer f.Close()

		// read tests from previous steps
		var results gotest.Results
		if err := json.NewDecoder(f).Decode(&results); err != nil {
			return fmt.Errorf("failed to decode tests from file: %v", err)
		}

		return nil
	},
}

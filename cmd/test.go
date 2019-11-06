package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kyroy/gochecks/pkg/gotest"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

const TestFile = "test.gochecks.json"

var (
	input     string
	directory string
	fail bool
)

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().StringVarP(&input, "input", "i", "", "path to test log file (can also be passed via pipe)")
	testCmd.Flags().StringVar(&directory, "dir", ".", "path to the test directory")
	testCmd.Flags().BoolVar(&fail, "fail", false, "if set, the command will fail on failed tests")
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		info, err := os.Stdin.Stat()
		if err != nil {
			return fmt.Errorf("failed to get stdin status: %v", err)
		}

		var reader io.Reader
		if input != "" {
			reader, err = os.Open(input)
			if err != nil {
				return fmt.Errorf("failed to open input: %v", err)
			}
		} else if info.Size() > 0 {
			reader = bufio.NewReader(os.Stdin)
		} else {
			res, err := gotest.Run(directory)
			if err != nil {
				if !bytes.HasPrefix(res, []byte("{")) {
					return fmt.Errorf("go test: %s\n%v", bytes.TrimSpace(res), err)
				}
			}
			reader = bytes.NewReader(res)
		}

		// parse test results for proper handling
		results, err := gotest.Parse(reader)
		if err != nil {
			fmt.Println(string(input))
			return fmt.Errorf("failed to parse go tests: %v", err)
		}

		// write test results to file for later steps
		f, err := os.OpenFile(TestFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to open file '%s': %v", TestFile, err)
		}
		defer f.Close()
		if err := json.NewEncoder(f).Encode(results); err != nil {
			return fmt.Errorf("failed to encode tests to file: %v", err)
		}

		// print failed tests
		var failed []string
		for _, r := range results {
			if r.Result == "fail" {
				failed = append(failed, r.String())
			}
		}
		if len(failed) > 0 {
			msg := fmt.Sprintf("some tests failed:\n%s", strings.Join(failed, "\n"))
			if fail {
				return fmt.Errorf(msg)
			} else {
				fmt.Println(msg)
			}
		} else {
			fmt.Println("tests successful")
		}
		return nil
	},
}

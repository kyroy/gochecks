package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func init() {
	rootCmd.AddCommand(fmtCmd)
	//testCmd.Flags().StringVarP(&input, "input", "i", "", "path to test log file (can also be passed via pipe)")
	fmtCmd.Flags().StringVar(&directory, "dir", ".", "path to the test directory")
}

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		var errors []string
		err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() || !strings.HasSuffix(info.Name(), ".go") {
				return nil
			}
			body, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			res, err := format.Source(body)
			if err != nil {
				return err
			}
			if !reflect.DeepEqual(body, res) {
				errors = append(errors, fmt.Sprintf("%s", path))
			}
			return nil
		})
		if err != nil {
			return fmt.Errorf("failed to traverse files: %v", err)
		}
		if len(errors) > 0 {
			return fmt.Errorf("go fmt: Following files are not formatted:\n%s", strings.Join(errors, "\n"))
		}
		return nil
	},
}

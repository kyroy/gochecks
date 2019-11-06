package gotest_test

import (
	"encoding/json"
	"github.com/kyroy/gochecks/pkg/gotest"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

const update = false

func testConverterParse(t *testing.T, input io.Reader, golden string) {
	result, err := gotest.Parse(input)
	require.NoError(t, err)
	if update {
		actual, err := json.MarshalIndent(result, "", "    ")
		require.NoError(t, err)
		require.NoError(t, ioutil.WriteFile(golden, []byte(actual), 0644))
	}
	expected, err := ioutil.ReadFile(golden)
	require.NoError(t, err)
	var expectedResults gotest.Results
	require.NoError(t, json.Unmarshal(expected, &expectedResults))
	require.ElementsMatch(t, expectedResults, result)
}

func TestGoTestLogs(t *testing.T) {
	gotests := path.Join("testdata", "testresults")
	files, err := ioutil.ReadDir(gotests)
	require.NoError(t, err)

	for _, f := range files {
		if !f.IsDir() && !strings.HasSuffix(f.Name(), ".golden") {
			t.Run(f.Name(), func(t *testing.T) {
				file := path.Join(gotests, f.Name())
				golden := file + ".golden"
				f, err := os.Open(file)
				require.NoError(t, err)
				defer f.Close()
				testConverterParse(t, f, golden)
			})
		}
	}
}

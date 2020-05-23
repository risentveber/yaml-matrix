package matrix

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/tj/assert"
)

type testCase struct {
	input  string
	output string
	err    string
}

func TestContainerCreation(t *testing.T) {
	cases := []testCase{
		{
			input:  "../examples/test-input-1.yaml",
			output: "../examples/test-output-1.yaml",
		},
		{
			input:  "../examples/nested-input-1.yaml",
			output: "../examples/nested-output-1.yaml",
		},
		{
			input:  "../examples/root-array-input.yaml",
			output: "../examples/root-array-output.yaml",
		},
		{
			input: "../examples/error-input.yaml",
			err:   ".task.matrix.matrix invalid type map[interface {}]interface {}",
		},
	}

	for _, c := range cases {
		filename, _ := filepath.Abs(c.input)
		inputBytes, err := ioutil.ReadFile(filename)
		assert.NoError(t, err)

		s := NewService("matrix")
		bytes, err := s.Convert(inputBytes)
		if c.err != "" {
			assert.EqualError(t, err, c.err)
			continue
		}

		filename, _ = filepath.Abs(c.output)
		outputBytes, err := ioutil.ReadFile(filename)
		assert.NoError(t, err)
		fmt.Println(string(bytes))
		assert.Equal(t, string(outputBytes), string(bytes), c.output)
	}
}

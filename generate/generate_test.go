package generate

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSampleOutput(t *testing.T) {
	buf := &bytes.Buffer{}
	err := GenerateTypeScript(buf, []string{"testfiles/sample.go"})
	assert.NoError(t, err)

	f, err := os.Open("testfiles/sample.ts")
	require.NoError(t, err)

	b, err := ioutil.ReadAll(f)
	require.NoError(t, err)

	assert.Equal(t, string(b), buf.String())
}

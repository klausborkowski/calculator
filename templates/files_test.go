package templates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTemplatesAreEmbedded(t *testing.T) {
	require.NotEmpty(t, IndexHTML, "IndexHTML template should not be empty")
}

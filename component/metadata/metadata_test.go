package metadata

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_inferMetadata(t *testing.T) {
	tests := []struct {
		name     string
		expected Metadata
	}{
		{
			name:     "discovery.dns",
			expected: Metadata{Exports: []Type{TypeTargets}},
		},
		{
			name: "discovery.relabel",
			expected: Metadata{
				Accepts: []Type{TypeTargets},
				Exports: []Type{TypeTargets},
			},
		},
		{
			name:     "loki.echo",
			expected: Metadata{Accepts: []Type{TypeLokiLogs}},
		},
		{
			name: "loki.source.file",
			expected: Metadata{
				Accepts: []Type{TypeTargets},
				Exports: []Type{TypeLokiLogs},
			},
		},
		{
			name: "loki.process",
			expected: Metadata{
				Accepts: []Type{TypeLokiLogs},
				Exports: []Type{TypeLokiLogs},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ForComponent(tt.name)
			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}

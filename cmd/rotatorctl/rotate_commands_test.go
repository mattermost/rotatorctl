package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotateCommand(t *testing.T) {
	cmd := newRotateCmd()

	t.Run("missing required --cluster", func(t *testing.T) {
		err := cmd.Execute()
		require.Error(t, err)
		require.Contains(t, err.Error(), "\"cluster\" not set")
	})
	t.Run("wrong type --max-scaling", func(t *testing.T) {
		cmd.SetArgs([]string{"--cluster", "my-rotated-cluster"})
		cmd.SetArgs([]string{"--max-scaling", "'0'"})
		err := cmd.Execute()
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid argument \"'0'\"")
	})
}

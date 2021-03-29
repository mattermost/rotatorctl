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
	t.Run("negative --max-scaling", func(t *testing.T) {
		cmd.SetArgs([]string{"--cluster", "my-rotated-cluster"})
		cmd.SetArgs([]string{"--max-scaling", "-1"})
		err := cmd.Execute()
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid argument \"-1\"")
	})
	t.Run("negative --max-drain-retries", func(t *testing.T) {
		cmd.SetArgs([]string{"--cluster", "my-rotated-cluster"})
		cmd.SetArgs([]string{"--max-drain-retries", "-1"})
		err := cmd.Execute()
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid argument \"-1\"")
	})
	t.Run("negative --evict-grace-period", func(t *testing.T) {
		cmd.SetArgs([]string{"--cluster", "my-rotated-cluster"})
		cmd.SetArgs([]string{"--evict-grace-period", "-1"})
		err := cmd.Execute()
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid argument \"-1\"")
	})
	t.Run("negative --wait-between-rotations", func(t *testing.T) {
		cmd.SetArgs([]string{"--cluster", "my-rotated-cluster"})
		cmd.SetArgs([]string{"--wait-between-rotations", "-1"})
		err := cmd.Execute()
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid argument \"-1\"")
	})
	t.Run("negative --wait-between-drains", func(t *testing.T) {
		cmd.SetArgs([]string{"--cluster", "my-rotated-cluster"})
		cmd.SetArgs([]string{"--wait-between-drains", "-1"})
		err := cmd.Execute()
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid argument \"-1\"")
	})
}

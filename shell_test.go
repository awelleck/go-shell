package main

import (
  "testing"
  "github.com/stretchr/testify/require"
)

func testEcho (t *testing.T) {
	echo_cmd := "echo Hello World"
	test_output, err := executor(echo_cmd)
	require.NoError(t, err)
	require.Equal(t, "Hello World", test_output)
}

package cflare_test

import (
	"context"
	"os"
	"testing"

	"github.com/stellaraf/zonebackup/internal/cflare"
	"github.com/stretchr/testify/require"
)

func Test_Collect(t *testing.T) {
	token := os.Getenv("CLOUDFLARE_API_TOKEN")
	require.NotEmpty(t, token, "missing cloudflare api token from env")
	t.Parallel()
	ctx := context.Background()
	dir := t.TempDir()
	err := cflare.Collect(ctx, token, dir)
	require.NoError(t, err)
}

package cmd

import (
	"github.com/stellaraf/zonebackup/internal/cflare"
	"github.com/urfave/cli/v2"
)

var cloudflareToken string
var outDir string

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "cloudflare-token",
		Usage:       "Cloudflare API Token",
		Destination: &cloudflareToken,
		Aliases:     []string{"ct"},
	},
	&cli.StringFlag{
		Name:        "out-dir",
		Usage:       "Output directory",
		Destination: &outDir,
		Aliases:     []string{"o", "out"},
	},
}

var CLI = &cli.App{
	Name:  "zonebackup",
	Usage: "Export DNS zone files from DNS providers",
	Flags: flags,
	Action: func(ctx *cli.Context) error {
		err := cflare.Collect(ctx.Context, cloudflareToken, outDir)
		if err != nil {
			return err
		}
		return nil
	},
}

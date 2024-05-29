package cflare

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/destel/rill"
	"github.com/joomcode/errorx"
)

func Collect(ctx context.Context, token, dir string) error {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(dir, os.ModePerm)
			if err != nil {
				err = errorx.Decorate(err, "failed to create %s", dir)
				return err
			}
		} else {
			err = errorx.Decorate(err, "failed to read %s", dir)
			return err
		}
	}
	api, err := cloudflare.NewWithAPIToken(token)
	if err != nil {
		err = errorx.Decorate(err, "failed to initialize connection to Cloudflare")
		return err
	}
	zones, err := api.ListZones(ctx)
	if err != nil {
		err = errorx.Decorate(err, "failed to retrieve zones")
		return err
	}
	now := time.Now()
	date := now.Format("20060102")
	batches := rill.FromSlice(zones, nil)
	err = rill.ForEach(batches, len(zones), func(zone cloudflare.Zone) error {
		export, err := api.ZoneExport(ctx, zone.ID)
		if err != nil {
			err = errorx.Decorate(err, "failed to export zone %s", zone.Name)
			return err
		}

		fileName := path.Join(dir, fmt.Sprintf("%s-cloudflare-%s.zone", date, zone.Name))
		file, err := os.Create(fileName)
		if err != nil {
			err = errorx.Decorate(err, "failed to create file at '%s'", fileName)
			return err
		}
		_, err = file.WriteString(export)
		if err != nil {
			err = errorx.Decorate(err, "failed to write to file '%s'", fileName)
			return err
		}
		log.Printf("exported %s to %s", zone.Name, fileName)
		return nil
	})
	return err
}

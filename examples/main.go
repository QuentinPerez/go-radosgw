package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/QuentinPerez/go-radosgw/pkg/api"
	"github.com/Sirupsen/logrus"
)

func printRawMode(out io.Writer, data interface{}) error {
	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s\n", js)
	return nil
}

func main() {
	api := radosAPI.New("http://192.168.42.40", os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

	now := time.Now().AddDate(0, 0, -1)
	usage, err := api.GetUsage(&radosAPI.UsageConfig{
		UID:         "qperez",
		ShowEntries: true,
		ShowSummary: true,
		Start:       &now,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	printRawMode(os.Stdout, usage)
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/QuentinPerez/go-radosgw/pkg/api"
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
	api, err := radosAPI.New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))
	if err != nil {
		log.Fatal(err)
	}

	// create a new user named JohnDoe
	user, err := api.CreateUser(radosAPI.UserConfig{
		UID:         "JohnDoe",
		DisplayName: "John Doe",
	})
	if err != nil {
		log.Fatal(err)
	}
	printRawMode(os.Stdout, user)

	// remove JohnDoe
	err = api.RemoveUser(radosAPI.UserConfig{
		UID: "JohnDoe",
	})
	if err != nil {
		log.Fatal(err)
	}
}

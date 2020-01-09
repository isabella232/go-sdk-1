package main

import (
	"fmt"
	"github.com/configcat/go-sdk"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	config := configcat.DefaultClientConfig()
	config.Logger = logger

	client := configcat.NewCustomClient("PKDVCLf-Hq-h-kCzMp-L7Q/psuH7BGHoUmdONrzzUOY7A", config)

	// create a user object to identify your user (optional)
	user := configcat.NewUser("##SOME-USER-IDENTIFICATION##")

	// get individual config values identified by a key and a user object
	value := client.GetValueForUser("keySampleText", "", user)

	fmt.Println("keySampleText: ", value)
}

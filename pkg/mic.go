package pkg

import (
	"fmt"
	"log"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/inputs"
)

func AssertDeviceID(obs *goobs.Client, audioSource string, deviceID string) error {
	if deviceID == "" {
		return fmt.Errorf("device flag is required")
	}

	res, err := obs.Inputs.GetInputSettings(&inputs.GetInputSettingsParams{
		InputName: audioSource,
	})

	if err != nil {
		return err
	}

	if id := res.InputSettings["device_id"]; id != deviceID {
		log.Fatalf("expected \"%s\", got \"%s\"", deviceID, id)
	}

	return nil
}

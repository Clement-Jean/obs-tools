package pkg

import (
	"fmt"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/events"
	"github.com/andreykaipov/goobs/api/requests/config"
)

func RecordWithName(obs *goobs.Client, name string, timestamp string) error {
	waitc := make(chan struct{})
	go obs.Listen(func(event any) {
		switch e := event.(type) {
		case *events.RecordStateChanged:
			if e.OutputState == "OBS_WEBSOCKET_OUTPUT_STARTED" {
				close(waitc)
			}
		}
	})

	oldFileFormat, err := obs.Config.GetProfileParameter(&config.GetProfileParameterParams{
		ParameterCategory: "Output",
		ParameterName:     "FilenameFormatting",
	})

	if err != nil {
		return err
	}

	if len(timestamp) == 0 && len(name) == 0 {
		return fmt.Errorf("if no timestamp is provided, name should be set")
	}

	var value string = name

	if len(timestamp) != 0 {
		if len(name) != 0 {
			value += " - "
		}

		value += timestamp
	}

	_, err = obs.Config.SetProfileParameter(&config.SetProfileParameterParams{
		ParameterCategory: "Output",
		ParameterName:     "FilenameFormatting",
		ParameterValue:    value,
	})

	if err != nil {
		return err
	}

	if _, err = obs.Record.StartRecord(); err != nil {
		return err
	}

	<-waitc
	_, err = obs.Config.SetProfileParameter(&config.SetProfileParameterParams{
		ParameterCategory: "Output",
		ParameterName:     "FilenameFormatting",
		ParameterValue:    oldFileFormat.ParameterValue,
	})

	return err
}

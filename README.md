# obs-tools

## Example

### Recording with file name

```
go run main.go --pass <PASSWORD> record --name <NAME>
```

### Recording with file name and timestamp "%MM-%DD %hh-%mm"

```
go run main.go --pass <PASSWORD> record --name <NAME> --timestamp "%MM-%DD %hh-%mm"
```

### Checking microphone is correct before recording

```
go run main.go --pass <PASSWORD> and -- mic --device <DEVICE_ID> --source <AUDIO_SOURCE_NAME> - record --name <NAME>
```

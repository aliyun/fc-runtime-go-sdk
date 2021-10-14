
## build and zip
```bash
GOOS=linux CGO_ENABLED=0 go build -o event-map event-map.go
zip -r event-map.zip event-map
```
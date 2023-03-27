NAME = gpx-converter

build:
	go build -o bin/${NAME} cmd/gpx-converter/main.go

clean:
	rm -f bin/${NAME}
	rm -rf bin/

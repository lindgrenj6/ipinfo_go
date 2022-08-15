all: build

build: 
	go build -o ipinfo

clean:
	go clean

install: build
	install ipinfo ~/bin/

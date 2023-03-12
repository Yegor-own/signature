build:
	mkdir build
	cp -r static build/
	cp .env build/
	go build -o build/bin/signature main.go

clear-build:
	rm -rf build

run:
	make clear-build
	make build
	./build/bin/signature
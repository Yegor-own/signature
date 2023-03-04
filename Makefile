build:
	#rm -rf build
	mkdir build
	mkdir build/bin
	#mkdir build/template
	go build -o build/bin/signature main.go
	cp -r src/template build/
#go build -o $GOPATH/bin/outputfile.exe source.go

run:

clear:
	rm -rf build
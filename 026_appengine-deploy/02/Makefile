
APP=bin/application

# env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${APP} main.go
#To read more on the two method to deploy your 
build: clean
	mkdir bin
	cp -rf css bin/css
	cp -rf img bin/img
	cp application.go bin/application.go
	cp -rf index.html bin/index.html
	cp go.mod bin/go.mod
	cd bin && zip -r Archive.zip .
	cd ..

clean:
	rm -rf bin/ || true
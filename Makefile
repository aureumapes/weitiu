install:
	cd resource && npm run compile
	go get -v
	go build -v
	cp -v -r ./resource ~/.weitiu
	sudo cp ./weitiu /usr/bin
	sudo chmod +xrw /usr/bin/weitiu
	weitiu
	echo "passwd"
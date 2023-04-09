build-and-replace:
	# delete old binary (-f to not fail if it is not existing yet)
	sudo rm -f /usr/bin/locvis
	# build
	go build
	# move
	sudo mv ./locvis /usr/bin

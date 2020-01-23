# teoh
text editor over http

## how to use

	teoh file-to-edit

	# see the file
	curl http://localhost

	# change the file
	curl --data 'please make me better' http://localhost

Point your browser to http://localhost/edit to edit over a web UI.

## how to install

	go get github.com/rganardi/teoh

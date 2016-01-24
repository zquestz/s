APPNAME = s

all:
	go build .

install: all
	sudo install -d /usr/local/bin
	sudo install -c ${APPNAME} /usr/local/bin/s

uninstall:
	sudo rm /usr/local/bin/s

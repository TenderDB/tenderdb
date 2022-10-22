

all: test mkdirs build template asset db start

test:
	sudo go test ./repositories/...

mkdirs:
	sudo mkdir -p /usr/local/app
	sudo mkdir -p /usr/local/app/bin
	sudo mkdir -p /usr/local/app/assets
	sudo mkdir -p /usr/local/app/templates
	sudo mkdir -p /usr/local/app/db

build:
	sudo OOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /usr/local/app/bin/tenderdb main.go
	sudo chmod 700 /usr/local/app/bin/tenderdb
	sudo chown www-data:www-data /usr/local/app/bin/.
	sudo chown www-data:www-data /usr/local/app/bin/*
template:	
	sudo cp -r templates/ /usr/local/app
	sudo chmod 400 `find /usr/local/app/templates/ -type f`
	sudo chown www-data:www-data /usr/local/app/templates/*
asset:
	sudo cp -r -u assets/ /usr/local/app	
	sudo chmod 400 `find /usr/local/app/assets/ -type f`
	sudo chown www-data:www-data /usr/local/app/assets/.
	sudo chown www-data:www-data /usr/local/app/assets/*
db:
	sudo gunzip db/*
	sudo cp -r -u db/ /usr/local/app	
	sudo chmod 600 /usr/local/app/db/* 
	sudo chown www-data:www-data /usr/local/app/db/.
	sudo chown  www-data:www-data /usr/local/app/db/*
	sudo gzip db/*
permit:
	sudo chmod 700 /usr/local/app/bin/tenderdb
	sudo chmod 400 `find /usr/local/app/bin/templates/ -type f`
	sudo chmod 400 `find /usr/local/app/assets/ -type f`
	sudo chmod 600 `find /usr/local/app/db/ -type f`
	sudo chown www-data:www-data /usr/local/app/bin/.
	sudo chown www-data:www-data /usr/local/app/bin/*
	sudo chown www-data:www-data /usr/local/app/assets/.
	sudo chown www-data:www-data /usr/local/app/assets/*
	sudo chown www-data:www-data /usr/local/app/templates/.
	sudo chown www-data:www-data /usr/local/app/templates/*
	sudo chown www-data:www-data /usr/local/app/db/.
	sudo chown www-data:www-data /usr/local/app/db/*
start:	
	sudo cp daemon.conf /etc/systemd/system/app.service 
	sudo systemctl daemon-reload
	sudo systemctl start app
stop:
	sudo systemctl stop app



.PHONY: test mkdirs build template asset db permit all

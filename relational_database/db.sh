#!/bin/bash
podman build -t example-go/mysql:latest .
podman run --rm -d -p 3306:3306 --name=mysql mysql 

sleep 7
podman exec -it mysql mysql -u root -proot -e "USE recordings; SELECT * FROM album;"


go build data-access/
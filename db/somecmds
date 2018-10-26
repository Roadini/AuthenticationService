#!/bin/bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null && pwd )"

docker run --name db -p 3306:3306 \
    -v $DIR/startupscripts:/docker-entrypoint-initdb.d \
    -v $DIR/log:/var/logs \
    -e MYSQL_ROOT_PASSWORD=pass \
    -e MYSQL_DATABASE=db \
    -d mysql:latest



    # -v ~/Development/web/myproject/docker/mysql:/var/lib/mysql \

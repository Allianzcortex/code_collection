#!/usr/bin/env bash

function start() {
    docker-compose up --build
}

function stop() {
    docker-compose down
}

function test() {
    docker exec -it wave_app sh -c "pipenv run pytest -s"
}

function send() {
    echo "The data of $1 is going to be stored in database"

    curl -i -X POST -F file=@$1 http://localhost:5000/api/upload
}

function summary() {
    curl -X GET http://127.0.0.1:5000/api/summary
}

function usage() {
    name=`basename $0`
    
    echo "USAGE: $name COMMAND [ARGS...]"
    echo ""
    echo "Commands:"
    echo "   start   Start the backend "
    echo "   stop    Stop the backend "
    echo "   test    Run all tests "
    echo "   summary Get Summary of all records"
    echo "   send    add filename(path) "
}

case $1 in
"start"    ) start          ;;
"stop"     ) stop           ;;
"test"     ) test           ;;
"summary"  ) summary        ;;
"send"     ) send $2        ;;
*          ) usage          ;;	
esac

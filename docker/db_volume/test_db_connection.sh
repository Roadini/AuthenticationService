#!/bin/sh
until nc -z db 3306; do sleep 1; echo "Waiting for DB to come up..."; done

echo "executing db"
app

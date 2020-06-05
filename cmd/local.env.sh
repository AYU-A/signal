#!/bin/bash
export GOOGLE_APPLICATION_CREDENTIALS=../../hotel-alarm-2240fe8d3432.json

# Local Postgres
export SIGNAL_PG_PROTO=tcp
export SIGNAL_PG_ADDR=localhost:5432
export SIGNAL_PG_DB=signals
export SIGNAL_PG_USER=postgres
export SIGNAL_PG_PASSWORD=postgres

export SIGNAL_GC_PROJECT=hotel-signal
export SIGNAL_GC_PROJECT_LOCATION=europe-west3
export SIGNAL_CALL_ROOM_ENDPOINT=https://$SIGNAL_GC_PROJECT_LOCATION-$SIGNAL_GC_PROJECT.cloudfunctions.net/CallRoom
export SIGNAL_SCHEDULER_LOCATION=projects/$SIGNAL_GC_PROJECT/locations/$SIGNAL_GC_PROJECT_LOCATION
export SIGNAL_SCHEDULER_TIMEZONE=Europe/Moscow
export SIGNAL_SCHEDULER_MAX_RETRY_COUNT=2
export SIGNAL_SCHEDULER_RETRY_PERIOD=1m
export SIGNAL_CALL_ENDPOINT=https://www.google.com
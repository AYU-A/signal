#!/bin/bash

# Google Cloud Postgres
export SIGNAL_PG_PROTO=unix
export SIGNAL_PG_ADDR=/cloudsql/hotel-signal:europe-west3:signals/.s.PGSQL.5432
export SIGNAL_PG_DB=signals
export SIGNAL_PG_USER=signals
export SIGNAL_PG_PASSWORD=OqGlOvyNGmNmraas

export SIGNAL_GC_PROJECT=hotel-signals
export SIGNAL_GC_PROJECT_LOCATION=europe-west3
export SIGNAL_CALL_ROOM_ENDPOINT=https://$SIGNAL_GC_PROJECT_LOCATION-$SIGNAL_GC_PROJECT.cloudfunctions.net/CallRoom
export SIGNAL_SCHEDULER_LOCATION=projects/$SIGNAL_GC_PROJECT/locations/$SIGNAL_GC_PROJECT_LOCATION
export SIGNAL_SCHEDULER_TIMEZONE=Europe/Moscow
export SIGNAL_SCHEDULER_MAX_RETRY_COUNT=2
export SIGNAL_SCHEDULER_RETRY_PERIOD=1m
export SIGNAL_CALL_ENDPOINT=https://www.google.com

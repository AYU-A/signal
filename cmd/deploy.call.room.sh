#!/bin/bash
source ./deploy.env.sh
cd ..
gcloud functions deploy CallRoom \
--runtime go113 \
--allow-unauthenticated \
--trigger-http \
--project $SIGNAL_GC_PROJECT \
--set-env-vars SIGNAL_PG_PROTO=$SIGNAL_PG_PROTO,SIGNAL_PG_ADDR=$SIGNAL_PG_ADDR,SIGNAL_PG_DB=$SIGNAL_PG_DB,SIGNAL_PG_USER=$SIGNAL_PG_USER,SIGNAL_PG_PASSWORD=$SIGNAL_PG_PASSWORD,SIGNAL_GC_PROJECT=$SIGNAL_GC_PROJECT,SIGNAL_GC_PROJECT_LOCATION=$SIGNAL_GC_PROJECT_LOCATION,SIGNAL_CALL_ROOM_ENDPOINT=$SIGNAL_CALL_ROOM_ENDPOINT,SIGNAL_SCHEDULER_LOCATION=$SIGNAL_SCHEDULER_LOCATION,SIGNAL_SCHEDULER_TIMEZONE=$SIGNAL_SCHEDULER_TIMEZONE,SIGNAL_CALL_ENDPOINT=$SIGNAL_CALL_ENDPOINT,SIGNAL_SCHEDULER_MAX_RETRY_COUNT=$SIGNAL_SCHEDULER_MAX_RETRY_COUNT,SIGNAL_SCHEDULER_RETRY_PERIOD=$SIGNAL_SCHEDULER_RETRY_PERIOD \
--region $SIGNAL_GC_PROJECT_LOCATION \
--memory 128MB \
--max-instances 10
cd cmd
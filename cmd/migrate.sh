#!/bin/bash
source ./ayu.env.sh
go run ../db/migrations/*.go $1
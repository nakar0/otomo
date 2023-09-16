#!/bin/sh

printf '\033[32m%s\033[m\n' 'Starting Firestore Emulator...'

gcloud beta emulators firestore start --host-port=localhost:8200 --quiet &>/dev/null &
curl http://localhost:8200 --silent --retry 30 --retry-delay 1 --retry-connrefused &>/dev/null

printf '\033[32m%s\033[m\n' 'Started Firestore Emulator'

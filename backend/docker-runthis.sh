#!/bin/sh

for bin in /nyg-backend/bin/*; do
  echo "Starting $bin"
  $bin &
done
wait
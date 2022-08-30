#!/bin/sh


set -xe

for path in $(ls -d */) 
do
    go test ./$path
done
#!/bin/sh

padded=$(printf "%02d" $1)
mkdir $padded -p
touch $padded/inp.txt
cp template.go $padded/$padded.go -n 
cp template_test.go $padded/${padded}_test.go -n 

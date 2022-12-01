#!/bin/bash

padded=$(printf "%02d" $1)
mkdir $padded -p
touch $padded/inp.txt
cp template.go $padded/$padded.go -n 

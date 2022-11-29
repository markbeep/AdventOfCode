#!/bin/bash

mkdir $1 -p
touch $1/inp.txt
cp template.go $1/$1.go -n 

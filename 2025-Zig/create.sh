#!/bin/bash

if [ -z "$1" ]; then
  echo "Error: Please provide a day number."
  echo "Usage: ./create_day.sh <day_number>"
  exit 1
fi

CLEAN_INPUT=$(echo "$1" | sed 's/^0*//')
NEW_DAY=$(printf "%02d" "$CLEAN_INPUT")

PREV_DAY_NUM=$((CLEAN_INPUT - 1))
PREV_DAY=$(printf "%02d" "$PREV_DAY_NUM")

echo "Setting up Day $NEW_DAY (replacing Day $PREV_DAY in runner)..."

cp src/template.zig "src/days/$NEW_DAY.zig"

sed -i "s/xx.txt/$NEW_DAY.txt/g" "src/days/$NEW_DAY.zig"

sed -i "s/$PREV_DAY/$NEW_DAY/g" src/main.zig

touch "days/$NEW_DAY.txt"

echo "Success! Created src/days/$NEW_DAY.zig and updated src/main.zig."

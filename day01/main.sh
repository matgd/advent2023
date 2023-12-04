#!/bin/bash

echo "Day 1 (1/2):"
sed 's/[^0-9]//g' input.txt | awk -F '' '{ s += $1$NF } END { print s }'

echo "Day 1 (2/2):"
sed -e s'/one/one1one/g' \
    -e s'/two/two2two/g' \
    -e s'/three/three3three/g' \
    -e s'/four/four4four/g' \
    -e s'/five/five5five/g' \
    -e s'/six/six6six/g' \
    -e s'/seven/seven7seven/g' \
    -e s'/eight/eight8eight/g' \
    -e s'/nine/nine9nine/g' \
    -e 's/[^0-9]//g' input.txt | awk -F '' '{ s += $1$NF } END { print s }'

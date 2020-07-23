#!/bin/sh
# example 1: correct usage
echo "I'm an input file." > input
./fakessim2qff -t '201809010005' -o output -d input
if [ $? -ne 0 ] || [ ! -f "./output" ] || [ "$(cat output)" != "I'm an output file." ]
then
	echo "Test 1: not working properly."
fi

# example 2: incorrect usage (misses output flags)
echo "I'm an input file." > input
./fakessim2qff -t '201809010005' -d input
if [ $? -eq 0 ]
then
	echo "Test 2: not working properly."
fi

# example 3: incorrect usage (wrong input file content)
echo "I'm an incorrect input file." > input
./fakessim2qff -t '201809010005' -o output -d input
if [ $? -eq 0 ]
then
	echo "Test 3: not working properly."
fi

# example 4: incorrect usage (wrong timestamp)
echo "I'm an input file." > input
./fakessim2qff -t '0' -o output -d input
if [ $? -eq 0 ]
then
	echo "Test 4: not working properly."
fi

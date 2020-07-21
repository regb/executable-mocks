#!/bin/sh
# correct usage example
echo "I'm an input file." > input
./fakessim2qff -t '201809010005' -o output -d input
if [ $? -eq 1 ]
then
	echo "Test 1: not working properly."
fi
# incorrect usage example
./fakessim2qff -t '20180901005' -d input
if [ $? -eq 0 ]
then
	echo "Test 2: not working properly."
fi

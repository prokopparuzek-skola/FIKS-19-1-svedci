#!/bin/bash

sed -e "/^$/d" input.txt > in.txt
./svedci <in.txt >output.txt

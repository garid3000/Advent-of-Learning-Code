#!/bin/bash

for i in {1..27625..20}
do
	tput cup 0 0
	cat "grid/$i"
done
#tput cup 0 0

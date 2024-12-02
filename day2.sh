#!/bin/bash

isSafe () {
	arr2=("$@")	
	curr=${arr2[0]}
	sec=${arr2[1]}
	if (( curr > sec )) ; then
		inc=false
	elif (( curr< sec )) ; then
		inc=true
	else
		echo "Not increasing or decreasing"
		return 0

	fi 

	for i in "${arr2[@]:1}"; do
	echo "previous item is $curr, current one is $i"

		if $inc; then
			if (( curr >= i )) ; then  echo "Not increasing"; return 0; fi
			diff=$((i-curr))
			echo "difference is $diff"
			if (( diff < 1 ))  || (( diff > 3 )) ; then return 0; fi
			curr=$i

		else
			if (( curr <= i )) ; then echo "Not decreasing";return 0; fi
			diff=$((curr-i))
			echo "difference is $diff"

			if (( diff < 1 ))  || (( diff > 3 )) ; then return 0; fi
			curr=$i

		fi;

	done
	return 1;
}

part1 () {
count=0
while IFS= read -r line; do
read -r -a arr <<< "$line"
isSafe "${arr[@]}"
((count+=$?))
done < day2.txt

echo $count
}

part1


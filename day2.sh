#!/bin/bash

main() {
	arr2=("$@")
	curr=${arr2[0]}
	sec=${arr2[1]}
	if ((curr > sec)); then
		inc=false
	elif ((curr < sec)); then
		inc=true
	else
		return 0

	fi

	for i in "${arr2[@]:1}"; do

		if $inc; then
			if ((curr >= i)); then
				return 0
			fi
			diff=$((i - curr))
			if ((diff < 1)) || ((diff > 3)); then return 0; fi
			curr=$i

		else
			if ((curr <= i)); then
				return 0
			fi
			diff=$((curr - i))

			if ((diff < 1)) || ((diff > 3)); then return 0; fi
			curr=$i

		fi

	done
	return 1
}

run1() {
	count=0
	while IFS= read -r line; do
		read -r -a arr <<<"$line"
		main "${arr[@]}"
		total=$?
		((count += total))
	done <day2.txt

	echo $count
}

run2() {
	count=0
	while IFS= read -r line; do
		read -r -a arr <<<"$line"
		length="${#arr[@]}"
		echo "Original array: ${arr[@]}"
		for ((i = 0; i < length; i++)); do
			arr2=("${arr[@]}")
			unset 'arr2[i]'
			arr2=("${arr2[@]}")
			main "${arr2[@]}"
			total=$?
			echo "Total: $total"
			if ((total > 0)); then
				((count += total))
				break
			fi
		done
	done <day2.txt
	echo "Count: $count"
}

run2

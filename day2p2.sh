#!/bin/bash

run() {
	count=0
	while IFS= read -r line; do
		read -r -a arr <<<"$line"
		$1 "${arr[@]}"
		((count += $?))
	done <day2.txt

	echo $count
}

d2p2() {
	arr2=("$@")
	p1=0
	p2=1
	length=${#arr2[@]}
	while ((p2 < length)); do
		echo $p2

	done
}

#!/bin/bash

part1() {
	arr2=("$@")
	curr=${arr2[0]}
	sec=${arr2[1]}
	if ((curr > sec)); then
		inc=false
	elif ((curr < sec)); then
		inc=true
	else
		echo "Not increasing or decreasing"
		return 0

	fi

	for i in "${arr2[@]:1}"; do
		echo "previous item is $curr, current one is $i"

		if $inc; then
			if ((curr >= i)); then
				echo "Not increasing"
				return 0
			fi
			diff=$((i - curr))
			echo "difference is $diff"
			if ((diff < 1)) || ((diff > 3)); then return 0; fi
			curr=$i

		else
			if ((curr <= i)); then
				echo "Not decreasing"
				return 0
			fi
			diff=$((curr - i))
			echo "difference is $diff"

			if ((diff < 1)) || ((diff > 3)); then return 0; fi
			curr=$i

		fi

	done
	return 1
}

run() {
	count=0
	while IFS= read -r line; do
		read -r -a arr <<<"$line"
		$1 "${arr[@]}"
		temp1=$?
		echo "Trying without first input"
		$1 "${arr[@]:1}"
		temp2=$?
		reversed_arr=($(printf '%s\n' "${arr[@]}" | gtac))
		temp3=$?
		if ((temp1)) || ((temp2)) || ((temp3)); then
			total=1
		else
			total=0
		fi
		((count = count + total))
	done <day2.txt

	echo $count
}

part2() {
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
	freeBie=true

	for i in "${arr2[@]:1}"; do

		if $inc; then
			if ((curr >= i)); then

				if $freeBie; then
					echo "Using freebie as not strict increasing"
					freeBie=false
					continue
				else

					echo "Not increasing, used freeBie so unsafe"
					return 0
				fi

			fi
			diff=$((i - curr))
			if ((diff < 1)) || ((diff > 3)); then
				if $freeBie; then
					echo "Using freebie as difference is too big: $diff"
					freeBie=false
					continue
				else
					echo "Difference is too big, unsafe"
					return 0
				fi
			fi
			curr=$i

		else
			if ((curr <= i)); then
				if $freeBie; then
					echo "Using freebie as not decreasing"
					freeBie=false
					continue
				else
					echo "Not decreasing, used freebie so unsafe"
					return 0
				fi
			fi

			diff=$((curr - i))
			if ((diff < 1)) || ((diff > 3)); then

				if $freeBie; then
					echo "Using freebie as difference is too big: $diff"
					freeBie=false
					continue
				else
					echo "Difference is too big, unsafe"
					return 0
				fi
			fi
			curr=$i
		fi

	done
	echo "Successful"
	return 1

}

run part2

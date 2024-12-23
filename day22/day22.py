import json
from collections import defaultdict


def mix(num1: int, num2:int) -> int:
	return num1^num2

def prune(num: int) -> int:
	return num % 16777216

def prune2(num: int) -> int:
	return num % 1024

def simulate(num: int) -> int:
	for i in range(2000):
		num2 = num << 6
		num = mix(num2,num)
		num = prune(num)

		num2 = num >> 5
		num = mix(num2, num)
		num = prune(num)

		num2 = num << 11
		num = mix(num2, num)
		num = prune(num)
	return num

def sim2(num: int) -> list[int]:
	latest_digits = [num % 10]
	for i in range(2000):
		num2 = num << 6
		num = mix(num2,num)
		num = prune(num)

		num2 = num >> 5
		num = mix(num2, num)
		num = prune(num)

		num2 = num << 11
		num = mix(num2, num)
		num = prune(num)

		last_digit = num % 10
		latest_digits.append(last_digit)


	return latest_digits
# count = 0
# with open('day22/input.txt') as f:
# 	for line in f:
# 		line = int(line.strip())
# 		ans = (simulate(line))
# 		count+=ans

def rolling_changes(arr: list[int]) -> dict[tuple[int,int,int,int], int]:
	sub_changes = []
	for i in range(len(arr)-1):
		diff = arr[i+1]-arr[i]
		sub_changes.append(diff)
	mapping = defaultdict(int)
	for i in range(3, len(sub_changes)):
		item = arr[i+1]
		changes = tuple(sub_changes[i-3: i+1])
		if changes not in mapping.keys():
			mapping[changes] = item
	return mapping

overall_mapping = defaultdict(list)
with open('day22/input.txt') as f:
	for line in f:
		line = int(line.strip())
		ans = sim2(line)
		mapping = rolling_changes(ans)
		for key, value in mapping.items():
			overall_mapping[key].append(value)

max_value = 0
pattern = []

for key, value in overall_mapping.items():
	total = sum(value)
	if total > max_value:
		pattern = key
		max_value = total

print(pattern)
print(len(overall_mapping[pattern]))
print(max_value)


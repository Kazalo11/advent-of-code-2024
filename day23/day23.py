import json
from collections import defaultdict
from functools import cache

connections = defaultdict(list[str])

t_values = set()
with open('day23/input.txt') as f:
	for line in f:
		line = line.strip()
		c1,c2 = line.split("-")
		if c1[0] == 't':
			t_values.add(c1)
		if c2[0] == 't':
			t_values.add(c2)
		connections[c1].append(c2)
		connections[c2].append(c1)

with open("day23/debug.txt", "w") as f:
	f.write(json.dumps(connections))


count = 0

def check_three_set(t: str) -> int:
	total = 0
	print(f'Checking for value {t}')
	conns = connections[t]
	for i in range(len(conns)-1):
		start = conns[i]
		for j in range(i+1, len(conns)):
			end = conns[j]
			if end in connections[start]:
				if start in t_values and end in t_values:
					total+=1/3
				elif start not in t_values and end not in t_values:
					total+=1
				else:
					total+=0.5
	return total



for t in t_values:
	val = check_three_set(t)
	count+=val

print(count)

# max_length = 0
# total_connecting = []
# magic_key=""

# @cache
# def is_connected(c1: str, c2: str) -> bool:
# 	return c2 in connections[c1] 
		


# def longest_connections(value: list[str]) -> int:
# 	count = 0
# 	for i in range(len(value)-1):
# 		for j in range(i+1, len(value)):
# 			if is_connected(value[i], value[j]):
# 				count+=1
# 	return count


# for key, value in connections.items():
# 	print(len(value))
# 	length = longest_connections(value)
# 	if max_length < length:
# 		max_length = length
# 		total_connecting = value + [key]
# 		magic_key=key

# total_connecting.sort()
# print(magic_key)
# print(','.join(total_connecting))


def longest_connections(path: list[str], curr: str) -> list[str]:
	neighbours = connections[curr]
	for item in path:
		if item not in neighbours:
			return path
	
	path.append(curr)
	max_length = 0
	max_val = []
	for neighbour in neighbours:
		val = longest_connections(path, neighbour)
		if max_length < len(val):
			max_length= len(val)
			max_val = val
	
	return max_val

ans_array = []
ans_key = ""
for key in connections.keys():
	long = longest_connections([], key)
	if len(long) > len(ans_array):
		ans_array = long
		ans_key = key

print(ans_key)

ans_array.sort()
print(",".join(ans_array))


	


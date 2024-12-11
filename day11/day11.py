with open('day11/day11.txt') as f:
	for line in f:
		arr = [int(x) for x in line.split(" ")]

start = {}
for val in arr:
	start[val] = 1

for i in range(25):
	for key, val in start.items():
		if key == 0:


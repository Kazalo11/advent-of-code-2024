from collections import defaultdict

index = 0
keys: list[list[int]] = []
locks: list[list[int]] = []
isKey = True
val = [0]*5
with open("day25/input.txt") as f:
	for line in f:
		line = line.strip()
		arr = list(line)
		match index % 8:
			case 0:
				if arr[0] == "#":
					isKey = False
				else:
					isKey = True
				val = [0]*len(arr)
			case 1 | 2 | 3 | 4 | 5:
				for i in range(len(arr)):
					if arr[i] == "#":
						val[i]+=1
			case 7:
				if isKey:
					keys.append(val)
				else:
					locks.append(val)

		index+=1
	if isKey:
		keys.append(val)
	else:
		locks.append(val)

count = 0		
for lock in locks:
	for key in keys:
		isOverlap = False
		for x,y in zip(lock,key):
			if x+y >5:
				isOverlap = True
				break
		if not isOverlap:
			count+=1

print(count)
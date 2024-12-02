def canBeSafe(line):
	if isSafe(line):
		return True
	for i in range(len(line)):
		line2 = line.copy()
		line2.pop(i)
		if isSafe(line2):
			return True
	return False
		

def isSafe(line):
	line2 = sorted(line)
	line3 = sorted(line,reverse=True)
	if line != line2 and line !=line3:
		print(line)
		print(line2)
		print(line3)
		print("Not all increasing or decreasing")
		return False
	for i in range(0, len(line)-1):
		curr = line[i]
		nexti = line[i+1]
		diff = abs(curr-nexti)
		if diff < 1 or diff > 3:
			return False	
	return True

total = 0

with open('day2.txt') as f:
	for line in f:
		line = line.strip().split(" ")
		if canBeSafe([int(i) for i in line]):
			total+=1

print(total)

import re
import statistics
from collections import defaultdict

width = 101
height = 103

# width = 11
# height = 7


iterations = 100

digits = r'-?\d+,-?\d+'

robots = defaultdict(list)

def extractNumbers(cood):
	first, second = cood.split(",")
	return (int(first), int(second))

def iteration(robots):
	ans = defaultdict(list)
	for pos, velArray in robots.items():
		x,y = pos
		for vel in velArray:
			vel_x, vel_y = vel
			new_x = (x + vel_x) % width
			new_y = (y + vel_y) % height
			ans[(new_x,new_y)].append(vel)
	return ans
			
def finalPosition(robots):
	ans = defaultdict(int)
	for pos, velArray in robots.items():
		x,y = pos
		for vel in velArray:
			vel_x, vel_y = vel
			new_x = (x + iterations*vel_x) % width
			new_y = (y + iterations*vel_y) % height
			ans[(new_x,new_y)]+=1
	return ans


with open('day14/day14.txt') as f:
	for line in f:
		line = line.strip()
		numbers = re.findall(digits, line)
		
		pos = extractNumbers(numbers[0])
		vel = extractNumbers(numbers[1])
		robots[pos].append(vel)
		
total = {}

def variance(robots):
	x_values = []
	y_values = []

	for pos, val in robots.items():
		x, y = pos
		for i in range(len(val)):
			x_values.append(x)
			y_values.append(y)
	var_x = statistics.variance(x_values)
	var_y = statistics.variance(y_values)
	return var_y


min_var = 1e9
min_i = 0
for i in range(10000):
	robots = iteration(robots)
	quadrants = defaultdict(list)

	# for pos, val in robots.items():
	# 	x, y = pos
	# 	if x < (width//2) and y < (height//2):
	# 		quadrants[0].append((x,y))
	# 	elif x > (width//2) and y <  (height//2):
	# 		quadrants[1].append((x,y))
	# 	elif x < (width//2) and y > (height//2):
	# 		quadrants[2].append((x,y))
	# 	elif x > (width//2) and y > (height//2):
	# 		quadrants[3].append((x,y))

	# count = 1
	# for key, val in quadrants.items():
	# 	count*=len(val)

	# total[i] = count

	var = variance(robots)
	if var < min_var:
		min_var = var
		min_i = i

print(min_i)

	




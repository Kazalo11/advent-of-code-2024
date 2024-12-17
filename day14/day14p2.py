import re
from collections import defaultdict
from itertools import groupby
from typing import List, Tuple

width = 101
height = 103

# width = 11
# height = 7


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

def checkIsTree(coords: List[Tuple[int,int]]):
	for _, group in groupby(coords, key=lambda x: x[0]):
		if len(list(group)) > 10:
			return True
	return False
	

with open('day14/day14.txt') as f:
	for line in f:
		line = line.strip()
		numbers = re.findall(digits, line)
		
		pos = extractNumbers(numbers[0])
		vel = extractNumbers(numbers[1])
		robots[pos].append(vel)
		

for i in range(10000):
	robots = iteration(robots)
	coords = list(robots.keys())
	if checkIsTree(coords):
		print(i)
		break





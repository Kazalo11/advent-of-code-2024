from itertools import groupby

matrix = []

row = 0

visited = set()

with open('day12/day12.txt') as f:
	for line in f:
		col = 0
		arr = list(line.strip())
		matrix.append(arr)

shapes = []

def findShape(row, col, item):
	if row < 0 or row >= len(matrix) or col < 0 or col >= len(matrix[0]):
		return []
	
	if matrix[row][col] != item:
		return []
	
	if (row,col) in visited:
		return []

	visited.add((row,col))

	return [(row,col)] + findShape(row+1, col, item) + findShape(row,col+1, item) + findShape(row, col-1, item) + findShape(row-1, col, item)


for rowIndex, row in enumerate(matrix):
	for colIndex, item in enumerate(row):
		if (rowIndex,colIndex) not in visited:
			shape = findShape(rowIndex, colIndex, item )
			if shape is not None:
				shapes.append(shape)


def calculatePerimeter1(shape):
	total = 0
	for item in shape:
		count = 4
		if (item[0]-1, item[1]) in shape:
			count-=1
		if (item[0]+1, item[1]) in shape:
			count-=1
		if (item[0], item[1]-1) in shape:
			count-=1
		if (item[0], item[1]+1) in shape:
			count-=1
		total+=count
	return total


total = 0

for shape in shapes:
	area = len(shape)
	perimeter = calculatePerimeter1(shape)
	total = total + (area*perimeter)

print(f'Answer for part 1: {total}')

def calculatePerimeter2(shape):
	total=0
	counter = 0
	counter2=0
	print(shape)
	for key, group in groupby(shape, lambda x: x[0]):
		group_list = list(group)
		print(group_list)
		for item in group_list:
			print("Counter")
			print(counter)
			if (item[0]+1, item[1]) in shape:
				if counter > 0:
					total+=1
					counter=0
			else:
				counter+=1
			print("Total")
			print(total)
		if counter > 0:
			total+=1

	for key, group in groupby(shape, lambda x: x[0]):	
		for item in group:
			if (item[0]-1, item[1]) in shape:
				if counter2 > 0:
					total+=1
					counter2=0
			else:
				counter2+=1
		if counter2 > 0:
			total+=1
		
	
	return total



total2 = 0
for shape in shapes:
	area = len(shape)
	print("Calculating perimeter")
	perimeter = calculatePerimeter2(shape)
	print(perimeter)
	total2+=(area*perimeter)

print(f'Answer for part2: {total2}')
arr1, arr2 = [],[]

with open('day1.txt') as f:
	for line in f:
		arr = line.strip().split(' ')
		arr1.append(int(arr[0]))
		arr2.append(int(arr[-1]))
arr1.sort()

arr2.sort()
total=0

for item1, item2 in zip(arr1,arr2):
	total+= abs(item1-item2)

print(total)

total2 = 0

for item in arr1:
	total2+= item*arr2.count(item)
print(total2)



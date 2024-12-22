
def mix(num1: int, num2:int) -> int:
	return num1^num2

def prune(num: int) -> int:
	return num % 16777216

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

count = 0

with open('day22/input.txt') as f:
	for line in f:
		line = int(line.strip())
		ans = (simulate(line))
		count+=ans

print(count)
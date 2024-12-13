import re

from sympy import Symbol, solve

num = r'\d+'
inputArr = []
item = []
with open("day13/day13.txt") as f:
	for line in f:
		line = line.strip()
		if line == "":
			inputArr.append(item)
			item = []
		else:
			x,y = re.findall(num,line)
			item.append((int(x),int(y)))
	inputArr.append(item)

a = Symbol('a')
b = Symbol('b')
sol = []

total = 0

for arr in inputArr:
	eq1 = arr[0][0]*a + arr[1][0]*b - (10000000000000+arr[2][0])
	eq2 = arr[0][1]*a + arr[1][1]*b - (10000000000000+arr[2][1])
	ans = solve([eq1,eq2], [a,b], dict=True)
	ans = sorted(ans, key=lambda i: arr[0][0]*i[a] + arr[0][1]*i[b])
	min_sol = ans[0]
	a_ans = min_sol[a]
	b_ans = min_sol[b]
	if int(a_ans) == a_ans and int(b_ans) == b_ans:
		total = total + 3* min_sol[a] + min_sol[b]
print(total)
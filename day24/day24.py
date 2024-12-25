from collections import defaultdict

wires: dict[str,int] = {}
calculations: dict[tuple[str,str], list[tuple[str,str]]] = defaultdict(list)

z_values = ""

with open('day24/input1.txt') as f:
	for line in f:
		line = line.strip()
		arr = line.split(": ")
		wire, value = arr[0], int(arr[1])
		wires[wire] = value

for i in range(44,-1,-1):
	val = str(i).zfill(2)
	x_value = f'x{val}'
	y_value = f'y{val}'
	z_value = wires[x_value] & wires[y_value]
	z_values+=str(z_value)

print(int(z_values, 2))



def calculate(input1: str, input2: str, operator: str, output: str) -> None:
	ans =  0
	match operator:
		case "AND":
			ans = wires[input1] & wires[input2]
		case "XOR":
			ans = wires[input1] ^ wires[input2]
		case "OR":
			ans = wires[input1] | wires[input2]
	wires[output] = ans


seen: set[tuple[str,str]] = set()

with open('day24/input2.txt') as f:
	for line in f:
		line = line.strip()
		arr = line.split(" ")
		input1 = arr[0]
		operator = arr[1]
		input2 = arr[2]
		output = arr[-1]
		calculations[(input1,input2)].append((operator, output))
while True:
	for key, values in calculations.items():
		if key in seen:
			continue
		input1, input2 = key[0], key[1]
		if input1 in wires.keys() and input2 in wires.keys():
			for value in values:
				calculate(input1, input2, operator=value[0], output=value[1])
				seen.add(key)
	if len(seen) == len(calculations.keys()):
		break

filtered_wires = {k:v for (k,v) in wires.items() if k[0] == 'z'}


sorted_wires = dict(sorted(filtered_wires.items()))
print(sorted_wires)

total = 0
for index, val in enumerate(sorted_wires.values()):
	total = total + (2**index)*val

print(total)

print(bin(total))

print(f'0b0{z_values}')	


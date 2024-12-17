register = {
	'A': 64012472,
	'B': 0,
	'C': 0,
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3
}

mapping = {
	0: '0',
	1: '1',
	2: '2',
	3: '3',
	4: 'A',
	5: 'B',
	6: 'C'
}

program = [0,1,5,4,3,0]

ans = []

for j in range(10000,20000):
	register['A'] = j
	i = 0

	while i < len(program)-1:

		opcode = program[i]
		operand = program[i+1]
		if operand != 7:
			combo = register[mapping[operand]]
		match opcode:
			case 0:
				register['A'] = register['A'] >> combo
				i+=2
			case 1:
				register['B'] = register['B']^operand
				i+=2
			case 2:
				register['B'] = combo % 8
				i+=2
			case 3:
				if register['A'] == 0:
					i+=2
				else:
					i=operand

			case 4:
				register['B'] = register['B'] ^ register['C']
				i+=2
			case 5:
				ans.append(str(combo%8))
				i+=2
			case 6:
				register['B'] = register['A'] >> combo
				i+=2

			case 7:
				register['C'] = register['A'] >> combo
				i+=2
	if register['A'] == j:
		print(j)
		break


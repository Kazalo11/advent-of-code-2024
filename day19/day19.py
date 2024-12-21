import re

i = 0
count = 0
words = ""
with open('day19/day19.txt') as f:
	for line in f:
		line = line.strip()
		if i == 0:
			line = line.replace(", ","|")
			words = f"(${line})+$"
		else:
			if re.fullmatch(words, line):
				count+=1
			
		i+=1

print(count)
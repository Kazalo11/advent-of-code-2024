from collections import deque
from functools import lru_cache
from typing import Any, Dict, List, Tuple

numeric_keypad: Dict[str, list[Tuple[str, str]]] = {
    '1': [('4', '^'), ('2', '>')],
    '2': [('5', '^'), ('1', '<'), ('3', '>')],
    '3': [('6', '^'), ('2', '<'), ('A', 'v')],
    '4': [('1', 'v'), ('5', '>'), ('7', '^')],
    '5': [('2', 'v'), ('4', '<'), ('6', '>'), ('8', '^')],
    '6': [('3', 'v'), ('5', '<'), ('9', '^')],
    '7': [('4', 'v'), ('8', '>')],
    '8': [('5', 'v'), ('7', '<'), ('9', '>')],
    '9': [('6', 'v'), ('8', '<')],
    '0': [('2', '^'), ('A', '>')],
    'A': [('0', '<'), ('3', '^')]
}

symbol_keypad: Dict[str, List[Tuple[str, str]]] = {
    '^': [('A', '>'), ('v', 'v')],
    'A': [('^', '<'), ('>', 'v')],
    '<': [('v', '>')],
    'v': [('<', '<'), ('^', '^'), ('>','>')],
    '>': [('v', '<'), ('A', '^')],
}



def shortest_path(start: Tuple[str,str], end: str, grid:  Dict[str, List[Tuple[str, str]]]) -> List[str]:
    queue = deque([(start, [start])])
    visited = set() 

    while queue:
        current, path = queue.popleft()

        if current[0] == end:
            return path

        visited.add(current[0])

        for neighbor in grid.get(current[0], []):
            if neighbor[0] not in visited:
                queue.append((neighbor, path + [neighbor]))

    return []

def full_numeric_path(route: str) -> List[str]:
    full_path = []
    for i in range(len(route)-1):
        start = route[i]
        end = route[i+1]
        path = shortest_path((start, "."), end, numeric_keypad)
        symbols = [i[1] for i in path if i[1] != '.']
        full_path = full_path + symbols + ["A"]
    return full_path


def full_symbol_path(route: List[str]) -> List[str]:
    full_path = []
    for i in range(len(route)-1):
        start = route[i]
        end = route[i+1]
        path = shortest_path((start, "."), end, symbol_keypad)
        symbols = [i[1] for i in path if i[1] != '.']
        full_path = full_path + symbols + ["A"]
    return full_path

with open('day21/day21.txt') as f:
    for line in f:
        line = line.strip()
        line = 'A' + line
        numeric_path = full_numeric_path(line)
        numeric_path.insert(0,'A')
        print(numeric_path)
        symbol_path = full_symbol_path(numeric_path)
        print(symbol_path)
        
        
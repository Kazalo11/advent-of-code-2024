from collections import deque
from functools import lru_cache
from typing import Any, Dict, List, Tuple

numeric_grid = {
    "0": [("A", ">"), ("2", "^")], 
    "A": [("3", "^"), ("0", "<")], 
    "1": [("2", ">"), ("4", "^")], 
    "2": [("5", ">"), ("3", "^"), ("1", "<"), ("0", "v")], 
    "3": [("6", ">"), ("2", "v"), ("A", "<")], 
    "4": [("7", ">"), ("5", "^"), ("1", "v")],  
    "5": [("8", ">"), ("6", "^"), ("4", "v"), ("2", "<")],  
    "6": [("9", ">"), ("5", "^"), ("3", "v")], 
    "7": [("8", ">"), ("4", "^")], 
    "8": [("9", ">"), ("7", "^"), ("5", "v")],
    "9": [("8", "<"), ("6", "v")], 
}

directional_grid = {
    "^": ["<", ">", "A", "v"],
    "A": ["^", "v"],
    "<": ["^", "v"],
    "v": ["<", ">", "^", "A"],
    ">": ["v", "^"],
}

from typing import Dict, List


def shortest_path(start: str, end: str, grid: Dict[str, List[str]]) -> List[str]:
    queue = deque([(start, [start])])
    visited = set() 

    while queue:
        current, path = queue.popleft()

        if current == end:
            return path

        visited.add(current)

        for neighbor in grid.get(current, []):
            if neighbor not in visited:
                queue.append((neighbor, path + [neighbor]))

    return []

def full_path(route: str) -> List[str]:
    total_path = [route[0]]
    for i in range(len(route)-1):
        path = shortest_path(route[i], route[i+1], numeric_grid)
        total_path = total_path + path[1:] + ["*"]
    return total_path

with open("day21/day21.txt") as f:
    for line in f:
        line = line.strip()
        line = 'A'+line
        total_path = full_path(line)
        print(total_path)
        
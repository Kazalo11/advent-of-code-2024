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



from collections import deque
from functools import lru_cache
from typing import Dict, List, Tuple


@lru_cache(None) 
def cached_paths(start: Tuple[str, str], end: str, grid: Tuple[Tuple[str, str]]) -> List[List[str]]:
    grid_dict = {key: [tuple(neighbor) for neighbor in value] for key, value in grid}

    queue = deque([(start, [start])])
    visited = set()
    shortest_paths = []
    shortest_distance = float('inf')

    while queue:
        current, path = queue.popleft()

        if len(path) > shortest_distance:
            continue

        if current[0] == end:
            path.append(("A", "A"))
            if len(path) < shortest_distance:
                shortest_paths = [path]
                shortest_distance = len(path)
            elif len(path) == shortest_distance:
                shortest_paths.append(path)

        visited.add(current[0])

        for neighbor in grid_dict.get(current[0], []):
            if neighbor[0] not in visited:
                queue.append((neighbor, path + [neighbor]))

    return shortest_paths

def shortest_paths(start: Tuple[str, str], end: str, grid: Dict[str, List[Tuple[str, str]]]) -> List[List[str]]:
    grid_tuple = tuple((key, tuple(value)) for key, value in grid.items())
    return cached_paths(start, end, grid_tuple)


def full_path(route: str, grid=numeric_keypad) -> List[List[str]]:
    total_paths = []

    for i in range(len(route) - 1):
        start = route[i]
        end = route[i + 1]
        
        paths = shortest_paths((start, "."), end, grid)
        new_paths = []
        for path in paths:
            if len(total_paths) == 0:
                new_paths = paths
            else:
                for i in total_paths:
                        new_path = i + path
                        new_paths.append(new_path)
        total_paths = new_paths

    ans = []

    for total_path in total_paths:
        new = [j[1] for j in total_path if j[1] != '.']
        ans.append(new)

    return ans




def get_min_length_lists(lists: List[List[str]]) -> List[List[str]]:
    min_length = min(len(lst) for lst in lists)
    return [lst for lst in lists if len(lst) == min_length]


def extract_number(item: str) -> int:
    return int(item[:-1])

count =0

with open('day21/day21.txt') as f:
    for line in f:
        line = line.strip()
        num = extract_number(line)
        line = 'A' + line
        numeric_paths = full_path(line)
        min_numeric_paths = get_min_length_lists(numeric_paths)

        for item in min_numeric_paths:
            item.insert(0,"A")

        total_symbol_paths = []
        for item in min_numeric_paths:
            symbols_path = full_path(item, grid=symbol_keypad)
            total_symbol_paths+=symbols_path
        min_symbols_paths = get_min_length_lists(total_symbol_paths)
        for item in min_symbols_paths:
            item.insert(0,"A")
        final_symbols_paths = []
        for item in min_symbols_paths:
            final_path = full_path(item, grid=symbol_keypad)
            final_symbols_paths+=final_path
        min_final_paths = get_min_length_lists(final_symbols_paths)
        final_length = len(min_final_paths[0])
        print(num,final_length)
        count= count + (final_length*num)

print(count)
        
        
        
        
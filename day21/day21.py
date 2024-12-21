from functools import lru_cache
from typing import Any, Dict, List, Tuple

numeric_keypad: Dict[Tuple[int, int], str] = {
    (0, 0): "7", 
    (1, 0): "8", 
    (2, 0): "9", 
    (0, 1): "4", 
    (1, 1): "5", 
    (2, 1): "6", 
    (0, 2): "1", 
    (1, 2): "2", 
    (2, 2): "3", 
    (3, 1): "0", 
    (3, 2): "A"
}

directional_keypad: Dict[Tuple[int, int], str] = {
    (1, 0): "^", 
    (2, 0): "A", 
    (0, 1): "<", 
    (1, 1): "v", 
    (1, 2): ">"
}


total_path: Dict[str, str] = {}

def find_path(route: List[str], keypad: Dict[Tuple[int, int], str], index: int, coord: Tuple[int, int], path: List[str]) -> List[Tuple[int, int]]:
    if coord not in list(keypad.keys()):
        return []
    next_item = route[index]
    curr_item = keypad[coord]
    if curr_item in path:
        return []
    if next_item == curr_item:
        if index == len(route):
            path.append('A')
            route2 = ''.join(route)
            total_path[route2] = path
            return path
        index+=1 
    path.append(curr_item)
    left_coord = (coord[0]-1, coord[1])
    right_coord = (coord[0]+1, coord[1])
    up_coord = (coord[0], coord[1]-1)
    down_coord = (coord[0], coord[1]+1)
    find_path(route, keypad, index,left_coord, path )
    find_path(route, keypad, index,right_coord, path )
    find_path(route, keypad, index,up_coord, path )
    find_path(route, keypad, index,down_coord, path )

def generate_paths(route: List[str], keypad: Dict[Tuple[int, int], str]) -> None:
    find_path(route, keypad, 0, (3,2), [])

with open("day21/day21.txt") as f:
    for line in f:
        line = line.strip()
        route = [i for i in line]
        generate_paths(line, keypad=numeric_keypad)
    print(total_path)

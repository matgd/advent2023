#!/usr/bin/env python3.11

from itertools import combinations


INPUT = "itt9724940nput.txt"

with open(INPUT, "r") as f:
    input = f.read().splitlines()

def expand_horizontally(input_lines: list[str]) -> list[str]:
    expanded_lines = []
    for line in input_lines:
        if not "#" in line:
            expanded_lines.append(line)
        expanded_lines.append(line)
    return expanded_lines

def expand_vertically(input_lines: list[str]) -> list[str]:
    expanded_lines = []
    vert = [''.join(list(l)) for l in zip(*input_lines)]
    for vline in vert:
        if not "#" in vline:
            expanded_lines.append(vline)
        expanded_lines.append(vline)
    return [''.join(list(l)) for l in zip(*expanded_lines)]

def manhattan_distance(a: tuple[int, int], b: tuple[int, int]) -> int:
    return abs(a[0] - b[0]) + abs(a[1] - b[1])

space = expand_vertically(expand_horizontally(input))
galaxies = [(i, j) for i in range(len(space)) for j in range(len(space[i])) if space[i][j] == "#"]
print(sum(manhattan_distance(a, b) for a, b in combinations(galaxies, 2)))

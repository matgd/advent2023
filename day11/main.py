#!/usr/bin/env python3.11

from itertools import combinations
import bisect

PART = 2
INPUT = "input.txt"

with open(INPUT, "r") as f:
    input = f.read().splitlines()

horizontal_plus_million = []
vertical_plus_million = []

def expand_horizontally(input_lines: list[str]) -> list[str]:
    for i, line in enumerate(input_lines):
        if not "#" in line:
            horizontal_plus_million.append(i)
    return input_lines

def expand_vertically(input_lines: list[str]) -> list[str]:
    vert = [''.join(list(l)) for l in zip(*input_lines)]
    for i, vline in enumerate(vert):
        if not "#" in vline:
            vertical_plus_million.append(i)
    return input_lines

def manhattan_distance(a: tuple[int, int], b: tuple[int, int]) -> int:
    return abs(a[0] - b[0]) + abs(a[1] - b[1])

expand_vertically(input)
expand_horizontally(input)
galaxies = [(i, j) for i in range(len(input)) for j in range(len(input[i])) if input[i][j] == "#"]

galaxies2 = []
for g in galaxies:
    mult = 1 if PART == 1 else 999_999
    plus_horizontal = bisect.bisect(horizontal_plus_million, g[0]) * mult
    plus_vertical = bisect.bisect(vertical_plus_million, g[1]) * mult
    galaxies2.append((g[0] + plus_horizontal, g[1] + plus_vertical))

print(sum(manhattan_distance(a, b) for a, b in combinations(galaxies2, 2)))

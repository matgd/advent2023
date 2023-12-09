#!/usr/bin/env python3.11

INPUT = "input.txt"

with open(INPUT, "r") as f:
    lines = f.readlines()

histories = []
for line in lines:
    histories.append([int(x) for x in line.split(" ")])

# Part 1
def part1(history):
    if not any(history):
        return 0
    sub_history = []
    for i in range(len(history) - 1):
        sub_history.append(history[i+1] - history[i])
    return history[-1] + part1(sub_history)

# Part 2
def part2(history):
    if not any(history):
        return 0
    sub_history = []
    for i in range(len(history) - 1):
        sub_history.append(history[i+1] - history[i])
    return history[-1] + part1(sub_history)

print("[P1]", 114, sum((part1(h) for h in histories)))
for h in histories:
    h.reverse()
print("[P2]", 2, sum((part1(h) for h in histories)))


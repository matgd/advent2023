INPUT = "input.txt"

with open(INPUT, 'r') as f:
    lines = f.readlines()

GO_UP = (-1, 0)
GO_DOWN = (1, 0)
GO_LEFT = (0, -1)
GO_RIGHT = (0, 1)

FROM_UP_TO_DOWN = (1, 0)
FROM_LEFT_TO_RIGHT = (0, 1)
FROM_DOWN_TO_UP = (-1, 0)
FROM_RIGHT_TO_LEFT = (0, -1)

pipes_connect_offset = {
    # direction
    # symbol, up_to_down, left_to_right, down_to_up, right_to_left
    '|': {
        FROM_UP_TO_DOWN: GO_DOWN,
        FROM_LEFT_TO_RIGHT: (0, 0),
        FROM_DOWN_TO_UP: GO_UP,
        FROM_RIGHT_TO_LEFT: (0, 0),
    },
    '-': {
        FROM_UP_TO_DOWN: (0, 0),
        FROM_LEFT_TO_RIGHT: GO_RIGHT,
        FROM_DOWN_TO_UP: (0, 0),
        FROM_RIGHT_TO_LEFT: GO_LEFT,
    },
    'L': {
        FROM_UP_TO_DOWN: GO_RIGHT,
        FROM_LEFT_TO_RIGHT: (0, 0),
        FROM_DOWN_TO_UP: (0, 0),
        FROM_RIGHT_TO_LEFT: GO_UP,
    },
    'J': {
        FROM_UP_TO_DOWN: GO_LEFT,
        FROM_LEFT_TO_RIGHT: GO_UP,
        FROM_DOWN_TO_UP: (0, 0),
        FROM_RIGHT_TO_LEFT: (0, 0),
    },
    '7': {
        FROM_UP_TO_DOWN: (0, 0),
        FROM_LEFT_TO_RIGHT: GO_DOWN,
        FROM_DOWN_TO_UP: GO_LEFT,
        FROM_RIGHT_TO_LEFT: (0, 0),
    },
    'F': {
        FROM_UP_TO_DOWN: (0, 0),
        FROM_LEFT_TO_RIGHT: (0, 0),
        FROM_DOWN_TO_UP: GO_RIGHT,
        FROM_RIGHT_TO_LEFT: GO_DOWN,
    },
}

# Find 'S' coordinates in lines
s = (-1, -1)
for i, line in enumerate(lines):
    if 'S' in line:
        s = (i, line.index('S'))

pipe_coords = s
pipe_type = 'F' if INPUT == 'tinput.txt' else 'J'

first_path_coords, second_path_coords = list(s), list(s)
first_direction, second_direction = (0, 0), (0, 0)

if pipe_type == 'F':
    first_direction = GO_RIGHT
    second_direction = GO_DOWN
if pipe_type == 'J':
    first_direction = GO_UP
    second_direction = GO_LEFT

steps = 0
# Find first path
while True:
    first_path_coords[0] += first_direction[0]
    first_path_coords[1] += first_direction[1]

    second_path_coords[0] += second_direction[0]
    second_path_coords[1] += second_direction[1]

    first_symbol = lines[first_path_coords[0]][first_path_coords[1]]
    second_symbol = lines[second_path_coords[0]][second_path_coords[1]]

    print(first_symbol, second_symbol)

    first_direction = pipes_connect_offset[first_symbol][first_direction]
    second_direction = pipes_connect_offset[second_symbol][second_direction]
    steps += 1
    if first_path_coords == second_path_coords:
        break

print(steps)

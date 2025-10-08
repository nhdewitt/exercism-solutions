def annotate(garden):
    if not garden:
        return []

    if not all(len(row) == len(garden[0]) for row in garden):
        raise ValueError("The board is invalid with current input.")

    rows = len(garden)
    cols = len(garden[0])

    flowers = set()
    for i, row in enumerate(garden):
        for j, cell in enumerate(row):
            if cell == "*":
                flowers.add((i, j))
            elif cell != " ":
                raise ValueError("The board is invalid with current input.")

    res = []
    for i in range(rows):
        row = ""
        for j in range(cols):
            if (i, j) in flowers:
                row += "*"
            else:
                count = count_adjacent_flowers(i, j, flowers, rows, cols)
                row += str(count) if count > 0 else " "
        res.append(row)

    return res

def count_adjacent_flowers(row, col, flowers, rows, cols):
    count = 0
    for dr in [-1, 0, 1]:
        for dc in [-1, 0, 1]:
            if dr == 0 and dc == 0:
                continue
            neighbor = (row + dr, col + dc)
            if (0 <= neighbor[0] < rows and
                0 <= neighbor[1] < cols and
                neighbor in flowers):
                count += 1

    return count
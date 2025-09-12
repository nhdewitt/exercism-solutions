def saddle_points(matrix):
    if not matrix:
        return []

    n_rows = len(matrix)
    n_cols = len(matrix[0])
    if any(len(row) != n_cols for row in matrix):
        raise ValueError("irregular matrix")
    if n_cols == 0:
        return []

    col_mins = list(matrix[0])
    for r in range(1, n_rows):
        row = matrix[r]
        for c in range(n_cols):
            if row[c] < col_mins[c]:
                col_mins[c] = row[c]

    res = []
    for i, row in enumerate(matrix):
        row_max = max(row)
        for j, val in enumerate(row):
            if val == row_max and val == col_mins[j]:
                res.append({"row": i + 1, "column": j + 1})

    return res
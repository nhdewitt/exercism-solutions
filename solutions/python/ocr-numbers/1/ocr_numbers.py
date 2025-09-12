def convert(input_grid):
    templates = [
        [" _ ", "| |", "|_|", "   "], # 0
        ["   ", "  |", "  |", "   "], # 1
        [" _ ", " _|", "|_ ", "   "], # 2
        [" _ ", " _|", " _|", "   "], # 3
        ["   ", "|_|", "  |", "   "], # 4
        [" _ ", "|_ ", " _|", "   "], # 5
        [" _ ", "|_ ", "|_|", "   "], # 6
        [" _ ", "  |", "  |", "   "], # 7
        [" _ ", "|_|", "|_|", "   "], # 8
        [" _ ", "|_|", " _|", "   "]  # 9
    ]
    ocr = {"".join(g): str(i) for i, g in enumerate(templates)}

    if len(input_grid) % 4 != 0:
        raise ValueError("Number of input lines is not a multiple of four")
    width = len(input_grid[0])
    if any(len(row) != width for row in input_grid):
        raise ValueError("Number of input rows is inconsistent")
    if width % 3 != 0:
        raise ValueError("Number of input columns is not a multiple of three")

    rows_out = []
    for r in range(0, len(input_grid), 4):
        block = input_grid[r:r + 4]
        chunks = [[line[c:c + 3] for c in range(0, width, 3)] for line in block]
        digits = list(zip(*chunks))
        rows_out.append("".join(ocr.get("".join(glyph), "?") for glyph in digits))

    return ",".join(rows_out)
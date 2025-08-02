def adjacency_check(mine_dict: dict, coords: tuple, edge_row=None, corner=(False, None)):
    offsets = [[(-1, -1),(-1, 0),(-1, 1),(0, -1),(0, 1),(1, -1),(1, 0),(1, 1)], #[0] - 8
               [(0, -1),(0, 1),(1, -1),(1, 0),(1, 1)],                          #[1] - top row
               [(-1, -1),(-1, 0),(-1, 1),(0, -1),(0, 1)],                       #[2] - bottom row
               [(0, 1),(1, 0),(1, 1)],                                          #[3] - top left corner
               [(-1, 0),(-1, 1),(0, 1)],                                        #[4] - bottom left corner
               [(0, -1),(1, -1),(1, 0)],                                        #[5] - top right corner
               [(-1, -1),(-1, 0),(0, -1)],                                      #[6] - bottom right corner
               [(-1, 0),(-1, 1),(0, 1),(1, 0),(1,1)],                           #[7] - left side
               [(-1, -1),(-1, 0),(0, -1),(1, -1),(1,0)]]                        #[8] - right side

    x, y = coords[0], coords[1]
    
    if edge_row == "top":
        if corner[0]:
            corner_loc = corner[1]
            if corner_loc == "left":
                offset = offsets[3]
            elif corner_loc == "right":
                offset = offsets[5]
        else:
            offset = offsets[1]
    elif edge_row == "bottom":
        if corner[0]:
            corner_loc = corner[1]
            if  corner_loc == "left":
                offset = offsets[4]
            elif corner_loc == "right":
                offset = offsets[6]
        else:
            offset = offsets[2]
    elif edge_row == "left":
        offset = offsets[7]
    elif edge_row == "right":
        offset = offsets[8]
    else:
        offset = offsets[0]

    for (dx, dy) in offset:
        neighbor = (x + dx, y + dy)
        if isinstance(mine_dict[neighbor], int):
            mine_dict[neighbor] += 1
        else:
            continue
        

def annotate(minefield):
    if minefield == []:
        return []
        
    symmetry_check = len(minefield[0])
    right_end = len(minefield[0]) - 1
    bottom_row = len(minefield) - 1
    oned_horizontal = False
    oned_vertical = False
    valid_chars = [" ","*"]

    offsets = [[(-1, -1),(-1, 0),(-1, 1),(0, -1),(0, 1),(1, -1),(1, 0),(1, 1)], #[0] - 8
               [(0, -1),(0, 1),(1, -1),(1, 0),(1, 1)],                          #[1] - top row
               [(-1, -1),(-1, 0),(-1, 1),(0, -1),(0, 1)],                       #[2] - bottom row
               [(0, 1),(1, 0),(1, 1)],                                          #[3] - top left corner
               [(-1, 0),(-1, 1),(0, 1)],                                        #[4] - bottom left corner
               [(0, -1),(1, -1),(1, 0)],                                        #[5] - top right corner
               [(-1, -1),(-1, 0),(0, -1)],                                      #[6] - bottom right corner
               [(-1, 0),(-1, 1),(0, 1),(1, 0),(1,1)],                           #[7] - left side
               [(-1, -1),(-1, 0),(0, -1),(1, -1),(1,0)]]                        #[8] - right side

    minefield_dict = {}
    mines = []
    for i, row in enumerate(minefield):
        if len(row) != symmetry_check:
            raise ValueError("The board is invalid with current input.")
        for j, item in enumerate(minefield[i]):
            if minefield[i][j] == " ":
                minefield_dict[(i, j)] = 0
            elif minefield[i][j] == "*":
                minefield_dict[(i, j)] = "*"
                mines.append((i, j))
            else:
                raise ValueError("The board is invalid with current input.")
    
    if len(minefield) == 1:
        oned_horizontal = True
    elif len(minefield[0]) == 1:
        oned_vertical = True
    
    if not oned_horizontal and not oned_vertical:
        for index, coords in mines:
            x, y = index, coords
            if y == 0 or y == right_end:
                if x == 0 and y == 0:
                    adjacency_check(minefield_dict, (x, y), "top", (True, "left"))
                elif x == bottom_row and y == 0:
                    adjacency_check(minefield_dict, (x, y), "bottom", (True, "left"))
                elif x == 0 and y == right_end:
                    adjacency_check(minefield_dict, (x, y), "top", (True, "right"))
                elif x == bottom_row and y == right_end:
                    adjacency_check(minefield_dict, (x, y), "bottom", (True, "right"))
                elif y == 0:
                    adjacency_check(minefield_dict, (x, y), "left")
                elif y == right_end:
                    adjacency_check(minefield_dict, (x, y), "right")
            elif x == 0 or x == bottom_row:
                if x == 0:
                    adjacency_check(minefield_dict, (x, y), "top")
                else:
                    adjacency_check(minefield_dict, (x, y), "bottom")
            else:
                adjacency_check(minefield_dict, (x, y))
    else:
        if oned_horizontal:
            for index, coords in mines:
                x, y = index, coords
                if y == 0:
                    if isinstance(minefield_dict[(x, y + 1)], int):
                        minefield_dict[(x, y + 1)] += 1
                elif y == right_end:
                    if isinstance(minefield_dict[(x, y - 1)], int):
                        minefield_dict[(x, y - 1)] += 1
                else:
                    if isinstance(minefield_dict[(x, y - 1)], int):
                        minefield_dict[(x, y - 1)] += 1
                    if isinstance(minefield_dict[(x, y + 1)], int):
                        minefield_dict[(x, y + 1)] += 1
        else:
            for index, coords in mines:
                x, y = index, coords
                if x == 0:
                    if isinstance(minefield_dict[(x + 1, y)], int):
                        minefield_dict[(x + 1, y)] += 1
                elif x == bottom_row:
                    if isinstance(minefield_dict[(x - 1, y)], int):
                        minefield_dict[(x - 1, y)] += 1
                else:
                    if isinstance(minefield_dict[(x - 1, y)], int):
                        minefield_dict[(x - 1, y)] += 1
                    if isinstance(minefield_dict[(x + 1, y)], int):
                        minefield_dict[(x + 1, y)] += 1

    swept = []
    for i, row in enumerate(minefield):
        swept_row = ""
        for j, item in enumerate(minefield[i]):
            if type(minefield_dict[i, j]) is int:
                if minefield_dict[i, j] == 0:
                    swept_row += " "
                else:
                    swept_row += str(minefield_dict[i, j])
            else:
                swept_row += minefield_dict[i, j]
        swept.append(swept_row)

    return swept
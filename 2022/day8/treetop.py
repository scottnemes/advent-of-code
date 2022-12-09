# load list of trees
def import_trees(file):
    with open(file) as f:
        lines = f.readlines()
    return lines


# create grid of trees using a 2d list
def get_tree_grid(tree_raw, cols, rows):
    # tree_grid[y][x]
    tree_grid = [[0 for i in range(cols)] for j in range(rows)]
    x = 0
    y = 0
    for row in tree_raw:
        row = row.strip()
        for height in row:
            tree_grid[y][x] = height
            x += 1
        x = 0
        y += 1
    return tree_grid


# find which trees are visible from the outside
# set is used to enforce uniqueness to avoid double counting trees
def check_visible_trees(tree_grid):
    visible_trees = set()
    rows = len(tree_grid)
    cols = len(tree_grid[0])
    # check rows from left to right
    tree_height = 0
    for y in range(rows):
        max_tree_height = 0
        for x in range(cols):
            tree_height = int(tree_grid[y][x])
            # if in the first or last row add each x,y pair
            if y == 0 or y == (rows-1):
                visible_trees.add(f"{x}-{y}")
            # if in the first or last column add each x,y pair
            elif x == 0 or x == (cols-1):
                visible_trees.add(f"{x}-{y}")
            # get next tree
            # should not be reached if we hit an edge first above
            elif tree_height > max_tree_height:
                visible_trees.add(f"{x}-{y}")
            max_tree_height = max(max_tree_height, tree_height)
            #print(f"({x},{y}) {tree_height} ({max_tree_height})")

    # check rows from right to left
    tree_height = 0
    for y in range(rows):
        max_tree_height = 0
        for x in range(cols-1,-1,-1):
            tree_height = int(tree_grid[y][x])
            # if in the first or last row add each x,y pair
            if y == 0 or y == (rows-1):
                visible_trees.add(f"{x}-{y}")
            # if in the first or last column add each x,y pair
            elif x == 0 or x == (cols-1):
                visible_trees.add(f"{x}-{y}")
            # get next tree
            # should not be reached if we hit an edge first above
            elif tree_height > max_tree_height:
                visible_trees.add(f"{x}-{y}")
            max_tree_height = max(max_tree_height, tree_height)
    # check cols from top to bottom
    tree_height = 0
    for x in range(cols):
        max_tree_height = 0
        for y in range(rows):
            tree_height = int(tree_grid[y][x])
            # if in the first or last row add each x,y pair
            if y == 0 or y == (rows-1):
                visible_trees.add(f"{x}-{y}")
            # if in the first or last column add each x,y pair
            elif x == 0 or x == (cols-1):
                visible_trees.add(f"{x}-{y}")
            # get next tree
            # should not be reached if we hit an edge first above
            elif tree_height > max_tree_height:
                visible_trees.add(f"{x}-{y}")
            max_tree_height = max(max_tree_height, tree_height)
    # check cols from bottom to top
    tree_height = 0
    for x in range(cols):
        max_tree_height = 0
        for y in range(rows-1,-1,-1):
            tree_height = int(tree_grid[y][x])
            # if in the first or last row add each x,y pair
            if y == 0 or y == (rows-1):
                visible_trees.add(f"{x}-{y}")
            # if in the first or last column add each x,y pair
            elif x == 0 or x == (cols-1):
                visible_trees.add(f"{x}-{y}")
            # get next tree
            # should not be reached if we hit an edge first above
            elif tree_height > max_tree_height:
                visible_trees.add(f"{x}-{y}")
            max_tree_height = max(max_tree_height, tree_height)
    return visible_trees                 


# calculate the scenic view score
def calculate_view_score(tree_grid):
    rows = len(tree_grid)
    cols = len(tree_grid[0])
    view_scores = {}
    for y in range(rows):
        for x in range(cols):
            left_score = 0
            right_score = 0
            up_score = 0
            down_score = 0
            tree_height = int(tree_grid[y][x])
            # if the tree is on a edge, the score will be zero
            if x == 0 or x == (cols-1):
                view_scores[f"{x}-{y}"] = 0
                continue
            if y == 0 or y == (rows-1):
                view_scores[f"{x}-{y}"] = 0
                continue
            # look left
            for x_sub in range(x-1,-1,-1):
                next_tree_height = int(tree_grid[y][x_sub])
                if next_tree_height >= tree_height:
                    left_score += 1
                    break
                else:
                    left_score += 1
            # look right
            for x_sub in range(x+1,cols,1):
                next_tree_height = int(tree_grid[y][x_sub])
                if next_tree_height >= tree_height:
                    right_score += 1
                    break
                else:
                    right_score += 1
            # look up
            for y_sub in range(y-1,-1,-1):
                next_tree_height = int(tree_grid[y_sub][x])
                if next_tree_height >= tree_height:
                    up_score += 1
                    break
                else:
                    up_score += 1
            # look down
            for y_sub in range(y+1,rows,1):
                next_tree_height = int(tree_grid[y_sub][x])
                if next_tree_height >= tree_height:
                    down_score += 1
                    break
                else:
                    down_score += 1
            view_scores[f"{x}-{y}"] = left_score * right_score * up_score * down_score
    best_view_tree = ""
    best_view_score = 0
    for tree, view_score in view_scores.items():
        if view_score > best_view_score:
            best_view_score = view_score
            best_view_tree = tree
    return [best_view_tree, best_view_score]


def main():
    grid_cols = 99
    grid_rows = 99
    tree_raw = import_trees("input.txt")
    tree_grid = get_tree_grid(tree_raw, grid_cols, grid_rows)
    visible_trees = check_visible_trees(tree_grid)
    print(f"Number of trees visible from the outside (part 1): {len(visible_trees)}")
    tree_view_scores = calculate_view_score(tree_grid)
    print(f"Tree with best scenic view score (part 2): {tree_view_scores[0]} ({tree_view_scores[1]})")


if __name__ == '__main__':
    main()
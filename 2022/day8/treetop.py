"""
--- Day 8: Treetop Tree House ---

--- Part One ---
The expedition comes across a peculiar patch of tall trees all planted carefully in a grid. The Elves explain that a previous expedition planted these trees as a reforestation effort. Now, they're curious if this would be a good location for a tree house.

First, determine whether there is enough tree cover here to keep a tree house hidden. To do this, you need to count the number of trees that are visible from outside the grid when looking directly along a row or column.

The Elves have already launched a quadcopter to generate a map with the height of each tree (your puzzle input). For example:

30373
25512
65332
33549
35390
Each tree is represented as a single digit whose value is its height, where 0 is the shortest and 9 is the tallest.

A tree is visible if all of the other trees between it and an edge of the grid are shorter than it. Only consider trees in the same row or column; that is, only look up, down, left, or right from any given tree.

All of the trees around the edge of the grid are visible - since they are already on the edge, there are no trees to block the view. In this example, that only leaves the interior nine trees to consider:

The top-left 5 is visible from the left and top. (It isn't visible from the right or bottom since other trees of height 5 are in the way.)
The top-middle 5 is visible from the top and right.
The top-right 1 is not visible from any direction; for it to be visible, there would need to only be trees of height 0 between it and an edge.
The left-middle 5 is visible, but only from the right.
The center 3 is not visible from any direction; for it to be visible, there would need to be only trees of at most height 2 between it and an edge.
The right-middle 3 is visible from the right.
In the bottom row, the middle 5 is visible, but the 3 and 4 are not.
With 16 trees visible on the edge and another 5 visible in the interior, a total of 21 trees are visible in this arrangement.

Consider your map; how many trees are visible from outside the grid?

--- Part Two ---
Content with the amount of tree cover available, the Elves just need to know the best spot to build their tree house: they would like to be able to see a lot of trees.

To measure the viewing distance from a given tree, look up, down, left, and right from that tree; stop if you reach an edge or at the first tree that is the same height or taller than the tree under consideration. (If a tree is right on the edge, at least one of its viewing distances will be zero.)

The Elves don't care about distant trees taller than those found by the rules above; the proposed tree house has large eaves to keep it dry, so they wouldn't be able to see higher than the tree house anyway.

In the example above, consider the middle 5 in the second row:

30373
25512
65332
33549
35390
Looking up, its view is not blocked; it can see 1 tree (of height 3).
Looking left, its view is blocked immediately; it can see only 1 tree (of height 5, right next to it).
Looking right, its view is not blocked; it can see 2 trees.
Looking down, its view is blocked eventually; it can see 2 trees (one of height 3, then the tree of height 5 that blocks its view).
A tree's scenic score is found by multiplying together its viewing distance in each of the four directions. For this tree, this is 4 (found by multiplying 1 * 1 * 2 * 2).

However, you can do even better: consider the tree of height 5 in the middle of the fourth row:

30373
25512
65332
33549
35390
Looking up, its view is blocked at 2 trees (by another tree with a height of 5).
Looking left, its view is not blocked; it can see 2 trees.
Looking down, its view is also not blocked; it can see 1 tree.
Looking right, its view is blocked at 2 trees (by a massive tree of height 9).
This tree's scenic score is 8 (2 * 2 * 1 * 2); this is the ideal spot for the tree house.

Consider each tree on your map. What is the highest scenic score possible for any tree?
"""


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
from functools import reduce
from typing import Iterator
from copy import deepcopy
import pprint as pp


class InvalidPuzzleError(Exception):
    pass


def is_solved(puzzle: list[list[int]]) -> bool:
    return 0 not in reduce(
        lambda accumulator, value: accumulator | value, map(set, puzzle), set()
    )


def get_box_set(puzzle: list[list[int]], row_index: int, col_index: int) -> set[int]:
    row_min = (row_index // 3) * 3
    col_min = (col_index // 3) * 3
    return {
        value
        for row in puzzle[row_min : row_min + 3]
        for value in row[col_min : col_min + 3]
    }


def get_possible_cell_values(
    puzzle: list[list[int]], row_index: int, col_index: int
) -> Iterator[int]:
    for value in range(1, 10):
        if (
            (value not in puzzle[row_index])
            and (value not in {row[col_index] for row in puzzle})
            and (value not in get_box_set(puzzle, row_index, col_index))
        ):
            yield value


max_level = 0


def solve(puzzle: list[list[int]], level: int) -> list[list[list[int]]]:
    global max_level
    # print(f"solve() called, {level=}, {max_level=}")
    if max_level < level:
        max_level = level

    # pp.pprint(puzzle)
    solutions = []
    for row_index in range(9):
        for col_index in range(9):
            if puzzle[row_index][col_index] == 0:
                # values = list(get_possible_cell_values(puzzle, row_index, col_index))
                # print(f"We have {len(values)=}, {values=}")
                for new_value in get_possible_cell_values(puzzle, row_index, col_index):
                    # _puzzle = deepcopy(puzzle)
                    _puzzle = tuple(
                        map(
                            lambda _row_index_and_row: tuple(
                                value
                                if (_row_index_and_row[0], _col_index)
                                != (row_index, col_index)
                                else new_value
                                for _col_index, value in enumerate(
                                    _row_index_and_row[1]
                                )
                            ),
                            enumerate(puzzle),
                        )
                    )
                    # _puzzle[row_index][col_index] = new_value
                    # print(f"Calling solve() with {row_index=}, {col_index=}, {new_value=}")
                    if is_solved(_puzzle):
                        solutions.append(_puzzle)
                    else:
                        for solution in solve(_puzzle, level + 1):
                            solutions.append(solution)
                else:
                    return solutions

    return solutions


def sudoku_solver(puzzle: list[list[int]]) -> list[list[int]]:
    # It should raise an error in cases of:
    # invalid grid (not 9x9, cell with values not in the range 1~9);
    # multiple solutions for the same puzzle or the puzzle is unsolvable
    if len(puzzle) != 9 or any(len(row) != 9 for row in puzzle):
        raise InvalidPuzzleError("Invalid puzzle shape")
    for row in puzzle:
        for item in row:
            if not 0 <= item <= 9:
                raise InvalidPuzzleError(
                    "Cell value out of range 1~9 and 0 for empty cells"
                )
    solutions = solve(tuple(map(tuple, puzzle)), 1)
    if not solutions:
        raise InvalidPuzzleError("Unsolvable puzzle")
    elif len(solutions) > 1:
        raise InvalidPuzzleError("Puzzle has more than 1 solution")
    return list(map(list, solutions[0]))


if __name__ == "__main__":
    puzzle = [
        [0, 0, 6, 1, 0, 0, 0, 0, 8],
        [0, 8, 0, 0, 9, 0, 0, 3, 0],
        [2, 0, 0, 0, 0, 5, 4, 0, 0],
        [4, 0, 0, 0, 0, 1, 8, 0, 0],
        [0, 3, 0, 0, 7, 0, 0, 4, 0],
        [0, 0, 7, 9, 0, 0, 0, 0, 3],
        [0, 0, 8, 4, 0, 0, 0, 0, 6],
        [0, 2, 0, 0, 5, 0, 0, 8, 0],
        [1, 0, 0, 0, 0, 2, 5, 0, 0],
    ]

    solution = [
        [3, 4, 6, 1, 2, 7, 9, 5, 8],
        [7, 8, 5, 6, 9, 4, 1, 3, 2],
        [2, 1, 9, 3, 8, 5, 4, 6, 7],
        [4, 6, 2, 5, 3, 1, 8, 7, 9],
        [9, 3, 1, 2, 7, 8, 6, 4, 5],
        [8, 5, 7, 9, 4, 6, 2, 1, 3],
        [5, 9, 8, 4, 1, 3, 7, 2, 6],
        [6, 2, 4, 7, 5, 9, 3, 8, 1],
        [1, 7, 3, 8, 6, 2, 5, 9, 4],
    ]

    zeros = 0
    for row in puzzle:
        for i in row:
            if i == 0:
                zeros += 1
    print(f"{zeros=}")

    solver_output = sudoku_solver(puzzle)
    print("** Solution found! **")
    pp.pprint(solver_output)
    assert solver_output == solution, "Failed to solve sample sudoku"
import pytest
from dataclasses import dataclass
from collections import deque
import math


class InvalidMazeShapeException(Exception):
    """Maze should be a square"""

    pass


def path_finder(maze: str) -> bool | int:
    """
    Returns number of steps in the shortest path from the labyrinth
    If no path exists - returns False
    Uses BFS - complexity O(N) where N is the maze size
    """
    if not maze:
        return False

    maze = maze.replace("\n", "")
    side_length = math.floor(math.sqrt(len(maze)))
    if side_length * side_length != len(maze):
        raise InvalidMazeShapeException("Maze should be a square!")

    finish_postion = (side_length - 1, side_length - 1)  # Exit is always [N-1, N-1]

    start_pos = (0, 0)
    queue = deque(((0, start_pos),))
    visited = {(start_pos,)}

    while queue:
        dist, current_pos = queue.popleft()
        if current_pos == finish_postion:
            return dist
        else:
            for next_pos in (
                (current_pos[0] - 1, current_pos[1]),
                (current_pos[0] + 1, current_pos[1]),
                (current_pos[0], current_pos[1] - 1),
                (current_pos[0], current_pos[1] + 1),
            ):
                if (
                    (0 <= next_pos[0] < side_length)
                    and (0 <= next_pos[1] < side_length)
                    and maze[next_pos[0] * side_length + next_pos[1]] == "."
                    and next_pos not in visited
                ):
                    queue.append((dist + 1, next_pos))
                    visited.add(next_pos)
    return False


@dataclass
class PathFinderTestCase:
    name: str
    maze: str
    expected: int | bool


PATH_FINDER_TEST_CASES = [
    PathFinderTestCase(
        name="3x3 maze with a solution",
        maze="\n".join([".W.", ".W.", "..."]),
        expected=4,
    ),
    PathFinderTestCase(
        name="3x3, no solution",
        maze="\n".join([".W.", ".W.", "W.."]),
        expected=False,
    ),
    PathFinderTestCase(
        name="6x6 with a solution",
        maze="\n".join(["......", "......", "......", "......", "......", "......"]),
        expected=10,
    ),
    PathFinderTestCase(
        name="6x6, no solution",
        maze="\n".join(["......", "......", "......", "......", ".....W", "....W."]),
        expected=False,
    ),
]


@pytest.mark.parametrize("test_case", PATH_FINDER_TEST_CASES)
def test_path_finder(test_case: PathFinderTestCase):
    result = path_finder(test_case.maze)

    assert result == test_case.expected, f"{test_case.name}: result differs"

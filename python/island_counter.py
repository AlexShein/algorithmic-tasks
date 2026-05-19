import pytest
from dataclasses import dataclass
from collections import deque


def count_islands(image: list[list[int]]):
    """
    An "island" is a set of adjacent pixels of one color (1)
    which is surrounded by pixels of a different color (0).
    Pixels are considered adjacent if their coordinates differ by no more than 1 on the X or Y axis.
    Uses BFS approach. Time complexity is O(N)
    """

    island_count = 0

    start = (0, 0)
    length, height = len(image[0]), len(image)

    queue = deque(((False, start),))
    visited = set()

    while queue:
        came_from_island, position = queue.popleft()
        if position in visited:
            continue
        visited.add(position)
        if not came_from_island and image[position[0]][position[1]] == 1:
            island_count += 1
        if image[position[0]][position[1]] == 1:
            came_from_island = True
        else:
            came_from_island = False

        for dy in range(-1, 2):
            for dx in range(-1, 2):
                if (
                    not (dx == 0 and dy == 0)
                    and 0 <= (nex_row := position[0] + dy) < height
                    and 0 <= (nex_col := position[1] + dx) < length
                    and (next_pos := (nex_row, nex_col)) not in visited
                ):
                    if image[next_pos[0]][next_pos[1]] == 1:
                        queue.appendleft((came_from_island, next_pos))
                    else:
                        queue.append((came_from_island, next_pos))

    return island_count


@dataclass
class CountIslandsTestCase:
    name: str
    map: str
    expected: int


COUNT_ISLANDS_TEST_CASES = [
    CountIslandsTestCase(
        name="4x4 map with no islands",
        map=[[0, 0, 0, 0], [0, 0, 0, 0], [0, 0, 0, 0], [0, 0, 0, 0]],
        expected=0,
    ),
    CountIslandsTestCase(
        name="4x4 map with 1 island",
        map=[[0, 0, 0, 0], [0, 1, 1, 0], [0, 1, 1, 0], [0, 0, 0, 0]],
        expected=1,
    ),
    CountIslandsTestCase(
        name="4x4 map with 4 island",
        map=[[1, 1, 0, 1], [0, 0, 0, 1], [1, 0, 0, 0], [1, 0, 1, 1]],
        expected=4,
    ),
    CountIslandsTestCase(
        name="Larger map with 2 islands",
        map=[
            [0, 0, 0, 0, 1, 0, 0, 0, 0, 0],
            [0, 0, 1, 1, 1, 0, 0, 0, 0, 0],
            [0, 0, 1, 1, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 1, 0],
            [0, 0, 0, 0, 0, 1, 1, 1, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        ],
        expected=2,
    ),
]


@pytest.mark.parametrize("test_case", COUNT_ISLANDS_TEST_CASES)
def test_count_islands(test_case: CountIslandsTestCase):
    result = count_islands(test_case.map)

    assert result == test_case.expected, f"{test_case.name}: result differs"

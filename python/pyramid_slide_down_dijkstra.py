import pytest
from dataclasses import dataclass
import heapq


def longest_slide_down(pyramid: list[list[int]]) -> int:
    """
    Takes a pyramid representation as an argument and returns its largest 'slide down'.
    The 'slide down' is the maximum sum of consecutive numbers from the top to the bottom of the pyramid.
    It is assumed that pyramid numbers are always positive.
    The solution uses a pure Dijkstra algorithm
    """
    priority_queue = []
    max_score = 0
    # Best previous cell to get to the cell
    previous_points_map: dict[tuple[int, int], tuple[int, int]] = {}
    # Best score path to the cell, not including the cell itself
    score_map: dict[tuple[int, int], int] = {}

    heapq.heappush(priority_queue, (0, (0, 0)))

    while priority_queue:
        neg_score, point = heapq.heappop(priority_queue)
        current_score = -neg_score + pyramid[point[0]][point[1]]

        if point[0] == len(pyramid) - 1:
            if current_score > max_score:
                max_score = current_score
            continue  # We've reached the bottom, no points below this one.

        for adj_point in ((point[0] + 1, point[1]), (point[0] + 1, point[1] + 1)):
            if current_score > score_map.get(
                adj_point, -1
            ):  # Unexplored cells score -1
                heapq.heappush(
                    priority_queue,
                    (
                        -current_score,
                        adj_point,
                    ),
                )
                previous_points_map[adj_point] = point
                score_map[adj_point] = current_score

    return max_score


@dataclass
class LongestSlideDownTestCase:
    name: str
    pyramid: list[list[int]]
    expected_length: int


SLIDE_TEST_CASES = [
    LongestSlideDownTestCase(
        name="Trivial case: single level",
        pyramid=[[3]],
        expected_length=3,
    ),
    LongestSlideDownTestCase(
        name="Four levels",
        pyramid=[[3], [7, 4], [2, 4, 6], [8, 5, 9, 3]],
        expected_length=23,
    ),
    LongestSlideDownTestCase(
        name="Four levels 2",
        pyramid=[[10], [10, 20], [10, 10, 20], [10, 90, 10, 20]],
        expected_length=130,
    ),
    LongestSlideDownTestCase(
        name="Four levels, starts with 0",
        pyramid=[[0], [10, 20], [10, 10, 20], [10, 90, 10, 20]],
        expected_length=120,
    ),
]


@pytest.mark.parametrize("test_case", SLIDE_TEST_CASES)
def test_longest_slide_down(test_case: LongestSlideDownTestCase):
    length = longest_slide_down(test_case.pyramid)

    assert length == test_case.expected_length, f"{test_case.name}: length differs"

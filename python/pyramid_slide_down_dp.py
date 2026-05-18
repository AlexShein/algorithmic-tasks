import pytest
from dataclasses import dataclass


def longest_slide_down(pyramid: list[list[int]]) -> int:
    """
    Takes a pyramid representation as an argument and returns its largest 'slide down'.
    The 'slide down' is the maximum sum of consecutive numbers from the top to the bottom of the pyramid.
    It is assumed that pyramid numbers are always positive.
    The solution uses the dynamic programming approach calculating the max slide from the bottom
    """
    prev_row = pyramid.pop()
    while pyramid:
        curr_row = pyramid.pop()
        for i, _ in enumerate(curr_row):
            prev_row[i] = max(curr_row[i] + prev_row[i], curr_row[i] + prev_row[i + 1])
    return prev_row[0]


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

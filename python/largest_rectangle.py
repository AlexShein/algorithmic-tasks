import pytest
from dataclasses import dataclass


def largest_rect(histogram: list[int]) -> int:
    """
    Given a list of the number whitish pixels on each line in the background,
    returns the area of the largest rectangle that fits on that background.
    Stack-based O(N) solution
    """
    max_area = 0
    stack = []
    for i, value in enumerate(
        histogram + [0]
    ):  # Add zero height column to force stack emptying
        while stack and value < histogram[stack[-1]]:
            current_pillar_index = stack.pop()
            left_boundary = stack[-1] if stack else -1
            right_boundary = i
            width = right_boundary - left_boundary - 1
            area = width * histogram[current_pillar_index]
            if area > max_area:
                max_area = area

        if not stack or value >= histogram[stack[-1]]:
            stack.append(i)

    return max_area


def largest_rect_naive(histogram: list[int]) -> int:
    """
    Given a list of the number whitish pixels on each line in the background,
    returns the area of the largest rectangle that fits on that background.
    Naive O(n^2) solution
    """
    max_area = 0
    if not histogram:
        return max_area

    for i, value in enumerate(histogram):
        temp_max_area = value
        n, m = i - 1, i + 1
        while n >= 0 and histogram[n] >= value:
            temp_max_area += value
            n -= 1
        while m < len(histogram) and histogram[m] >= value:
            temp_max_area += value
            m += 1
        if temp_max_area > max_area:
            max_area = temp_max_area
    return max_area


@dataclass
class LargestRectangleTestCase:
    name: str
    histogram: list[int]
    expected_largest_area: int  # number or -1 for impossible puzzles


TEST_CASES = [
    LargestRectangleTestCase(
        name="Trivial case: no items",
        histogram=[],
        expected_largest_area=0,
    ),
    LargestRectangleTestCase(
        name="Trivial case: one item",
        histogram=[1],
        expected_largest_area=1,
    ),
    LargestRectangleTestCase(
        name="Simple case: ones",
        histogram=[1, 1, 1],
        expected_largest_area=3,
    ),
    LargestRectangleTestCase(
        name="Simple case",
        histogram=[1, 2, 3],
        expected_largest_area=4,
    ),
    LargestRectangleTestCase(
        name="Simple case reversed",
        histogram=[3, 2, 1],
        expected_largest_area=4,
    ),
    LargestRectangleTestCase(
        name="Harder caase",
        histogram=[9, 7, 5, 4, 2, 5, 6, 7, 7, 5, 7, 6, 4, 4, 3, 2],
        expected_largest_area=36,
    ),
]


@pytest.mark.parametrize("test_case", TEST_CASES)
def test_largest_rect(test_case: LargestRectangleTestCase):
    area = largest_rect(test_case.histogram)

    assert area == test_case.expected_largest_area, (
        f"{test_case.name}: step number differs"
    )


if __name__ == "__main__":
    pass

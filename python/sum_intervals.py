import pytest
from dataclasses import dataclass


def sum_of_intervals(intervals: list[tuple[int, int]]) -> int:
    """
    https://www.codewars.com/kata/52b7ed099cdc285c300001cd
    Accepts an array of intervals, and returns the sum of all the interval lengths.
    Sliding the single position along intervals
    Adding to the length if we "jump" inside of an interval.
    """
    length_sum = 0
    current_position = float("-inf")
    for interval in sorted(intervals):
        if interval[0] > current_position:
            current_position = interval[0]
        if interval[1] > current_position:
            length_sum += interval[1] - current_position
            current_position = interval[1]
    return length_sum


def sum_of_intervals_first_impl(intervals: list[tuple[int, int]]) -> int:
    """
    https://www.codewars.com/kata/52b7ed099cdc285c300001cd
    Accepts an array of intervals, and returns the sum of all the interval lengths.
    """

    res_length = 0
    if not intervals:
        return res_length
    intervals = sorted(
        intervals, key=lambda interval: interval[0]
    )  # Sort intervals by start - complexity is n log n
    current = 0
    while True:
        interval = intervals[current]
        start = interval[0]
        finish = interval[1]
        for i in range(current, len(intervals)):  # sum((n-1), (n-2),..., 2, 1) ~ n/2
            next_interval = intervals[i]
            if next_interval[0] >= finish:
                res_length += (
                    finish - start
                )  # We found a cliff, finish prev interval and add it to the length
                current = i  # Switch to the next one
                start, finish = next_interval
                break
            elif next_interval[1] > finish:  # Greedily extending current interval
                finish = next_interval[1]
            current = i
        if current == len(intervals) - 1:  # We're at the last el
            res_length += finish - start  # Final sum
            break
    return res_length


@dataclass
class SumOfIntervalsTestCase:
    name: str
    intervals: list[tuple[int, int]]
    expected: bool


TEST_CASES = [
    SumOfIntervalsTestCase(
        name="Trivial case: single interval", intervals=[(1, 5)], expected=4
    ),
    SumOfIntervalsTestCase(
        name="Negative case: single interval", intervals=[(-5, -1)], expected=4
    ),
    SumOfIntervalsTestCase(
        name="Two intervals", intervals=[(1, 5), (6, 10)], expected=8
    ),
    SumOfIntervalsTestCase(
        name="Two intervals 2", intervals=[(7, 12), (1, 5)], expected=9
    ),
    SumOfIntervalsTestCase(
        name="Two intervals: negative", intervals=[(-5, -1), (-2, 2)], expected=7
    ),
    SumOfIntervalsTestCase(
        name="Complex case: multiple intervals",
        intervals=[(1, 5), (2, 6), (3, 8), (7, 10), (11, 13)],
        expected=11,
    ),
    SumOfIntervalsTestCase(
        name="Single interval: huge number",
        intervals=[(-1_000_000_000, 1_000_000_000)],
        expected=2_000_000_000,
    ),
    SumOfIntervalsTestCase(
        name="Three intervals: huge number",
        intervals=[(0, 20), (-100_000_000, 10), (30, 40)],
        expected=100_000_030,
    ),
]


@pytest.mark.parametrize("test_case", TEST_CASES)
def test_sum_of_intervals(test_case: SumOfIntervalsTestCase):
    res = sum_of_intervals(test_case.intervals)
    assert res == test_case.expected, f"{test_case.name}: result differs"

import pytest
from dataclasses import dataclass
from collections import deque
from math import gcd


def wpp(a: int, b: int, n: int) -> list[tuple[int, int]]:
    """
    https://en.wikipedia.org/wiki/Water_pouring_puzzle
    This kata's goal is to get one of jugs to have exactly n units of water inside
    Allowed actions: emptying / filling jugs, filling one from another, emptying one into another
    Solving using the breadth-first search approach
    """
    path = []
    if (n > a and n > b) or (n % gcd(a, b) != 0):
        return []

    visited_cells = {(0, 0)}
    prev: dict[tuple[int, int], tuple[int, int]] = {}
    queue = deque([(0, 0)])
    while queue:
        node = queue.popleft()

        if node[0] == n or node[1] == n:
            path.append(node)
            while node := prev.get(node):
                path.append(node)
            break
        possible_next_nodes = [
            # Emptying
            (0, node[1]),
            (node[0], 0),
            # Filling
            (a, node[1]),
            (node[0], b),
        ]
        # Empty one into another
        if (node[0] + node[1]) <= a:
            possible_next_nodes.append((node[0] + node[1], 0))
        if (node[0] + node[1]) <= b:
            possible_next_nodes.append((0, node[0] + node[1]))
        # Fill one from the another
        if (room_in_a := a - node[0]) > 0 and (left_in_b := node[1] - room_in_a) > 0:
            possible_next_nodes.append((a, left_in_b))
        if (room_in_b := b - node[1]) > 0 and (left_in_a := node[0] - room_in_b) > 0:
            possible_next_nodes.append((left_in_a, b))

        for new_node in possible_next_nodes:
            if new_node not in visited_cells:
                visited_cells.add(new_node)
                queue.append(new_node)
                prev[new_node] = node

    return list(reversed(path))[1:]  # Exclude the (0,0) start as per task's definition


@dataclass
class WaterPouringPuzzeTestCase:
    name: str
    a: int
    b: int
    n: int
    expected_step_number: int  # number or -1 for impossible puzzles


TEST_CASES = [
    WaterPouringPuzzeTestCase(
        name="Trivial case: target is larger than both jugs",
        a=1,
        b=2,
        n=3,
        expected_step_number=0,
    ),
    WaterPouringPuzzeTestCase(
        name="Trivial case: target is not a factor of GCD of both numbers",
        a=3,
        b=6,
        n=2,
        expected_step_number=0,
    ),
    WaterPouringPuzzeTestCase(
        name="3, 5, 4: 6 steps",
        a=3,
        b=5,
        n=4,
        expected_step_number=6,  # [(0,5), (3,2), (0,2), (2,0), (2,5), (3,4)]
    ),
    WaterPouringPuzzeTestCase(
        name="1 step",
        a=4,
        b=5,
        n=4,
        expected_step_number=1,  # [(4,0)]
    ),
    WaterPouringPuzzeTestCase(
        name="1 step",
        a=4,
        b=5,
        n=5,
        expected_step_number=1,  # [(0,5)]
    ),
]


@pytest.mark.parametrize("test_case", TEST_CASES)
def test_wpp(test_case: WaterPouringPuzzeTestCase):
    steps = wpp(test_case.a, test_case.b, test_case.n)

    assert len(steps) == test_case.expected_step_number, (
        f"{test_case.name}: step number differs"
    )


if __name__ == "__main__":
    pass

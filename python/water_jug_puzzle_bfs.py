import pytest
from dataclasses import dataclass
from collections import deque


def solve(capacities: list[int, int, int], goal: list[int, int, int]) -> int:
    """
    Given a set of 3 jugs of water that have capacities of a, b, and c liters,
    find the minimum number of operations performed before each jug has x, y, and z liters.
    *Important*: Only jug C will start completely filled.
    Following operations allowed:
    Water is poured from one jug to another until one of the jugs is either empty or full.

    https://en.wikipedia.org/wiki/Water_pouring_puzzle

    Function returns the minimum number of operations needed to reach the goal state.
    If there is no possible solution, returns -1.

    Naive BFS solution
    """
    capacities, goal = tuple(capacities), tuple(goal)

    if (
        sum(goal) > capacities[2]  # Not enough water in the jug C.
        or any(capacities[i] < goal[i] for i in range(3))
    ):
        return -1
    initial_node = (0, 0, capacities[2])  # Jug C is filled by default
    visited_cells = {initial_node}
    queue = deque(((0, initial_node),))
    while queue:
        step_count, node = queue.popleft()
        if node == goal:
            return step_count
        possible_next_nodes = []
        for i in range(3):
            for j in range(3):
                if i != j:  # We don't pour from current jug to the other
                    amount_to_pour = min(node[i], capacities[j] - node[j])
                    if amount_to_pour:
                        temp_node = list(node)
                        temp_node[i] -= amount_to_pour
                        temp_node[j] += amount_to_pour
                        possible_next_nodes.append(tuple(temp_node))
        for new_node in possible_next_nodes:
            if new_node not in visited_cells:
                visited_cells.add(new_node)
                queue.append((step_count + 1, new_node))
    return -1


@dataclass
class WaterPouringPuzzeTestCase:
    name: str
    capacities: list[int, int, int]
    goal: list[int, int, int]
    expected_step_number: int  # number or -1 for impossible puzzles


TEST_CASES = [
    WaterPouringPuzzeTestCase(
        name="Trivial case: target is larger than one of jugs",
        capacities=[3, 5, 8],
        goal=[4, 0, 4],
        expected_step_number=-1,
    ),
    WaterPouringPuzzeTestCase(
        name="Trivial case: Not enough water is jug C",
        capacities=[3, 5, 8],
        goal=[3, 5, 1],
        expected_step_number=-1,
    ),
    WaterPouringPuzzeTestCase(
        name="One step solution",
        capacities=[3, 5, 8],
        goal=[0, 5, 3],
        expected_step_number=1,
    ),
    WaterPouringPuzzeTestCase(
        name="No solution exists",
        capacities=[4, 17, 22],
        goal=[2, 5, 15],
        expected_step_number=-1,
    ),
    WaterPouringPuzzeTestCase(
        name="Tough puzzle",
        capacities=[4, 7, 10],
        goal=[0, 5, 5],
        expected_step_number=8,
    ),
]


@pytest.mark.parametrize("test_case", TEST_CASES)
def test_solve(test_case: WaterPouringPuzzeTestCase):
    step_count = solve(test_case.capacities, test_case.goal)

    assert step_count == test_case.expected_step_number, (
        f"{test_case.name}: step number differs"
    )


if __name__ == "__main__":
    pass

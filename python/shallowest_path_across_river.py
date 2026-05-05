import pytest
from dataclasses import dataclass
import heapq
from pprint import pprint

HUGE_NUMBER = int(10e10)  # arbitrary initial large number


class Solution:
    path: list[tuple[int]]
    length: int
    depth: int

    def __init__(
        self, point: tuple[int, int], prev: dict[tuple[int, int], int], depth: int
    ):
        self.depth = depth
        self.path = point_path(point, prev)
        self.length = len(self.path)


def point_path(
    point: tuple[int, int], prev: dict[tuple[int, int], int]
) -> list[tuple[int]]:
    """Returns path to the current point including the point itself by traversing prev references"""
    current_point = point
    path = [point]
    while current_point := prev.get(current_point):
        path.append(current_point)
    return list(reversed(path))  # We started traversing from the target


def adjacent_points(i: int, j: int, dim: tuple[int, int]):
    """
    Returns adjacent points
    Going back to the left bank is disallowed
    """

    for n in range(-1, 2):
        for m in range(-1, 2):
            if (n, m) == (0, 0) or (j + m) == 0:
                continue  # Ignore the current point or the left bank
            if 0 <= i + n < dim[0] and 0 <= j + m < dim[1]:
                # print("Adj points returning", (i + n, j + m))
                yield (i + n, j + m)
    return  # Default case when no point could be yielded


def shallowest_path(river: list[list[int]]) -> list[tuple[int, int]]:
    """
    shallowest_path takes a list of lists of positive ints showing the depths of the river
    and returns a shallowest path (i.e., the maximum depth is minimal) as a list of coordinate pairs.
    If there are several paths that are equally shallow, the function returns a shortest such path.
    All depths are given as positive integers.
    """
    print("Solving for: ")
    pprint(river)
    # The priority queue: a simple list of tuples (distance, node)
    # heapq maintains the min-heap invariant on this list
    priority_queue = []

    best_solution: Solution | None = None

    # Best previous cell to get to the cell
    previous_points_map: dict[tuple[int, int], tuple[int, int]] = {}
    # Best path depth to the cell, not including the cell itself
    depth_map: dict[tuple[int, int], int] = {}

    dim = (len(river), len(river[0]))

    for i in range(dim[0]):
        for j in range(dim[1]):
            depth_map[(i, j)] = HUGE_NUMBER

    # Put left bank in the queue:
    for i in range(dim[0]):
        heapq.heappush(priority_queue, (river[i][0], (i, 0)))

    while priority_queue:
        depth, point = heapq.heappop(priority_queue)
        print(f"Picked point, processing {depth=} {point=}")

        if best_solution and depth > best_solution.depth:
            break  # The best remaining point is deeper than the best know solution

        if point[1] == dim[1] - 1:  # We reached the right coast.
            solution = Solution(point, previous_points_map, depth)
            if not best_solution or (
                depth < best_solution.depth
                or (
                    depth == best_solution.depth
                    and solution.length < best_solution.length
                )
            ):
                best_solution = solution
                continue

        for adj_point in adjacent_points(point[0], point[1], dim):
            max_depth = max(depth, river[adj_point[0]][adj_point[1]])
            if (
                max_depth < depth_map[adj_point]
                or (  # The depth is lower
                    max_depth
                    == depth_map[
                        adj_point
                    ]  # Depth is the same, but path via the current point is shorter than the previously discovered one
                    and len(point_path(point, previous_points_map)) + 1
                    < len(point_path(adj_point, previous_points_map))
                )
            ):
                print(f"Pushing {max_depth=}, {adj_point=}")
                heapq.heappush(
                    priority_queue,
                    (
                        max_depth,
                        adj_point,
                    ),
                )
                previous_points_map[adj_point] = point
                depth_map[adj_point] = max_depth

    return best_solution.path


@dataclass
class AdjacentPointsTestCase:
    name: str
    i: int
    j: int
    dim: tuple[int, int]
    expected_adjacent_points: list[int]


ADJACENT_POINTS_TEST_CASES = [
    AdjacentPointsTestCase(
        name="Single row",
        i=1,
        j=0,
        dim=(5, 0),
        expected_adjacent_points=[],
    ),
    AdjacentPointsTestCase(
        name="Left top corner",
        i=0,
        j=0,
        dim=(10, 10),
        expected_adjacent_points=[(0, 1), (1, 1)],
    ),
    AdjacentPointsTestCase(
        name="Close to the left bank",
        i=1,
        j=1,
        dim=(10, 10),
        expected_adjacent_points=[(0, 1), (0, 2), (1, 2), (2, 1), (2, 2)],
    ),
    AdjacentPointsTestCase(
        name="All points around are OK",
        i=2,
        j=2,
        dim=(10, 10),
        expected_adjacent_points=[
            (1, 1),
            (1, 2),
            (1, 3),
            (2, 1),
            (2, 3),
            (3, 1),
            (3, 2),
            (3, 3),
        ],
    ),
    AdjacentPointsTestCase(
        name="Right bottom corner",
        i=3,
        j=2,
        dim=(4, 4),
        expected_adjacent_points=[(2, 1), (2, 2), (2, 3), (3, 1), (3, 3)],
    ),
]


@pytest.mark.parametrize("test_case", ADJACENT_POINTS_TEST_CASES)
def test_adjacent_points(test_case: AdjacentPointsTestCase):
    assert (
        list(adjacent_points(test_case.i, test_case.j, test_case.dim))
        == test_case.expected_adjacent_points
    ), test_case.name


@dataclass
class ShallowestDepthTestCase:
    name: str
    river: list[list[int]]
    expected_depth: int
    expected_length: int


DEPTH_TEST_CASES = [
    ShallowestDepthTestCase(
        name="Single column",
        river=[[8], [8], [8], [8], [8], [8], [5], [8], [8], [8], [8], [8]],
        expected_depth=5,
        expected_length=1,
    ),
    ShallowestDepthTestCase(
        name="Two columns",
        river=[[1, 8], [8, 8], [8, 8], [8, 1], [8, 8], [8, 8], [1, 8], [8, 1], [8, 8]],
        expected_depth=1,
        expected_length=2,
    ),
    ShallowestDepthTestCase(
        name="Proper test",
        river=[
            [1, 8, 8],
            [8, 8, 8],
            [8, 8, 1],
            [8, 8, 1],
            [8, 1, 8],
            [8, 8, 1],
            [1, 1, 8],
            [8, 8, 1],
            [8, 8, 8],
        ],
        expected_depth=1,
        expected_length=3,
    ),
    ShallowestDepthTestCase(
        name="Two paths",
        river=[
            [1, 8, 8],
            [8, 1, 8],
            [8, 1, 3],
            [8, 8, 2],
            [8, 8, 8],
            [8, 8, 1],
            [8, 1, 1],
            [8, 2, 8],
            [1, 8, 8],
        ],
        expected_depth=2,
        expected_length=3,
    ),
    ShallowestDepthTestCase(
        name="Two paths - wider",
        river=[
            [2, 2, 8, 8],
            [8, 5, 2, 5],
            [8, 3, 2, 8],
            [1, 8, 8, 3],
        ],
        expected_depth=3,
        expected_length=4,
    ),
    ShallowestDepthTestCase(
        name="Hard puzzle",
        river=[
            [1, 1, 1, 8, 8, 1, 8, 8, 8, 1],
            [1, 1, 1, 8, 1, 8, 8, 8, 8, 8],
            [1, 1, 8, 1, 8, 1, 8, 1, 8, 8],
            [1, 1, 8, 8, 1, 8, 8, 1, 1, 8],
            [8, 1, 8, 1, 1, 1, 8, 1, 8, 8],
            [8, 1, 8, 1, 8, 1, 8, 1, 1, 8],
            [1, 1, 1, 8, 8, 8, 1, 8, 1, 1],
            [1, 8, 8, 8, 8, 1, 8, 1, 1, 1],
            [8, 8, 8, 1, 1, 1, 8, 8, 8, 1],
            [1, 8, 8, 1, 8, 1, 8, 8, 8, 1],
        ],
        expected_depth=1,
        expected_length=10,
    ),
    ShallowestDepthTestCase(
        name="Hard puzzle 2",
        river=[
            [95, 66, 51, 86, 63, 54, 51, 78, 98, 33, 78, 35, 15, 71, 98],
            [70, 28, 67, 53, 24, 53, 43, 40, 48, 41, 84, 55, 56, 56, 66],
            [59, 70, 25, 54, 7, 60, 95, 19, 93, 40, 74, 92, 3, 57, 66],
            [47, 80, 72, 74, 61, 82, 71, 53, 47, 82, 23, 24, 54, 53, 45],
            [78, 6, 10, 62, 39, 85, 58, 27, 14, 89, 75, 91, 78, 22, 6],
            [85, 35, 73, 76, 81, 35, 84, 54, 55, 74, 18, 51, 27, 21, 39],
            [28, 19, 16, 92, 65, 6, 93, 9, 5, 87, 52, 35, 92, 33, 7],
            [48, 43, 47, 62, 31, 89, 43, 6, 66, 18, 54, 73, 26, 36, 63],
            [30, 84, 63, 34, 90, 13, 90, 3, 47, 21, 3, 7, 40, 68, 52],
            [15, 67, 82, 70, 16, 67, 6, 32, 28, 56, 4, 61, 81, 62, 91],
            [97, 33, 96, 39, 36, 29, 66, 17, 26, 52, 49, 24, 74, 11, 47],
            [85, 36, 25, 98, 73, 52, 27, 14, 84, 49, 55, 73, 22, 37, 5],
            [41, 43, 8, 1, 38, 94, 60, 20, 55, 88, 46, 11, 33, 83, 54],
            [47, 8, 86, 42, 49, 98, 25, 27, 31, 68, 70, 29, 7, 90, 4],
            [52, 81, 9, 72, 45, 90, 4, 83, 75, 25, 14, 37, 19, 11, 4],
        ],
        expected_depth=39,
        expected_length=20,
    ),
    ShallowestDepthTestCase(
        name="Hard puzzle 3",
        river=[
            [95, 66, 51, 86, 63, 54, 51, 78, 98, 33, 78, 35, 15, 71, 98],
            [70, 28, 67, 53, 24, 53, 43, 40, 48, 41, 84, 55, 56, 56, 66],
            [59, 70, 25, 54, 7, 60, 95, 19, 93, 40, 74, 92, 3, 57, 66],
            [47, 80, 72, 74, 61, 82, 71, 53, 47, 82, 23, 24, 54, 53, 45],
            [78, 6, 10, 62, 39, 85, 58, 27, 14, 89, 75, 91, 78, 22, 6],
            [85, 35, 73, 76, 81, 35, 84, 54, 55, 74, 18, 51, 27, 21, 39],
            [28, 19, 16, 92, 65, 6, 93, 9, 5, 87, 52, 35, 92, 33, 7],
            [48, 43, 47, 62, 31, 89, 43, 6, 66, 18, 54, 73, 26, 36, 63],
            [30, 84, 63, 34, 90, 13, 90, 3, 47, 21, 3, 7, 40, 68, 52],
            [15, 67, 82, 70, 16, 67, 6, 32, 28, 56, 4, 61, 81, 62, 91],
            [97, 33, 96, 39, 36, 29, 66, 17, 26, 52, 49, 24, 74, 11, 47],
            [85, 36, 25, 98, 73, 52, 27, 14, 84, 49, 55, 73, 22, 37, 5],
            [41, 43, 8, 1, 38, 94, 60, 20, 55, 88, 46, 11, 33, 83, 54],
            [47, 8, 86, 42, 49, 98, 25, 27, 31, 68, 70, 29, 7, 90, 4],
            [52, 81, 9, 72, 45, 90, 4, 83, 75, 25, 14, 37, 19, 11, 4],
        ],
        expected_depth=39,
        expected_length=20,
    ),
    ShallowestDepthTestCase(
        name="Hard puzzle 4",
        river=[
            # 0. 1.  2.   3.  4. 5.  6.  7.  8.  9. 10 11. 12. 13. 14
            [93, 31, 32, 43, 8, 34, 16, 4, 78, 39, 65, 7, 35, 46, 24],  # 0
            [4, 97, 28, 25, 26, 20, 80, 95, 93, 16, 81, 79, 61, 58, 67],  # 1
            [68, 10, 95, 22, 97, 33, 76, 69, 74, 73, 9, 12, 32, 6, 19],  # 2
            [59, 63, 16, 47, 32, 83, 1, 11, 37, 22, 17, 15, 62, 61, 14],  # 3
            [55, 88, 82, 48, 99, 85, 5, 69, 19, 11, 6, 31, 8, 24, 31],  # 4
        ],
        expected_depth=33,
        expected_length=15,
    ),
]


@pytest.mark.parametrize("test_case", DEPTH_TEST_CASES)
def test_shallowest_path(test_case: ShallowestDepthTestCase):
    path = shallowest_path(test_case.river)

    assert path[0][1] == 0, "Path does not start on left bank."
    assert path[-1][1] == len(test_case.river[0]) - 1, (
        "Path does not end on right bank."
    )
    assert all(
        abs(r - path[j][0]) <= 1 and abs(c - path[j][1]) <= 1
        for j, (r, c) in enumerate(path[1:])
    ), "Path is not continuous."
    assert max(test_case.river[r][c] for (r, c) in path) <= test_case.expected_depth, (
        "Path is deeper than expected. Expected depth is %d." % test_case.expected_depth
    )
    assert len(path) <= test_case.expected_length, (
        "Path is longer than expected. Expected length is %d."
        % test_case.expected_length
    )


if __name__ == "__main__":
    print(
        shallowest_path(
            [
                [2, 2, 8, 8],
                [8, 5, 2, 5],
                [8, 3, 2, 8],
                [1, 8, 8, 3],
            ],
        )
    )

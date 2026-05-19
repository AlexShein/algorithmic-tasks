import pytest
from dataclasses import dataclass
from collections import defaultdict
import heapq


def shortest(vertex_count: int, edge_list: list[list[int]]):
    """Uses the Dijkstra algorithm"""
    graph = defaultdict(list)
    for edge in edge_list:
        # Each vertex points to a pair - weight, vertex
        graph[edge[0]].append((edge[2], edge[1]))

    start = 0  # We always start from the vertex number 0
    end = vertex_count - 1  # We always finish at the last vertex

    priority_queue: list[tuple[int, int]] = [(0, start)]  # weight, vertex
    previous_vertices_map: dict[int, int] = {}
    distance_map: dict[int, int] = {}

    while priority_queue:
        weight, current_vertex = heapq.heappop(priority_queue)
        if current_vertex == end:
            return weight

        for edge_weight, adj_vertex in graph[current_vertex]:
            if (
                not previous_vertices_map.get(adj_vertex)
                or distance_map[adj_vertex] > weight + edge_weight
            ):
                adj_vertex_weight = weight + edge_weight
                previous_vertices_map[adj_vertex] = current_vertex
                distance_map[adj_vertex] = adj_vertex_weight
                heapq.heappush(priority_queue, (adj_vertex_weight, adj_vertex))
    return -1


@dataclass
class ShortestPathTestCase:
    name: str
    vertex_count: int
    graph: list[list[int]]
    expected_length: int


PATH_TEST_CASES = [
    ShortestPathTestCase(
        name="Path exists",
        vertex_count=4,
        graph=[[0, 1, 1], [0, 2, 5], [0, 3, 5], [1, 3, 3], [2, 3, 1]],
        expected_length=4,
    ),
    ShortestPathTestCase(
        name="Path exists, 3 vertices",
        vertex_count=3,
        graph=[[0, 1, 7], [1, 2, 5], [0, 2, 12]],
        expected_length=12,
    ),
    ShortestPathTestCase(
        name="Path does not exist",
        vertex_count=5,
        graph=[[0, 2, 1], [1, 2, 1], [0, 3, 1], [1, 4, 1]],
        expected_length=-1,
    ),
]


@pytest.mark.parametrize("test_case", PATH_TEST_CASES)
def test_shortest(test_case: ShortestPathTestCase):
    length = shortest(test_case.vertex_count, test_case.graph)

    assert length == test_case.expected_length, f"{test_case.name}: length differs"

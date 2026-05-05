import math
from typing import Tuple, List
from collections import namedtuple


def merged_sort(input_list: List[int]) -> Tuple[List[int], int]:
    if len(input_list) > 2:
        mid = math.ceil(len(input_list) / 2)
        (sorted_part1, i_count_1), (sorted_part2, i_count_2) = merged_sort(input_list[:mid]), merged_sort(
            input_list[mid:]
        )

        i_count = i_count_1 + i_count_2
        i, j = 0, 0
        res = []

        while i + j < len(input_list):
            i_count += 1
            if i < len(sorted_part1) and j < len(sorted_part2):
                if sorted_part1[i] < sorted_part2[j]:
                    res.append(sorted_part1[i])
                    i += 1
                else:
                    res.append(sorted_part2[j])
                    j += 1
            elif i < len(sorted_part1):
                res.append(sorted_part1[i])
                i += 1
            else:
                res.append(sorted_part2[j])
                j += 1
        return res, i_count

    elif len(input_list) == 2:
        return [min(input_list), max(input_list)], 2

    return input_list, 1


TestCase = namedtuple("TestCase", ["name", "input", "output"])

TEST_CASES = [
    TestCase(
        name="Just a list",
        input=[3, 87, 34, 12, 89, 2, 5, 2, 47, 99, 6, 10, 3, 12, 58, 94, 24, 23, 17],
        output=sorted([3, 87, 34, 12, 89, 2, 5, 2, 47, 99, 6, 10, 3, 12, 58, 94, 24, 23, 17]),
    ),
    TestCase(
        name="Just a list",
        input=[4, 76, 23, 5, 23, 72, 19, 24, 10, 8, 2, 7, 9],
        output=sorted([4, 76, 23, 5, 23, 72, 19, 24, 10, 8, 2, 7, 9]),
    ),
    TestCase(
        name="Just a list",
        input=[3, 4, 1, 2, 6, 8, 7, 5],
        output=sorted([3, 4, 1, 2, 6, 8, 7, 5]),
    ),
    TestCase(
        name="Small list",
        input=[4, 3],
        output=sorted([3, 4]),
    ),
    TestCase(
        name="Empty list",
        input=[],
        output=sorted([]),
    ),
    TestCase(
        name="One-element list",
        input=[1],
        output=sorted([1]),
    ),
]

if __name__ == "__main__":
    for t_case in TEST_CASES:
        res, iter_count = merged_sort(t_case.input)
        assert res == t_case.output
        print(
            f"###\nSuccess:{res == t_case.output}\nResult  :{res}\nExpected:{t_case.output}\nIter count:{iter_count}\nInput len:{len(t_case.input)}",
        )

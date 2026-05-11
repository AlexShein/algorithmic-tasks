import pytest
from dataclasses import dataclass
from typing import Optional


@dataclass
class Node:
    value: int
    left: Optional["Node"] = None
    right: Optional["Node"] = None


def in_order(node: Node | None):
    result = []
    stack = [(False, node)]
    while stack:
        visited, node = stack.pop()
        if not node:
            continue
        if visited:
            yield node.value
        else:
            stack.append((False, node.right))
            stack.append((True, node))
            stack.append((False, node.left))
    return result


# Using the in_order generator to get next elements
def is_bst(node: Node | None) -> bool:
    order = None  # Order is unknown
    if node:
        in_order_generator = in_order(node)
        current_el = next(in_order_generator)
        for el in in_order_generator:
            print(f"Iteration {current_el=}, {el=} {order=}")
            if not order:  # order is not predefined and is inferred from the first pair
                order = -1 if el < current_el else 1
            if (order == -1 and el > current_el) or (order == 1 and el < current_el):
                return False  # Items order is inconsistent
            current_el = el
    return True


@dataclass
class BinarySearchTreeValidationTestCase:
    name: str
    node: Node | None
    expected: bool


n2 = Node(2)
n9 = Node(9)
n7 = Node(7, left=n9, right=n2)

n2 = Node(2)
n3 = Node(3)
n1 = Node(1, left=n2, right=n3)

Node(5, Node(2, Node(1), Node(3)), Node(7, None, Node(9)))

n08 = Node(7, right=Node(9))
n02 = Node(2, left=Node(1), right=Node(3))
n05 = Node(5, left=n02, right=n08)


TEST_CASES = [
    BinarySearchTreeValidationTestCase(
        name="Trivial case: None", node=None, expected=True
    ),
    BinarySearchTreeValidationTestCase(
        name="Base case, 1 item", node=n9, expected=True
    ),
    BinarySearchTreeValidationTestCase(
        name="Larger case, valid tree", node=n7, expected=True
    ),
    BinarySearchTreeValidationTestCase(
        name="Other case, valid tree", node=n05, expected=True
    ),
    BinarySearchTreeValidationTestCase(
        name="Base case, invalid tree", node=n1, expected=False
    ),
]


@pytest.mark.parametrize("test_case", TEST_CASES)
def test_is_bst(test_case: BinarySearchTreeValidationTestCase):
    res = is_bst(test_case.node)
    assert res == test_case.expected, f"{test_case.name}: result differs"

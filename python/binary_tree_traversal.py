import pytest
from dataclasses import dataclass
from typing import Optional


@dataclass
class Node:
    data: float | str  # A number or a string
    left: Optional["Node"] = None  # A Node, which is None if there is no left child.
    right: Optional["Node"] = None  # A Node, which is None if there is no right child.


# Pre-order traversal: iterative approach using a stack
def pre_order(node: Node | None):
    result = []
    stack = []
    while node:
        result.append(node.data)
        if node.right:
            stack.append(node.right)
        if node.left:
            node = node.left
        elif stack:
            node = stack.pop()
        else:
            node = None
    return result


# In-order traversal: iterative approach using a stack
def in_order(node: Node | None):
    result = []
    stack = [(False, node)]
    while stack:
        visited, node = stack.pop()
        if not node:
            continue
        if visited:
            result.append(node.data)
        else:
            stack.append((False, node.right))
            stack.append((True, node))
            stack.append((False, node.left))
    return result


# Post-order traversal: iterative approach using a stack
def post_order(node: Node | None):
    result = []
    stack = [(False, node)]
    while stack:
        visited, node = stack.pop()
        if not node:
            continue
        if visited:
            result.append(node.data)
        else:
            stack.append((True, node))
            stack.append((False, node.right))
            stack.append((False, node.left))
    return result


# Pre-order traversal: iterative approach using a stack
def pre_order_recursive(node: Node | None):
    return (
        [node.data, *pre_order_recursive(node.left), *pre_order_recursive(node.right)]
        if node
        else []
    )


# In-order traversal: naive recursive approach
def in_order_recursive(node: Node | None):
    return (
        [*in_order_recursive(node.left), node.data, *in_order_recursive(node.right)]
        if node
        else []
    )


# In-order traversal: naive recursive approach
def post_order_recursive(node: Node | None):
    return (
        [*post_order_recursive(node.left), *post_order_recursive(node.right), node.data]
        if node
        else []
    )


@dataclass
class BTreeTraversalTestCase:
    name: str
    node: Node | None
    expected: dict[str, int]


d = Node("leaf")
c = Node(2, left=d)
b = Node(1)
a = Node(5, left=b, right=c)

n2 = Node(2)
n3 = Node(3)
n4 = Node(4, left=n2, right=n3)
n8 = Node(8)
n23 = Node(23)
n9 = Node(9, left=n8, right=n23)
n6 = Node(6, left=n4, right=n9)


PRE_TEST_CASES = [
    BTreeTraversalTestCase(name="Trivial case: None", node=None, expected=[]),
    BTreeTraversalTestCase(name="Base case, 1 item", node=b, expected=[b.data]),
    BTreeTraversalTestCase(
        name="Base case, 2 items", node=c, expected=[c.data, d.data]
    ),
    BTreeTraversalTestCase(
        name="Base case, 4 items", node=a, expected=[a.data, b.data, c.data, d.data]
    ),
    BTreeTraversalTestCase(
        name="Larger case",
        node=n6,
        expected=[n6.data, n4.data, n2.data, n3.data, n9.data, n8.data, n23.data],
    ),
]

IN_TEST_CASES = [
    BTreeTraversalTestCase(name="Trivial case: None", node=None, expected=[]),
    BTreeTraversalTestCase(name="Base case, 1 item", node=b, expected=[b.data]),
    BTreeTraversalTestCase(
        name="Base case, 2 items", node=c, expected=[d.data, c.data]
    ),
    BTreeTraversalTestCase(
        name="Base case, 4 items", node=a, expected=[b.data, a.data, d.data, c.data]
    ),
    BTreeTraversalTestCase(
        name="Larger case",
        node=n6,
        expected=[n2.data, n4.data, n3.data, n6.data, n8.data, n9.data, n23.data],
    ),
]

POST_TEST_CASES = [
    BTreeTraversalTestCase(name="Trivial case: None", node=None, expected=[]),
    BTreeTraversalTestCase(name="Base case, 1 item", node=b, expected=[b.data]),
    BTreeTraversalTestCase(
        name="Base case, 2 items", node=c, expected=[d.data, c.data]
    ),
    BTreeTraversalTestCase(
        name="Base case, 4 items", node=a, expected=[b.data, d.data, c.data, a.data]
    ),
    BTreeTraversalTestCase(
        name="Larger case",
        node=n6,
        expected=[n2.data, n3.data, n4.data, n8.data, n23.data, n9.data, n6.data],
    ),
]


@pytest.mark.parametrize("test_case", PRE_TEST_CASES)
def test_pre_order(test_case: BTreeTraversalTestCase):
    res = pre_order(test_case.node)
    assert res == test_case.expected, f"{test_case.name}: result differs"


@pytest.mark.parametrize("test_case", IN_TEST_CASES)
def test_in_order(test_case: BTreeTraversalTestCase):
    res = in_order(test_case.node)
    assert res == test_case.expected, f"{test_case.name}: result differs"


@pytest.mark.parametrize("test_case", POST_TEST_CASES)
def test_post_order(test_case: BTreeTraversalTestCase):
    res = post_order(test_case.node)
    assert res == test_case.expected, f"{test_case.name}: result differs"

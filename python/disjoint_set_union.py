class DisjointSetUnion:
    parents: dict[int, int]
    sizes: dict[int, int]
    number_of_components: int

    def __init__(self, items: list[int]):
        self.parents = {}
        self.sizes = {}
        for i in items:
            self.parents[i] = i
            self.sizes[i] = 1
        self.number_of_components = len(items)

    def find(self, node: int) -> int:
        # Finds the parent of the node, compresses the path

        if self.parents[node] != node:
            self.parents[node] = self.find(self.parents[node])

        return self.parents[node]

    def union(self, node1: int, node2: int) -> bool:
        root1 = self.find(node1)
        root2 = self.find(node2)

        if root1 == root2:
            return False

        if self.sizes[root1] < self.sizes[root2]:
            root1, root2 = root2, root1

        self.parents[root2] = root1

        self.sizes[root1] += self.sizes[root2]

        self.number_of_components -= 1
        return True

    def is_connected(self, node1: int, node2: int) -> bool:
        return self.find(node1) == self.find(node2)


if __name__ == "__main__":
    dsu = DisjointSetUnion([6, 5, 4, 3, 2, 1, 0])
    print(f"Created, {dsu.parents=} {dsu.number_of_components=}")

    print(f"{dsu.find(5)=}")

    dsu.union(6, 0)
    print(f"After union 6, 0, {dsu.parents=} {dsu.number_of_components=}")

    dsu.union(0, 1)
    print(f"After union 0, 1, {dsu.parents=} {dsu.number_of_components=}")
    print(f"{dsu.is_connected(1, 6)=}")

sample_test_size = 10
large_test_size = 1000


class DSU:
    def __init__(self, n):
        self.p = list(range(n))
        self.sz = [1] * n

    def find(self, x):
        if self.p[x] != x:
            self.p[x] = self.find(self.p[x])
        return self.p[x]

    def union(self, a, b):
        ra, rb = self.find(a), self.find(b)
        if ra == rb:
            return False
        if self.sz[ra] < self.sz[rb]:
            ra, rb = rb, ra
        self.p[rb] = ra
        self.sz[ra] += self.sz[rb]
        return True


def solve_day8_part1(input_):
    input_ = [line.split(",") for line in input_]
    distances = []

    for i, line in enumerate(input_):
        for j in range(i+1, len(input_)):
            dis = get_distance(line, input_[j])
            distances.append((dis, (i, j)))

    distances.sort()

    dsu = DSU(len(input_))
    for i in range(sample_test_size):
        _, (j, k) = distances[i]
        dsu.union(j, k)

    # get the largest 3 circles
    roots = [dsu.sz[i] for i in range(len(input_)) if dsu.find(i) == i]
    roots.sort(reverse=True)
    return roots[0] * roots[1] * roots[2]


def get_distance(line1, line2):
    x_dis = int(line1[0])-int(line2[0])
    y_dis = int(line1[1])-int(line2[1])
    z_dis = int(line1[2])-int(line2[2])

    return x_dis**2+y_dis**2+z_dis**2

import itertools


def choose_best_sum(t, k, ls):
    combs = list(itertools.combinations(ls, k))

    sums = [s for s in map(sum, combs) if s <= t]

    if sums == []:
        return None

    return max(sums)

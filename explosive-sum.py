def exp_sum(n):
    table = [1] + [0] * n
    for i in range(1, n + 1):
        for j in range(i, n + 1):
            table[j] += table[j - i]
    return table[n]

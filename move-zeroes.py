def move_zeros(lst):
    zeroes = 0

    for n in lst:
        if n == 0:
            zeroes += 1

    lst = [x for x in lst if x != 0]

    return lst + [0] * zeroes

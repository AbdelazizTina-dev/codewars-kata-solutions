def two_sum(numbers, target):
    seen = {}

    for i, n in enumerate(numbers):
        complement = target - n

        if complement in seen:
            return (i, seen[complement])
        else:
            seen[n] = i

    return (0, 0)

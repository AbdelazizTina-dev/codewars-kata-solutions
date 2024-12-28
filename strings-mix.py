def mix(s1, s2):
    s1_count = {}
    s2_count = {}

    for c in s1:
        if not c.islower():
            continue
        s1_count[c] = s1_count.get(c, 0) + 1

    for c in s2:
        if not c.islower():
            continue
        s2_count[c] = s2_count.get(c, 0) + 1

    s1_count = dict(sorted(s1_count.items(), key=lambda item: (-item[1], item[0])))
    s2_count = dict(sorted(s2_count.items(), key=lambda item: (-item[1], item[0])))

    res = {}

    for key, value in s1_count.items():
        if not key in s2_count:
            if value == 1:
                continue
            res[key * value] = 1
        else:
            if value == 1:
                continue

            if value < s2_count[key]:
                res[key * s2_count[key]] = 2
            elif value > s2_count[key]:
                res[key * value] = 1
            else:
                res[key * value] = 3

            del s2_count[key]

    if len(s2_count) != 0:
        for key, value in s2_count.items():
            if value == 1:
                continue

            res[key * value] = 2

    res = dict(sorted(res.items(), key=lambda item: (-len(item[0]), item[1], item[0])))

    final = "/".join(
        [f"{'=' if value == 3 else value}:{key}" for key, value in res.items()]
    )

    return final

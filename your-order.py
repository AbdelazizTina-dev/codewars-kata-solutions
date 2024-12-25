def order(sentence):
    if sentence == "":
        return ""

    nums = [int(c) for c in sentence if c.isnumeric()]

    words = [t for t in zip(nums, sentence.split())]

    s = sorted(words, key=lambda x: x[0])

    return " ".join([x[1] for x in s])

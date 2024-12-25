def encode(message, key):
    dict = {chr(c): c - 96 for c in range(97, 123)}

    res = []

    flat_key = [int(digit) for digit in str(key)]

    for i, c in enumerate(message):
        res.append(dict[c] + flat_key[i % len(flat_key)])

    return res

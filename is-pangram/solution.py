def is_pangram(st):
    alpha = {chr(i): i - 96 for i in range(97, 123)}

    for c in st:
        if c.lower() in alpha:
            alpha.pop(c.lower())

    return alpha == {}

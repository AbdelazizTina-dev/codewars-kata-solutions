def valid_braces(string):
    if len(string) % 2 != 0:
        return False

    opp = {"]": "[", ")": "(", "}": "{"}

    stack = []

    for c in string:
        if c in opp:
            if stack and stack[-1] == opp[c]:
                stack.pop()
            else:
                return False
        else:
            stack.append(c)

    return stack == []

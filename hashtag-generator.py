def generate_hashtag(s):
    if s == "":
        return False

    list = [s.capitalize() for s in s.split()]

    result = "#" + "".join(list)

    if len(result) > 140:
        return False

    return result

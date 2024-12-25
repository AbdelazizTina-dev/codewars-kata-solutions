def alphanumeric(password):
    if password == "":
        return False

    for c in password:
        if not c.isalpha() and not c.isnumeric():
            return False

    return True

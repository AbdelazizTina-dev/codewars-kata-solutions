def cakes(recipe, available):
    buf = []

    for ingredient, amount in recipe.items():
        if ingredient not in available:
            return 0

        buf.append(available[ingredient] // amount)

    return min(buf)

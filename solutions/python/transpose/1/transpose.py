from itertools import zip_longest

def transpose(text):
    if not text:
        return ""

    rows = text.splitlines()
    cols = list(zip_longest(*rows))

    res = []
    for col in cols:
        lst = list(col)
        while lst and lst[-1] is None:
            lst.pop()
        res.append("".join(c if c else " " for c in lst))

    return "\n".join(res)
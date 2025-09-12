import re
import operator as op

OPS = {
    "plus": op.add,
    "multiplied by": op.mul,
    "minus": op.sub,
    "divided by": op.truediv,
}
TOKEN = re.compile(r"""
    (?P<num>-?\d+) |
    (?P<op>plus|minus|multiplied\ by| divided\ by)
""", re.IGNORECASE | re.VERBOSE)

def answer(question):
    q = question.strip()
    if not q.lower().startswith("what is") or not q.endswith("?"):
        raise ValueError("syntax error")

    core = q[7:-1].strip()
    if len(core) == 0:
        raise ValueError("syntax error")

    tokens = []
    idx = 0
    for m in TOKEN.finditer(core):
        gap = core[idx:m.start()]
        if gap.strip():
            raise ValueError("syntax error")
        idx = m.end()

        if m.lastgroup == "num":
            tokens.append(int(m.group("num")))
        else:
            tokens.append(m.group("op").lower())

    if core[idx:].strip():
        raise ValueError("unknown operation")

    if not tokens or not isinstance(tokens[0], int):
        raise ValueError("syntax error")

    res = tokens[0]
    i = 1
    while i < len(tokens):
        if i + 1 >= len(tokens) or not isinstance(tokens[i + 1], int):
            raise ValueError("syntax error")
        op = tokens[i]
        if op not in OPS:
            raise ValueError("unknown operation")
        try:
            res = OPS[op](res, tokens[i + 1])
        except ZeroDivisionError:
            raise ValueError("syntax error")
        i += 2

    return int(res)
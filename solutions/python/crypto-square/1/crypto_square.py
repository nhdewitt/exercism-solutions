from string import punctuation
from math import ceil, sqrt

def cipher_text(plain_text):
    normalized = "".join(
        [t for t in plain_text.lower().replace(" ","") if t not in punctuation]
    )
    n = len(normalized)

    c = ceil(sqrt(n))
    if c == 0:
        return ""
        
    r = ceil(n / c)
    
    strings = []

    for i in range(c):
        strings.append(f"{normalized[i * c:(i * c) + c]}")

    if len(strings) == 2:
        return strings[0]

    if len(strings[-1]) == 0 and len(strings) > 1:
        strings.pop()

    if len(strings[-1]) < len(strings[0]):
        diff = len(strings[0]) - len(strings[-1])
        strings[-1] += f"{' ' * diff}"
    

    output = [
        "".join(string)
        for string in zip(*strings)
    ]

    return " ".join(output)
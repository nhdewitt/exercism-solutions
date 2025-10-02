from string import ascii_lowercase
from secrets import randbelow

class Cipher:
    alpha = ascii_lowercase
    
    def __init__(self, key=None):
        if not key:
            self.key = "".join([Cipher.alpha[randbelow(26)] for _ in range(100)])
        else:
            self.key = key
        self.rot = [int(ord(k) - ord('a')) for k in self.key]

        self.n = len(self.rot)

    def shift(self, k, offset):
        base = ord('a')
        pos = ord(k) - base
        new_pos = (pos + offset) % 26
        return chr(base + new_pos)

    def encode(self, text):
        res = []
        for i, _ in enumerate(text):
            res.append(self.shift(text[i], self.rot[i % self.n]))

        return "".join(res)
            
    def decode(self, text):
        res = []
        for i, _ in enumerate(text):
            res.append(self.shift(text[i], -self.rot[i % self.n]))

        return "".join(res)

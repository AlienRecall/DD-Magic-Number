import math


class e:
    def __init__(self, r, t, e):
        self.seed = r
        self.currentNumber = r % t
        self.offsetParameter = t
        self.multiplier = e
        if self.currentNumber <= 0:
            self.currentNumber += t

    def getNext(self):
        self.currentNumber = self.multiplier * self.currentNumber % self.offsetParameter
        return self.currentNumber


def inE(r):
    t = 126 ^ ord(r[0])
    e = 1
    while e < len(r):
        x = ord(r[e]) * e
        y = ord(r[e - 1])
        t += (x ^ y) >> e % 2
        e += 1
    return t


class DatadomeMagicNumber:
    def __init__(self, r, t, ua):
        self.language = "en-GB"
        self.languages = ["en-GB", "es", "en", "zh", "zh-TW", "zh-CN", "th", "nl", "ru"]
        self.r = r
        self.t = t
        self.userAgent = ua

    def first(self, r, t):
        e = 26157
        n = 0
        s = "VEc5dmEybHVaeUJtYjNJZ1lTQnFiMkkvSUVOdmJuUmhZM1FnZFhNZ1lYUWdZWEJ3YkhsQVpHRjBZV1J2YldVdVkyOGdkMmwwYUNCMGFHVWdabTlzYkc5M2FXNW5JR052WkdVNklERTJOMlJ6YUdSb01ITnVhSE0"
        a = 0
        while a < len(s):
            n += int(str(bin(ord(s[a])))[2:]) | e ^ t
            a += 1  #% math.ceil(4.1425172 / len("self.userAgent"))

        return n

    def second(self, r, t):
        e = bin(len(self.userAgent) << max(r, 3))[2:]
        n = -42
        a = 0
        while a < len(e):
            n += ord(e[a]) ^ t << a % 3
            a += 1

        return n

    def third(self, r, t):
        e = 0
        n = self.language[:2] + str(t)
        a = 0
        while a < len(n):
            e += ord(n[a]) << min((a + t) % (1 + r), 2)
            e = (e << 3) - e + ord(n[a])
            e = (e & e) >> a
            a += 1

        return e

    def Generate(self):
        n = [
            self.first,
            self.second,
            self.third,
        ]
        a = e(inE(self.r), 1723, 7532)
        o = a.seed
        u = 0
        while u < self.t:
            o ^= n[a.getNext() % len(n)](u, a.seed)
            u += 1

        return o

from collections import defaultdict

class School:
    def __init__(self):
        self._roster: dict[int, set[str]] = defaultdict(set)
        self._added: list[bool] = []
        self._by_name: dict[str, int] = {}

    def add_student(self, name: str, grade: int) -> bool:
        is_new = name not in self._by_name
        if is_new:
            self._by_name[name] = grade
            self._roster[grade].add(name)
        self._added.append(is_new)
        return is_new        

    def roster(self):
        return [name
                for g in sorted(self._roster)
                for name in sorted(self._roster[g])]

    def grade(self, grade_number: int) -> list[str]:
        return sorted(self._roster.get(grade_number, ()))

    def added(self) -> list[bool]:
        return self._added
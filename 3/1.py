from pathlib import Path

engine: list[list[str]] = []
sum = 0


def check_cases(word_start: int, word_end: int, line_index: int, engine: list[list[str]]) -> bool:
    # Above cases
    if line_index != 0:
        if word_start != 0 and engine[line_index - 1][word_start - 1] != ".":
            return True
        
        for i in range(word_start, word_end + 1):
            if engine[line_index - 1][i] != ".":
                return True

        if word_end != len(engine[0]) - 2 and engine[line_index - 1][word_end + 1] != ".":
            return True

    # Left Side
    if word_start != 0 and engine[line_index][word_start - 1] != ".":
        return True

    # Below Cases
    if line_index != len(engine) - 1:
        if word_start != 0 and engine[line_index + 1][word_start - 1] != ".":
            return True
        
        for i in range(word_start, word_end + 1):
            if engine[line_index + 1][i] != ".":
                return True

        if word_end != len(engine[0]) - 2 and engine[line_index + 1][word_end + 1] != ".":
            return True

    return False

if __name__ == "__main__":
    cal: list[int] = []
    txt = Path("C:/Users/Andrew/Desktop/advent/3/input.txt")
    with txt.open() as f:
        for line in f:
            engine.append([*line])

    for i, line in enumerate(engine):
        word = ""
        word_start = None
        for j, char in enumerate(line):
            if char.isdigit():
                if word_start is None:
                    word_start = j
                word += char
            else:
                if word_start is None:
                    continue
                if (char != "." and char != "\n") or ((char == "." or char == "\n") and check_cases(word_start, j - 1, i, engine)):
                    # print(word)
                    sum += int(word)
                # Reset
                word = ""
                word_start = None

    print(sum)

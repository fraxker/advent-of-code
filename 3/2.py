from pathlib import Path
from math import prod

engine: list[list[str]] = []
sum = 0

# Returns gear ratio
def check_gear(line_index: int, column_index: int, engine: list[list[str]]) -> int:
    line_len = len(engine[0]) - 1

    gears: list[str] = []
    
    # Above cases
    if line_index != 0:
        word = ""
        # Check Center
        if engine[line_index - 1][column_index].isdigit():
            word += engine[line_index - 1][column_index]

            # Check Left
            if column_index != 0:
                for i in range(column_index - 1, -1, -1):
                    if not engine[line_index - 1][i].isdigit():
                        break
                    word += engine[line_index - 1][i]
            
                word = word[::-1]

            # Check Right
            if column_index != line_len - 1:
                for i in range(column_index + 1, line_len - 1):
                    if not engine[line_index - 1][i].isdigit():
                        break
                    word += engine[line_index - 1][i]
            
            gears.append(int(word))

        else:
            if column_index != 0 and engine[line_index - 1][column_index - 1].isdigit():
                word = engine[line_index - 1][column_index - 1]
                for i in range(column_index - 2, -1, -1):
                    if not engine[line_index - 1][i].isdigit():
                        break
                    word += engine[line_index - 1][i]
                gears.append(int(word[::-1]))
            

            if column_index != line_len - 1 and engine[line_index - 1][column_index + 1].isdigit():
                word = engine[line_index - 1][column_index + 1]
                for i in range(column_index + 2, line_len - 1):
                    if not engine[line_index - 1][i].isdigit():
                        break
                    word += engine[line_index - 1][i]
                gears.append(int(word))

    # Left Side
    if column_index != 0 and engine[line_index][column_index - 1].isdigit():
        word = engine[line_index][column_index - 1]
        for i in range(column_index - 2, -1, -1):
            if not engine[line_index][i].isdigit():
                break
            word += engine[line_index][i]
        gears.append(int(word[::-1]))

    # Right Side
    if column_index != line_len - 1 and engine[line_index][column_index + 1].isdigit():
        word = engine[line_index][column_index + 1]
        for i in range(column_index + 2, line_len - 1):
            if not engine[line_index][i].isdigit():
                break
            word += engine[line_index][i]
            
        gears.append(int(word))

    # Below Cases
    if line_index != line_len:
        word = ""
        # Check Center
        if engine[line_index + 1][column_index].isdigit():
            word += engine[line_index + 1][column_index]

            # Check Left
            if column_index != 0:
                for i in range(column_index - 1, -1, -1):
                    if not engine[line_index + 1][i].isdigit():
                        break
                    word += engine[line_index + 1][i]
            
                word = word[::-1]

            # Check Right
            if column_index != line_len - 1:
                for i in range(column_index + 1, line_len - 1):
                    if not engine[line_index + 1][i].isdigit():
                        break
                    word += engine[line_index + 1][i]
            
            gears.append(int(word))

        else:
            if column_index != 0 and engine[line_index + 1][column_index - 1].isdigit():
                word = engine[line_index + 1][column_index - 1]
                for i in range(column_index - 2, -1, -1):
                    if not engine[line_index + 1][i].isdigit():
                        break
                    word += engine[line_index + 1][i]
                gears.append(int(word[::-1]))

            if column_index != line_len - 1 and engine[line_index + 1][column_index + 1].isdigit():
                word = engine[line_index + 1][column_index + 1]
                for i in range(column_index + 2, line_len):
                    if not engine[line_index + 1][i].isdigit():
                        break
                    word += engine[line_index + 1][i]
                gears.append(int(word))

    if len(gears) == 2:
        print(gears)
        return prod(gears)

    return None

if __name__ == "__main__":
    cal: list[int] = []
    txt = Path("C:/Users/Andrew/Desktop/advent/3/input.txt")
    with txt.open() as f:
        for line in f:
            engine.append([*line])

    for i, line in enumerate(engine):
        for j, char in enumerate(line):
            if char == "*":
                ratio = check_gear(i, j, engine)
                if ratio != None:
                    sum += ratio

    print(sum)

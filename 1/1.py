from pathlib import Path


if __name__ == "__main__":
    cal: list[int] = []
    txt = Path("input.txt")
    with txt.open() as f:
        for line in f:
            f_digit = ""
            l_digit = ""

            for _, c in enumerate(line):
                if c.isdigit():
                    f_digit = c
                    break

            if f_digit == "":
                print("Could not find digit")
                exit()

            for _, c in enumerate(line[::-1]):
                if c.isdigit():
                    l_digit = c
                    break

            if l_digit == "":
                print("Could not find digit")
                exit()

            word = f_digit + l_digit

            cal.append(int(word))

    
    print(sum(cal))
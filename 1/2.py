from pathlib import Path

numbers = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]

if __name__ == "__main__":
    cal: list[int] = []
    txt = Path("input.txt")
    with txt.open() as f:
        for line in f:
            f_digit = ""
            f_word = ""
            l_digit = ""
            l_word = ""
            k = None

            for n in numbers:
                temp = line.find(n)
                if temp != -1 and (k is None or temp < k):
                    k = temp
                    f_word = n
                    break

            for i, c in enumerate(line):
                if c.isdigit() and (k is None or i < k):
                    f_digit = c
                    break

            if f_digit == "" and f_word == "":
                print("Could not find first digit:", line)
                exit()

            k = None
            line_reverse = line[::-1]

            for n in numbers:
                temp = line_reverse.find(n[::-1])
                if temp != -1 and (k is None or temp < k):
                    k = temp
                    l_word = n
                    break

            for i, c in enumerate(line_reverse):
                if c.isdigit() and (k is None or i < k):
                    l_digit = c
                    break

            if l_digit == "" and l_word == "":
                print("Could not find second digit:", line)
                exit()

            if f_word != "" and f_digit == "":
                f_digit = str(numbers.index(f_word) + 1)

            if l_word != "" and l_digit == "":
                l_digit = str(numbers.index(l_word) + 1)

            word = f_digit + l_digit

            cal.append(int(word))

    print(sum(cal))
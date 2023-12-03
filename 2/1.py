from pathlib import Path

MAX = {
    "red": 12,
    "green": 13,
    "blue": 14
}

sum = 0

if __name__ == "__main__":
    txt = Path("input.txt")
    with txt.open() as f:
        for i, line in enumerate(f):
            game = line.split(":")[1]
            for pull in game.split(";"):
                valid = True
                for color in pull.split(","):
                    num, color = color.split()
                    if int(num) > MAX[color]:
                        valid = False
                        break
                if valid is False:
                    break
            if valid is True:
                print(i+1)
                sum += i + 1
        print(sum)

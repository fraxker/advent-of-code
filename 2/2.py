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
        for line in f:
            game = line.split(":")[1]
            max = {
                "red": None,
                "green": None,
                "blue": None
            }
            for pull in game.split(";"):
                for color in pull.split(","):
                    num, color = color.split()
                    num = int(num)
                    if max[color] is None or max[color] < num:
                        max[color] = num
            pow = 1 
            for n in max.values():
                pow *= n
            sum += pow
        print(sum)

import sys


def get_input(file):
    with open(file) as f:
        lines = f.readlines()[0][:-1].split(",")
    return lines
  

def main():
    skus = get_input("input.txt")
    total = 0
    # for sku_range in skus:
    #     bounds = sku_range.split("-")
    #     start = int(bounds[0])
    #     end = int(bounds[1])
    #     for sku in range(start, end + 1):
    #         sku = str(sku)
    #         if len(sku) % 2 != 0:
    #             continue
    #         length = len(sku) // 2
    #         if sku[:length] == sku[length:]:
    #             total += int(sku)
    # print(total)

    # 34512560282 too high

    for sku_range in skus:
        bounds = sku_range.split("-")
        start = int(bounds[0])
        end = int(bounds[1])
        for sku in range(start, end + 1):
            sku = str(sku)
            length = len(sku)
            if length == 1:
                continue
            # length of 2 or 3 could only be first digit repeated
            if length <= 3:
                if sku[0] * length == sku:
                    total += int(sku)
                continue
           
            # simple case of first digit repeating
            if sku[0] * length == sku:
                total += int(sku)
                continue

            # if the SKU can't be broken into 2s or 3s then it's valid regardless
            if length % 2 != 0 and length % 3 != 0:
                continue

            for i in range(2, (length // 2) + 1):
                if length % i != 0:
                    continue
                if sku[0:i] * (length // i) == sku:
                    total += int(sku)
                    break

    print(total)


if __name__ == '__main__':
    main()

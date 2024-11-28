EPS = 1e-6

def read_file(name_file):
    try:
        with open(name_file, "r") as f:
            table = [list(map(float, string.split())) for string in list(f)]
        return table

    except:
        print("Ошибка чтения файла!")
        return []

def read_data(size_table):
    try:
        n = int(input("Введите степень аппроксимирующего полинома - n: "))

        if (n <= 0):
            print("Степень палинома должна быть > 0!")
            return 1, 0, 0
        elif (n >= size_table):
            print("Слишком большая степень аппроксимирующего полинома для данной таблицы!")
            return 2, 0, 0

        x = float(input("Введите значение аргумента, для которого выполняется интерполяция: "))
        return 0, n, x

    except:
        print("Ошибка ввода данных!")
        return 3, 0, 0

def print_table(table):
    print("\n{:^12}{:^12}{:^12}\n".format("x", "y", "y`"))

    for i in range(len(table)):
        for j in range(len(table[i])):
            print("%-12f" %(table[i][j]), end = '')
        print()

    print()

def search_index(table, x, n):
    index = 0

    for i in table:
        if (i[0] > x):
            break
        index += 1

    if index >= len(table) - n:
        return len(table) - n - 1

    l_border = index
    r_border = index

    while (n > 0):
        if (r_border - index == index - l_border):
            if (l_border > 0):
                l_border -= 1
            else:
                r_border += 1
        else:
            if (r_border < len(table) - 1):
                r_border += 1
            else:
                l_border -= 1
        n -= 1

    return l_border

def divided_difference(x0, y0, x1, y1, y_der):
    if (abs(x0 - x1) > EPS):
        return (y0 - y1) / (x0 - x1)
    else:
        return y_der

def newton_polynomial(table, n, x):
    index = search_index(table, x, n)
    np = table[index][1]

    for i in range(n):
        for j in range(n - i):
            table[index + j][1] = divided_difference(
                table[index + j][0],         table[index + j][1],
                table[index + j + i + 1][0], table[index + j + 1][1], 
                table[index + j][2])

        mult = 1
        for j in range(i + 1):
            mult *= (x - table[index + j][0])

        mult *= table[index][1]
        np += mult
        
    return np

def table_extension(table):
    new_size = 2 * len(table)

    for i in range (0, new_size, 2):
        table.insert(i + 1, table[i][:])

def hermit_polynomial(name_file, n, x):
    table = read_file(name_file)
    table.sort(key = lambda array: array[0])
    table_extension(table)

    return newton_polynomial(table, n, x)

def root_search(name_file, n):
    table = read_file(name_file)
    table.sort(key = lambda array: array[1])

    for i in table:
        i[0], i[1] = i[1], i[0]

    return newton_polynomial(table, n, 0)

def main():
    name_file = "table.txt"

    table = read_file(name_file)
    if (table == []):
        return

    table.sort(key = lambda array: array[0])
    print_table(table)

    r, n, x = read_data(len(table))
    if (r):
        return

    np = newton_polynomial(table, n, x)
    hp = hermit_polynomial(name_file, n, x)
    root = root_search(name_file, n)

    print("\nNewton_p = %.6f" %(np))
    print("\nHermit_p = %.6f" %(hp))
    print("\nКорень данного полинома5 = %.6f\n" %(root))

if __name__ == "__main__":
    main()
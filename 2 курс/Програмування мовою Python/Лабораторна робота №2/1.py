# Завдання 1. Варіант 7

# Реалізувати дві функції користувача в одній програмі.


import math


def calc_z(m):
    if m == 3:
        print("Значення m не може бути рівним 3")
        exit()

    return math.sqrt((m + 3) / (m - 3))


def calc_y(n):
    y = 1
    for i in range(1, n + 1):
        y *= 2 * i

    return y


m = int(input("Введіть значення m: "))
print(f"Значення z: {calc_z(m)}")

n = int(input("Введіть значення n: "))
print(f"Значення y: {calc_y(n)}")

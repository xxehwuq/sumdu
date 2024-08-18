# Завдання 2. Варіант 7

# Реалізувати функцію 2 із завдання 1 у вигляді окремого модуля, підключити її в основну програму і продемонструвати роботу з нею.

import math
from module import calc_y


def calc_z(m):
    if m == 3:
        print("Значення m не може бути рівним 3")
        exit()

    return math.sqrt((m + 3) / (m - 3))


m = int(input("Введіть значення m: "))
print(f"Значення z: {calc_z(m)}")

n = int(input("Введіть значення n: "))
print(f"Значення y: {calc_y(n)}")

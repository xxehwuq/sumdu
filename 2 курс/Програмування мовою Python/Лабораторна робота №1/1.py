# Завдання 1. Варіант 7

# Написати програму, яка обчислює значення X в залежності від значень a та b, введених користувачем з клавіатури.
# У варіантах 1-10 числа a та b можуть бути лише додатніми, у варіантах 10-20 можуть приймати значення від 1 до 100.
# Реалізувати у програмі перевірку чисел a та b, введених користувачем

a = int(input("Введіть значення a: "))
b = int(input("Введіть значення b: "))

if a < 0 or b < 0:
    print("Числа a та b повинні бути додатними!")
    exit()

x = 0

if a > b:
    x = 5 * a + b
elif a == b:
    x = -125
elif a < b:
    x = (a - 5) / b

print(f"Значення X: {x}")

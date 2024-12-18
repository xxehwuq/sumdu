# Завдання 3. Варіант 7

# Реалізувати функцію, яка виконує операції над множинами – задану за варіантом та друк множини на екран. У випадку,
# якщо задану варіантом операцію над множиною виконати не можна, перетворіть множину у список, а потім при виведенні
# на екран результуючий список перетворіть на множину. Скласти програму, яка формує і виводить на екран множину
# квадратів цілих чисел в порядку зростання з діапазону 1..1000, кількість елементів множини задає користувач.

def squares_set():
    squares = set()

    n = int(input("Введіть кількість елементів множини: "))

    i = 1
    while i <= n:
        if i ** 2 <= 1000:
            squares.add(i ** 2)
            i += 1
        else:
            print(f"Число {i}^2 перевищує 1000")
            break

    print("Множина квадратів:", *sorted(squares))


squares_set()

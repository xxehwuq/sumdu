# Завдання 1. Варіант 7

# Реалізувати функцію, яка виконує операції над списками – задану за варіантом та друк списку на екран.
# Список користувач має вводити з клавіатури. Доповнення списку з обох кінців.

def extend_list():
    lst = list(map(int, input('Введіть елементи списку через пробіл: ').split()))

    first_elem = int(input("Введіть елемент для додавання на початок списку: "))
    last_elem = int(input("Введіть елемент для додавання в кінець списку: "))

    lst.insert(0, first_elem)
    lst.append(last_elem)

    print("Оновлений список:", lst)


extend_list()

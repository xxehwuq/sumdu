# Завдання 1. Варіант 7

# Розробити програму, яка: а) створює текстовий файл TF10_1 із символьних рядків різної довжини; б) читає вміст файла
# TF10_1, вилучає всі цифри і записує у файл TF10_2 по 10 символів у рядок (останній рядок може бути не повним); в)
# читає вміст файла TF10_2 і друкує його по рядках. У програмі реалізувати обробку помилок відкриття файлу за
# допомогою конструкції try-except.

def create(filename):
    try:
        file = open(filename, 'w')

        lines = [
            "Рядок з цифрами 1 2 3",
            "Інший рядок з числом 52",
            "Звичайний рядок",
        ]

        for line in lines:
            file.write(line + "\n")

        print(f"Файл {filename} створений")

    except Exception as e:
        print(f"Помилка під час створення файлу {filename}: {e}")


def remove_nums(input_filename, output_filename):
    try:
        input_file = open(input_filename, 'r')
        output_file = open(output_filename, 'w')

        content = input_file.read()

        lines = content.split("\n")

        for line in lines:
            new_line = ''

            for symbol in line:
                if not symbol.isdigit():
                    new_line += symbol

                if len(new_line) == 10:
                    new_line += "\n"

            output_file.write(new_line + "\n")

        print(f"Файл {output_filename} оброблений і збережений")

    except Exception as e:
        print(f"Помилка під час обробки файлу {input_filename}: {e}")


def read(filename):
    try:
        file = open(filename, 'r')

        print(f"Вміст файлу {filename}:")

        for line in file:
            print(line.strip())

    except Exception as e:
        print(f"Помилка під час читання файлу {filename}: {e}")


create("TF10_1.txt")
remove_nums("TF10_1.txt", "TF10_2.txt")
read("TF10_2.txt")

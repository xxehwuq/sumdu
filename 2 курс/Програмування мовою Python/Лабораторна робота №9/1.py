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
        with open(input_filename, 'r') as input_file, open(output_filename, 'w') as output_file:
            content = input_file.read()

            filtered_content = ''.join(char for char in content if not char.isdigit())
            filtered_content = filtered_content.replace("\n", "").strip()

            for i in range(0, len(filtered_content), 10):
                output_file.write(filtered_content[i:i + 10] + "\n")
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

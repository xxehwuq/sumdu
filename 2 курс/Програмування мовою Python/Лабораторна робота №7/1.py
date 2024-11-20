# Завдання 1. Варіант 7

# Задано дані про кількість опадів, які випали за кожен день місяця, і про температуру повітря в ці дні. Скласти
# програму, що визначає, яка кількість опадів випала у вигляді снігу і яка – у вигляді дощу (вважати, що дощ іде,
# якщо температура повітря вище 0° С).

weather_data = {
    "01.11.2024": {"опади": 10, "температура": .2},
    "02.11.2024": {"опади": 5, "температура": 3},
    "03.11.2024": {"опади": 8, "температура": 1},
    "04.11.2024": {"опади": 0, "температура": .5},
    "05.11.2024": {"опади": 12, "температура": 4},
}


def print_one_record(data, date):
    values = data[date]
    print(f"Дата: {date}, кількість опадів: {values['опади']} mm, температура: {values['температура']}°C")


def print_all_values(data):
    for date, _ in data.items():
        print_one_record(data, date)


def add_record(data):
    try:
        date = input("Вкажіть дату (наприклад 01.11.2024): ")
        precipitation = int(input("Вкажіть кількість опадів (mm): "))
        temperature = int(input("Вкажіть температуру (°C): "))

        if date in data:
            print("Ця дата вже існує")
        else:
            data[date] = {"опади": precipitation, "температура": temperature}
    except ValueError:
        print("Неправильне введення, введіть числові значення для опадів і температури")


def delete_record(data):
    date = input("Вкажіть дату (наприклад 01.11.2024): ")
    if date in data:
        del data[date]
    else:
        print("Для цієї дати не знайдено записів")


def view_sorted_by_keys(data):
    for date in sorted(data.keys()):
        print_one_record(data, date)


def calculate_snow_and_rain(data):
    snow = sum(values["опади"] for values in data.values() if values["температура"] <= 0)
    rain = sum(values["опади"] for values in data.values() if values["температура"] > 0)

    print(f"Загальна кількість снігових опадів: {snow} mm")
    print(f"Загальна кількість дощових опадів: {rain} mm")


def menu():
    while True:
        print("\nМеню вибору:")
        print("1. Показати всі значення")
        print("2. Додати запис")
        print("3. Видалити запис")
        print("4. Переглянути словник, відсортований за ключами")
        print("5. Розрахувати кількість снігових та дощових опадів")
        print("6. Вийти")
        choice = input("Вкажіть свій вибір (1-6): ")

        print("\n")

        if choice == "1":
            print_all_values(weather_data)
        elif choice == "2":
            add_record(weather_data)
        elif choice == "3":
            delete_record(weather_data)
        elif choice == "4":
            view_sorted_by_keys(weather_data)
        elif choice == "5":
            calculate_snow_and_rain(weather_data)
        elif choice == "6":
            break
        else:
            print("Неправильний вибір")


menu()

# Завдання 1. Варіант 7

# Задано дані про кількість опадів, які випали за кожен день місяця, і про температуру повітря в ці дні.
# Скласти програму, що визначає, яка кількість опадів випала у вигляді снігу і яка – у вигляді дощу
# (вважати, що дощ іде, якщо температура повітря вище 0°С).
import json


def read_json_file(filename):
    try:
        with open(filename, "r", encoding="utf-8") as file:
            data = json.load(file)
        return data

    except FileNotFoundError:
        return []


def write_json_file(filename, data):
    with open(filename, "w", encoding="utf-8") as file:
        json.dump(data, file, ensure_ascii=False, indent=4)


def display_json_content(filename):
    data = read_json_file(filename)
    print(json.dumps(data, ensure_ascii=False, indent=4))


def add_record(filename):
    data = read_json_file(filename)
    date = input("Введіть дату (наприклад 01.01.2001): ")
    precipitation = float(input("Введіть кількість опадів (мм): "))
    temperature = float(input("Введіть температуру повітря (°С): "))

    data.append({"date": date, "precipitation": precipitation, "temperature": temperature})
    write_json_file(filename, data)

    print("Новий запис додано")


def delete_record(filename):
    data = read_json_file(filename)
    date = input("Введіть дату запису для видалення (наприклад 01.01.2001): ")

    data = [record for record in data if record["date"] != date]
    write_json_file(filename, data)

    print("Запис видалено")


def search_record(filename):
    data = read_json_file(filename)
    date = input("Введіть дату для пошуку (наприклад 01.01.2001): ")

    result = [record for record in data if record["date"] == date]
    if result:
        print(json.dumps(result, ensure_ascii=False, indent=4))
    else:
        print("Запис не знайдено")


def analyze_precipitation(input_filename, output_filename):
    data = read_json_file(input_filename)
    rain = sum(record["precipitation"] for record in data if record["temperature"] > 0)
    snow = sum(record["precipitation"] for record in data if record["temperature"] <= 0)

    result = {"rain": rain, "snow": snow}
    write_json_file(output_filename, result)

    print(f"Результати записані у файл '{output_filename}'")


def main():
    input_filename = "weather_data.json"
    output_filename = "precipitation_analysis.json"

    while True:
        print("\nМеню:")
        print("1. Вивести вміст JSON-файлу")
        print("2. Додати новий запис")
        print("3. Видалити запис")
        print("4. Пошук запису")
        print("5. Аналіз опадів (дощ/сніг)")
        print("6. Вийти")
        choice = input("Оберіть опцію (1-6): ")

        if choice == "1":
            display_json_content(input_filename)
        elif choice == "2":
            add_record(input_filename)
        elif choice == "3":
            delete_record(input_filename)
        elif choice == "4":
            search_record(input_filename)
        elif choice == "5":
            analyze_precipitation(input_filename, output_filename)
        elif choice == "6":
            break
        else:
            print("Невірний вибір")


main()

# Завдання 1. Варіант 7

# Знайти дані Exports of goods and services (% of GDP) для уcіх країн світу за 2019 рік. Вивести вміст .csv файлу на
# екран. Організувати пошук найнижчого та найвищого значень показника та записати результат пошуку у новий .csv файл.

import csv

try:
    csvfile1 = open("file1.csv", "r")

    reader = csv.DictReader(csvfile1, delimiter=",")

    print("Country | 2019")
    print("======================")

    for row in reader:
        print(row["Country Code"], "|", row["2019"])

    csvfile1.close()

except Exception as e:
    print(f"Помилка під час читання CSV файлу: {e}")
    exit()

try:
    csvfile1 = open("file1.csv", "r")
    csvfile2 = open("file2.csv", "w")

    reader = csv.DictReader(csvfile1, delimiter=",")
    writer = csv.DictWriter(csvfile2, delimiter=",", fieldnames=["Критерій", "Країна", "2019"])

    max_value_country, min_value_country = "", ""
    max_value, min_value = 0.0, 0.0

    for index, row in enumerate(reader):
        if index == 0:
            max_value, min_value = float(row["2019"]), float(row["2019"])
            max_value_country, min_value_country = row["Country Code"], row["Country Code"]

        if row["2019"] != '':
            if float(row["2019"]) > max_value:
                max_value, max_value_country = float(row["2019"]), row["Country Code"]
            elif float(row["2019"]) < min_value:
                min_value, min_value_country = float(row["2019"]), row["Country Code"]

    writer.writeheader()
    writer.writerow({"Критерій": "Найвищий показник", "Країна": max_value_country, "2019": max_value})
    writer.writerow({"Критерій": "Найнижчий показник", "Країна": min_value_country, "2019": min_value})

    csvfile2.close()

except Exception as e:
    print(f"Помилка під час запису в CSV файл: {e}")
    exit()

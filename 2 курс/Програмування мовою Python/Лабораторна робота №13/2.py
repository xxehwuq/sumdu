# Завдання 2. Варіант 7

# Візуалізація даних з порталу відкритих даних https://databank.worldbank.org/home.aspx. Використайте бібліотеку Matplotlib.
# Самостійно оберіть предметну область, наприклад, Education Statistics, з якої візьміть показник, наприклад, Children out of school,
# primary в динаміці за останніх двадцять років (або інший період, якщо дані для цього періоду на порталі відсутні) для України та
# однієї з країн світу на вибір, наприклад, США.  Сформуйте масив даних для побудови графіку та напишіть програму для їх візуалізації.
# 2.1. На одній координатній осі побудуйте графіки, що показують динаміку показника для двох країн, підпишіть осі – по осі Х має
# відображатися рік, а по осі Y має відображатися значення показника.
# 2.2 Побудуйте стовпчасті діаграми, які відображатимуть значення показника для кожної з країн. Назву країни для побудови діаграми
# має вводити користувач з клавіатури.

import csv
import matplotlib.pyplot as plot

data = {}

try:
    csvfile = open("data.csv", "r")

    reader = csv.DictReader(csvfile, delimiter=",")

    for row in reader:
        country = row['Country Name']
        data[country] = {}

        for key, value in row.items():
            if '[YR' in key:
                year = key.split(' [')[0]
                # year = int(key.split(' [')[0])
                data[country][year] = float(value)

    csvfile.close()

except Exception as e:
    print(f"Помилка під час читання CSV файлу: {e}")
    exit()

def show_plot(country):
    plot.figure(figsize=(12, 5))
    plot.xlabel("Рік")
    plot.ylabel("Значення показника")

    if country != "":
        years = list(data[country].keys())
        indicators = list(data[country].values())

        plot.title(f"Birth rate, crude (per 1,000 people) for {country}", fontsize=14)
        plot.grid(axis='y', alpha=0.5)
        plot.bar(years, indicators, color='orange')
        plot.show()
    else:
        for data_country, values in data.items():
            years = list(values.keys())
            indicators = list(values.values())

            plot.plot(years, indicators, label=data_country)

        plot.title("Birth rate, crude (per 1,000 people)", fontsize=14)
        plot.legend()
        plot.grid(alpha=0.5)
        plot.show()

input_country = input("Введіть назву країни (Ukraine або United States) або залиште поле пустим для відображення даних двох країн: ")
if (input_country in data) or input_country == "":
    show_plot(input_country)
else:
    print(f"Невірно введена назва країни")

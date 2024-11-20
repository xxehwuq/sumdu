# Завдання 2. Варіант 7

# Побудуйте кругову діаграму на основі даних з предметної області лабораторної роботи №12.
# Використайте бібліотеку Matplotlib. На круговій діаграмі мають відображатися значення
# показників у відсотках, наприклад, відсоток дівчат та хлопців, які навчаються у певному класі,
# сектори діаграми повинні бути розфарбовані в різний колір, на діаграмі мають бути підписи.

import json
import matplotlib.pyplot as plt

def read_json_file(filename):
    try:
        with open(filename, "r", encoding="utf-8") as file:
            data = json.load(file)
        return data

    except FileNotFoundError:
        return []


weather_data = read_json_file("../Лабораторна робота №12/weather_data.json")

dates = [data["date"] for data in weather_data]
precipitation = [data["precipitation"] for data in weather_data]

fig, ax = plt.subplots()
ax.pie(precipitation, labels=dates)
ax.axis("equal")
plt.legend()
plt.show()

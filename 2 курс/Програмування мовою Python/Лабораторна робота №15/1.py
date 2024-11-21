# Завдання 1. Варіант 7

# Перетворіть словник, створений у лабораторній роботі №7, на датафрейм,
# у разі необхідності доповніть словник даними, виведіть його на екран.
# Виконайте агрегацію та групування даних із заданої предметної області.

import pandas as pd

weather_data = {
    "01.11.2024": {"precipitation": 10, "temperature": 2},
    "02.11.2024": {"precipitation": 5, "temperature": 3},
    "03.11.2024": {"precipitation": 8, "temperature": 1},
    "04.11.2024": {"precipitation": 0, "temperature": 5},
    "05.11.2024": {"precipitation": 12, "temperature": 4},
}

df = pd.DataFrame(weather_data).T

print("Початковий DataFrame:")
print(df)

print("Середня кількість опадів:", df["precipitation"].mean(), "мм")
print("Середнє значення температури:", df["temperature"].mean(), "°C")

print("Мінімальна кількість опадів:", df["precipitation"].min(), "мм")

print("Максимальна температура:", df["temperature"].max(), "°C")

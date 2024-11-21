# Завдання 2. Варіант 7

# Створіть датафрейм з даними використання велодоріжок за рік, заданий варіантом.
# 2013 рік. Визначте, який місяць найбільш популярний у велосипедистів

import pandas as pd

months = ["Січні", "Лютому", "Березні", "Квітні", "Травні", "Червні", "Липні", "Серпні", "Вересні", "Жовтні",
          "Листопаді", "Грудні"]
file_path = 'comptagevelo2013.csv'

df = pd.read_csv(file_path)
df['Date'] = pd.to_datetime(df['Date'])
df['Month'] = df['Date'].dt.month
df['Sum'] = df.iloc[:, 1:-1].sum(axis=1, skipna=True)

months_sum = df.groupby('Month')["Sum"].sum()

print(f"Максимальна кількість велосипедистів {months_sum.max()} в {months[months_sum.idxmax() - 1]}")

# Завдання 2. Варіант 7

# Задано речення. Скласти програму, яка визначає і виводить на екран слова, які містять хоча б одну літеру «о».

sentence = str(input("Введіть речення: "))

words = sentence.split()

words_o = []

for word in words:
    if word.lower().count("о") >= 1:
        words_o.append(word)

print("Слова, що містять літеру 'о':", words_o)

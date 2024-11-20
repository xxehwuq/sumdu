# Завдання 2. Варіант 7

# Задано два слова. Скласти програму, яка визначає і виводить ті літери слів, які зустрічаються в обох словах тільки один раз.

word1 = str(input("Введіть перше слово: "))
word2 = str(input("Введіть друге слово: "))

letters = []

for letter in word1:
    if word1.count(letter) == 1 and word2.count(letter) == 1:
        letters.append(letter)

print("Літери, які зустрічаються в обох словах тільки один раз: ", letters)

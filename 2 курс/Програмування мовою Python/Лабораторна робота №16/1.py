# Завдання 1. Варіант 7

# Імпортувати бібліотеку NLTK та тексти із електронного архіву  текстів Project Gutenberg,
# для виконання завдань взяти текст, заданий варіантом.
# Визначити кількість слів у текст.
# Визначити 10 найбільш вживаних слів у тексті, побудувати на основі цих даних стовпчасту діаграму.
# Виконати видалення з тексту стоп-слів та пунктуації, після чого знову знайти 10 найбільш вживаних
# слів у тексті та побудувати на їх основі стовпчасту діаграму.
#
# burgess-busterbrown.txt

import nltk
from nltk.corpus import stopwords
import matplotlib.pyplot as plt
from collections import Counter
import string

nltk.download('stopwords')

try:
    file = open('burgess-busterbrown.txt', 'r', encoding='utf-8')

except FileNotFoundError:
    print("Файл не знайдено!")
    exit(0)

def count_words(text):
    sentences = nltk.sent_tokenize(text)
    k_words = 0

    for sentence in sentences:
        words = nltk.word_tokenize(sentence)
        k_words += len(words)

    return k_words

def remove_stopwords(text):
    stop_words = set(stopwords.words('english'))
    cleaned_text_arr = [word.lower() for word in text if word.lower() not in stop_words and word not in string.punctuation]
    return ''.join(str(letter) for letter in cleaned_text_arr)


def most_used_words(text, title):
    text1 = text.split()
    cnt = Counter(text1)
    cort = cnt.most_common(10)

    x = [cort[el][0] for el in range(len(cort))]
    y = [cort[el][1] for el in range(len(cort))]

    plt.bar(x, y)
    plt.title(title)
    plt.xlabel("Слова")
    plt.ylabel("Кількість у тексті")

    plt.show()


text = file.read()

print(f"\nКількість слів у тексті: {count_words(text)}")
print(f"Кількість слів у тексті (без стоп-слів): {count_words(text)} \n")

cleaned_text = remove_stopwords(text)

most_used_words(text, "10 найбільш вживаних слів у тексті")
most_used_words(cleaned_text, "10 найбільш вживаних слів у тексті (без стоп-слів)")

# Завдання 1. Варіант 7

# Перший студент має розмістити текст обсягом до 100 слів і написати програмний код обробки тексту,
# застосувавши щонайменше три вбудовані функції Python для обробки рядків. Програмний код має
# супроводжуватися коментарями із поясненням.
# Усі студенти повинні працювати з текстом, який розмістив перший студент. Кожен студент у відповіді
# на завдання для спільної роботи окрім посилання на програмний код розміщує короткий опис функцій,
# які він застосував, та зазначає три функції, які потрібно застосувати до цього тексту наступному студенту.
# Окремі функції по роботі з текстом можуть повторюватися, але кожен студент має використати щонайменше дві нові функції.


# Функція для визначення кількості слів у тексті
def count_words(text):
    words = text.split()
    return len(words)


# Функція для визначення літери, яка найчастіше зустрічається у всіх словах
def most_common_letter_in_words(text):
    from collections import Counter
    words = text.split()
    all_letters = "".join(words)  # Збираємо всі літери з усіх слів
    letter_counts = Counter(all_letters)
    return letter_counts.most_common(1)[0][0]


# Функція, яка виводить слова, що містять задану літеру
def words_with_letter(text, letter):
    words = text.split()
    return [word for word in words if letter in word]


# Функція для перевірки, чи текст починається і закінчується однаковою літерою
def starts_and_ends_with_same_letter(text):
    first_letter = text[0].lower()
    last_letter = text[-1].lower()

    if first_letter == last_letter:
        return "Так"
    else:
        return "Ні"


# Функція, яка підраховує кількість унікальних слів у тексті
def count_unique_words(text):
    words = text.lower().split()
    unique_words = set(words)
    return len(unique_words)


# Оригінальний текст
text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum accumsan aliquam mi, ac pretium justo finibus ut. Pellentesque viverra varius nisl, eget egestas massa consequat a. Pellentesque fermentum nulla nec luctus pharetra. Cras in nibh eu velit pellentesque auctor pellentesque vitae nisl. Quisque urna orci, elementum at sodales at, tempor."

# Використання функцій
word_count = count_words(text)
print(f"Кількість слів у тексті: {word_count}")

common_letter = most_common_letter_in_words(text)
print(f"Літера, яка найчастіше зустрічається в усіх словах: '{common_letter}'")

words_with_common_letter = words_with_letter(text, common_letter)
print(f"Слова, що містять літеру '{common_letter}': {words_with_common_letter}")

same_start_end = starts_and_ends_with_same_letter(text)
print(f"Чи починається і закінчується текст однаковою літерою? {same_start_end}")

unique_word_count = count_unique_words(text)
print(f"Кількість унікальних слів у тексті: {unique_word_count}")

# Імпортування необхідних бібліотек
import xml.etree.ElementTree as ET
from fileinput import filename

import pandas as pd
import matplotlib.pyplot as plot
from collections import Counter


# Функція для парсингу XML-файлу
def parse_file(filename):
    # Завантаження XML-файлу та отримання його кореневого елемента
    tree = ET.parse(filename)
    root = tree.getroot()

    # Ініціалізація списку для зберігання інформації про книги
    books = []

    # Ітерація по кожному елементу <book> у файлі
    for book in root.findall('book'):
        # Збереження даних про книгу у вигляді словника
        books.append({
            "Title": book.find("title").text,  # Назва книги
            "Author": book.find("author").text,  # Автор книги
            "Genre": book.find("genre").text,  # Жанр книги
            "Year": book.find("year").text  # Рік публікації
        })

    # Повернення списку книг
    return books


# Функція для відображення всіх книг у табличному форматі
def show_all_books(books):
    # Перетворення списку словників у DataFrame (таблицю) та її відображення
    df = pd.DataFrame(books)
    print(df)


# Функція для побудови графіка розподілу книг за роками
def show_plot_by_year(books):
    # Підрахунок кількості книг для кожного року за допомогою Counter
    year_counts = Counter(book["Year"] for book in books)

    # Сортування років для коректного відображення на графіку
    years = sorted(year_counts.keys())
    counts = [year_counts[year] for year in years]

    # Створення стовпчикового графіка
    plot.figure(figsize=(10, 6))  # розмір графіка
    plot.bar(years, counts)  # побудова стовпчиків
    plot.title("Розподіл книг за роками")  # заголовок графіка
    plot.xlabel("Рік")  # підпис осі X
    plot.ylabel("Кількість книг")  # підпис осі Y
    plot.xticks(rotation=45)  # нахил підписів на осі X для читабельності
    plot.grid(axis='y', linestyle="solid", alpha=0.5)  # додавання сітки для осі Y
    plot.show()  # відображення графіка


# Функція для побудови кругової діаграми розподілу книг за жанрами
def show_plot_by_genre(books):
    # Підрахунок кількості книг для кожного жанру за допомогою Counter
    genre_counts = Counter(book["Genre"] for book in books)

    # Отримання списків жанрів і їх кількостей
    labels = list(genre_counts.keys())
    sizes = list(genre_counts.values())

    # Створення кругової діаграми
    plot.figure(figsize=(8, 8))
    plot.pie(sizes, labels=labels)
    plot.title("Розподіл книг за жанрами")
    plot.show()


# Функція для відображення меню користувача та взаємодії з програмою
def menu(books):
    while True:
        # Виведення меню з доступними опціями
        print("\nМеню:")
        print("1. Вивести всі книги")
        print("2. Показати графік розподілу за роками")
        print("3. Показати графік розподілу за жанрами")
        print("4. Вийти")
        choice = input("Оберіть опцію (1-4): ")  # Запит вибору у користувача

        # Обробка вибору користувача
        if choice == "1":
            show_all_books(books)
        elif choice == "2":
            show_plot_by_year(books)
        elif choice == "3":
            show_plot_by_genre(books)
        elif choice == "4":
            break
        else:
            print("Некоректний вибір")  # Повідомлення про некоректний ввід


# Основна частина програми
filename = "data.xml"  # Змінна з назвою файла

books = parse_file(filename)  # Виклик функції парсингу для зчитування даних із XML-файлу
menu(books)  # Запуск меню для взаємодії з користувачем

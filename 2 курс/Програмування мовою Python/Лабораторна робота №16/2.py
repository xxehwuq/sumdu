# Завдання 2. Варіант 7

# Зберегти у текстовому файлі довільний текст до 100 слів. Виконати наступні дії над текстом:
# Токенізація по словам
# Лемматизація та стеммінг
# Видалення стоп-слів
# Видалення пункуації
# Записати оброблений текст у інший файл.
# Для виконання завдань лабораторної роботи використати бібліотеку NLTK.

import nltk
from nltk.tokenize import word_tokenize
from nltk.corpus import stopwords
from nltk.stem import WordNetLemmatizer, PorterStemmer
import string

nltk.download('punkt')
nltk.download('stopwords')
nltk.download('wordnet')
nltk.download('omw-1.4')

with open('input_text.txt', 'r') as file:
    text = file.read()

tokens = word_tokenize(text)

lemmatizer = WordNetLemmatizer()
stemmer = PorterStemmer()

lemmas = [lemmatizer.lemmatize(token.lower()) for token in tokens]
stems = [stemmer.stem(token.lower()) for token in tokens]

stop_words = set(stopwords.words('english'))
filtered_tokens = [token for token in lemmas if token not in stop_words]

processed_tokens = [token for token in filtered_tokens if token not in string.punctuation]

processed_text = ' '.join(processed_tokens)
with open('processed_text.txt', 'w') as file:
    file.write(processed_text)

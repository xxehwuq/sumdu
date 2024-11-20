# Завдання 2. Варіант 7

#  Заповнити двовимірний масив розміром 7x7 таким чином, як показано на рисунку згідно з Вашим варіантом.
#  Вивести масив на екран. Для виконання завдання використовуйте цикли.

# [7, 7, 7, 7, 7, 7, 7]
# [6, 6, 6, 6, 6, 6, 6]
# [5, 5, 5, 5, 5, 5, 5]
# [4, 4, 4, 4, 4, 4, 4]
# [3, 3, 3, 3, 3, 3, 3]
# [2, 2, 2, 2, 2, 2, 2]
# [1, 1, 1, 1, 1, 1, 1]

n = 7
arr = [[0] * n for _ in range(n)]

for i in range(n):
    for k in range(n):
        arr[i][k] = n - i

for r in arr:
    print(r)

# Завдання 1. Варіант 7

# Побудуйте графік функції. Оберіть суцільний тип лінії, задайте колір
# та товщину графіку та позначте осі, виведіть назву графіка на екран,
# вставте легенду. Використайте бібліотеку Matplotlib.
# Y(x)=sin(10*x)*sin(3*x)/(x^2), x=[0...4]

import math
import matplotlib.pyplot as plot

x_values = [i/100 for i in range(1, 400)]

y_values = [math.sin(10 * x) * math.sin(3 * x) / (x**2) for x in x_values]

plot.figure(figsize=(7, 7))
plot.plot(x_values, y_values, linestyle='solid', color='red', linewidth=2, label="Y(x)=sin(10*x)*sin(3*x)/(x^2)")

plot.xlabel("x")
plot.ylabel("y")
plot.title("Y(x)=sin(10*x)*sin(3*x)/(x^2)", fontsize=14)

plot.legend()
plot.grid(alpha=0.5)
plot.show()
# Завдання 1. Варіант 7

# Побудуйте графік функції. Оберіть суцільний тип лінії, задайте колір
# та товщину графіку та позначте осі, виведіть назву графіка на екран,
# вставте легенду. Використайте бібліотеку Matplotlib.
# Y(x)=sin(10*x)*sin(3*x)/(x^2), x=[0...4]

import numpy as np
import matplotlib.pyplot as plot

x = np.linspace(0.01, 4, 100)
y = np.sin(10 * x) * np.sin(3 * x) / (x ** 2)

plot.figure(figsize=(8, 8))
plot.plot(x, y, linestyle='solid', color='blue', linewidth=2, label="Y(x)=sin(10*x)*sin(3*x)/(x^2), x=[0...4]")

plot.xlabel("x")
plot.ylabel("y(x)")
plot.title("Y(x)=sin(10*x)*sin(3*x)/(x^2), x=[0...4]", fontsize=14)

plot.legend()
plot.grid(alpha=0.5)
plot.show()

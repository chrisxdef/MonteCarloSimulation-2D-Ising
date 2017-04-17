import pandas as pd
import matplotlib.pyplot as plt
import sys

df = pd.read_csv(sys.argv[1])
df.astype(float)
df.plot(kind='scatter', x='temperature', y='ave_magnetization', title='Average Magnetization')
df.plot(kind='scatter', x='temperature', y='ave_magnetization^2', title='Average Magnetization^2')
df.plot(kind='scatter', x='temperature', y='susceptibility', title='Susceptibility')
df.plot(kind='scatter', x='temperature', y='ave_energy', title='Average Energy')
df.plot(kind='scatter', x='temperature', y='ave_energy^2', title='Average Energy^2')
df.plot(kind='scatter', x='temperature', y='C_v', title='Specific Heat')
plt.show()

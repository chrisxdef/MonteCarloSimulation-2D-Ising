import pandas as pd
import matplotlib.pyplot as plt
import sys

df = pd.read_csv(sys.argv[1])

df.plot(kind='scatter', x='temperature', y='ave_magnetization')
df.plot(kind='scatter', x='temperature', y='ave_magnetization^2')
df.plot(kind='scatter', x='temperature', y='susceptibility')
df.plot(kind='scatter', x='temperature', y='ave_energy')
df.plot(kind='scatter', x='temperature', y='ave_energy^2')
df.plot(kind='scatter', x='temperature', y='C_v')
plt.show()

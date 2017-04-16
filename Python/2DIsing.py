# Monte Carlo Simulation of the 2-D Ising Model
# This program is largely adapted from Lisa Larrimore's FORTRAN 90 code

import math
import random
import json

# Function prints array, for testing
def print_array():
    for i in range(size):
        for j in range(size):
            print(str(spinArray[i][j]), end=" ")
        print()

# Read from input file
inputFile = open("input.json")
parsedJSON = json.loads(inputFile.read())

# Define parameters
# For this version of the program, many things will be hard coded
# Future versions may use File I/O
size = parsedJSON['nrows'] # size of the square 2-D Ising Model
nSamples = parsedJSON['npass'] # number of random samples
nEquil = parsedJSON['nequil'] # passes until equilibrium is "reached"
highTemp = parsedJSON['high_temp'] # Upper limit of temperatures to scan over
lowTemp = parsedJSON['low_temp'] # Lower limit of temperatures to scan over
tempInt = parsedJSON['temp_interval'] # Interval of scans

nScans = int((highTemp - lowTemp) / tempInt)

# Create output files and write function
magFile = open("magnetization.csv", "w")
magFile.write("Temp,Ave_magnetization,Ave_magnetization^2\n")
energyFile = open("energy.csv", "w")
energyFile.write("Temp,Ave_energy,Ave_energy^2\n")
bothFile = open ("both.csv", "w")
bothFile.write("temperature,ave_magnetization,ave_magnetization^2,ave_energy,ave_energy^2\n")

def write_vals(temp, magAvg, magAvgSq, energyAvg, energyAvgSq):
    magFile.write(str(round(temp, 4)) + "," + str(round(magAvg, 4)) + "," + str(round(magAvgSq, 4)) + "\n")
    energyFile.write(str(round(temp, 4)) + "," + str(round(energyAvg, 4)) + "," + str(round(energyAvgSq, 4)) + "\n")
    bothFile.write(str(round(temp, 4)) + "," + str(round(magAvg, 4)) + "," + str(round(magAvgSq, 4)) + ","
                   + str(round(energyAvg, 4)) + "," + str(round(energyAvgSq, 4)) + "\n")


# Begin scan loop
for iScan in range(nScans):
    # Initialize variables
    temp = float(highTemp - tempInt*iScan)
    beta = 1/temp
    outputCount = 0.0
    energyAvg = 0.0
    energyAvgSq = 0.0
    magAvg = 0.0
    magAvgSq = 0.0

    # The starting spin array begins as a checkerboard pattern.
    # Python lists are immediately circular, so the periodic boundary conditions hold regardless
    spinArray = [[0 for x in range(size)] for y in range(size)]
    for i in range(0, size, 2):
        for j in range(0, size, 2):
            spinArray[i][j] = 1.0
            spinArray[i][j + 1] = -1.0
            spinArray[i + 1][j] = -1.0
            spinArray[i + 1][j + 1] = 1.0

    for iPass in range(nSamples):
        # We only calculate the magnetization and energy once the equilibration steps have
        # been completed.
        if iPass > nEquil:
            outputCount += 1
            magnetization = 0.0
            energy = 0.0
            for i in range(size - 1): # last index needs to be specially treated
                for j in range(size - 1):
                    magnetization += spinArray[i][j]
                    energy += spinArray[i][j]*(spinArray[i-1][j] + spinArray[i+1][j] +
                                               spinArray[i][j-1] + spinArray[i][j+1])
            for i in range(size - 1): # last index in rows
                    magnetization += spinArray[i][size - 1]
                    energy += spinArray[i][size - 1] * (spinArray[i - 1][size - 1] + spinArray[i + 1][size - 1] +
                                                        spinArray[i][size - 2] + spinArray[i][0])
            for j in range(size - 1): # last index in columns
                    magnetization += spinArray[size - 1][j]
                    energy += spinArray[size - 1][j] * (spinArray[size - 2][j] + spinArray[0][j] +
                                                        spinArray[size - 1][j - 1] + spinArray[size - 1][j+1])
            # finally last index in both
            magnetization += spinArray[size - 1][size - 1]
            energy += spinArray[size - 1][size - 1] * (spinArray[size - 2][size - 1] + spinArray[0][size - 1] +
                                                       spinArray[size - 1][size - 2] + spinArray[size - 1][0])
            magnetization /= size**2
            energy /= (2.0*size**2)
            magAvg += magnetization
            magAvgSq += magnetization**2
            energyAvg += energy
            energyAvgSq += energy**2

        # Pick a random spin to change and determine if it is worth accepting
        m = int(size * random.random()) # pick a random column
        n = int(size * random.random()) # pick a random row
        trialSpin = -1 * spinArray[m][n] # flip the spin of the chosen atom

        # Find the energy change due to the spin change
        # if exp(-beta*deltaE) > a random number [0,1], we accept the state
        if m == (size - 1): # for outer periodic boundary condition
            i = 0
        else:
            i = m + 1
        if n == (size - 1): # for outer periodic boundary condition
            j = 0
        else:
            j = n + 1
        deltaE = -2 * trialSpin * (spinArray[m-1][n] + spinArray[i][n] +
                                   spinArray[m][n-1] + spinArray[m][j])
        exp = math.exp(-1 * beta * deltaE)
        if exp > random.random():
            spinArray[m][n] = trialSpin
    # Calculate and write final values
    magAvg = math.fabs(magAvg / outputCount)
    energyAvg = math.fabs(energyAvg / outputCount)
    magAvgSq = math.fabs(magAvgSq / outputCount)
    energyAvgSq = math.fabs(energyAvgSq / outputCount)
    write_vals(temp, magAvg, magAvgSq, energyAvg, energyAvgSq)


import os

year = input("Insert the current year: ")

os.system(f"mkdir {year}")

for i in range(1, 26):
    os.system(f"mkdir {year}/{i}")
    os.system(f"touch {year}/{i}/{i}-a.py {year}/{i}/{i}-a-test.txt {year}/{i}/{i}-b.py {year}/{i}/{i}-b-test.txt \
                {year}/{i}/{i}-a.txt {year}/{i}/{i}-b.txt")



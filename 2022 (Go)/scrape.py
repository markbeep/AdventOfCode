"""
Fetches the input data for the given day or simply today and throws it
in the folder
"""

import os
import requests, sys
from datetime import datetime
from dotenv import load_dotenv

load_dotenv()

# cookie session key
SESSION_KEY = os.getenv("SESSION")
HEADERS = {"User-Agent": f"Fetch tool by {os.getenv('EMAIL', 'n/a')}"}

def fetch(day: int):
    r = requests.get(
        f"https://adventofcode.com/2022/day/{day}/input",
        cookies={"session": SESSION_KEY},
        headers=HEADERS)
    try:
        with open(f"{str(day).zfill(2)}/inp.txt", "w") as f:
            f.write(r.text)
    except FileNotFoundError:
        print("Day folder doesn't exist")
        
if __name__ == "__main__":
    args = sys.argv
    if len(args) < 2:
        print("No day given. Taking today's date")
        day = datetime.now().day
    else:
        day = args[1]
    
    fetch(day)

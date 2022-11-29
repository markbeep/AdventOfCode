"""
Fetches the input data for the given day or simply today and throws it
in the folder
"""

import requests, sys
from datetime import datetime

# cookie session key
SESSION_KEY = "53616c7465645f5f22342b2890cce3b49cc607014daee0af67abbd7fd427651a4706773d32654d42a1e9c7673161f68f8b0c7a2e5d5e4bc2e62dad851071e58a"

def fetch(day: int):
    r = requests.get(
        f"https://adventofcode.com/2022/day/{day}/input",
        cookies={"session": SESSION_KEY})
    try:
        with open(f"{day}/inp.txt", "w") as f:
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

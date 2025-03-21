import requests
import time

BASE_URL = "http://localhost:8080"

def get_stocks():
    url = f"{BASE_URL}/get-stocks"
    response = requests.get(url)
    if response.status_code == 200:
        return response.json()
    return []

if __name__ == "__main__":
    while True:
        print(get_stocks())
        time.sleep(10)
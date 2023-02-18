import requests

resp = requests.get("https://data.lpdaac.earthdatacloud.nasa.gov/s3credentials")
print(resp.json())

import requests
import boto3
import json
import urllib.request

provider = 'ASF'
CMR_OPS = 'https://cmr.earthdata.nasa.gov/search'
url = f'{CMR_OPS}/{"collections"}'
response1 = requests.get(url,
                        params={
                            'cloud_hosted': 'True',
                            'has_granules': 'True',
                            'provider': provider,
                        },
                        headers={
                            'Accept': 'application/json'
                        }
                        )

print(response1.json()['feed']['entry'][0])



response = requests.get(url,
                        params={
                            'concept_id': 'C1214470488-ASF',
                            # 'temporal': '2020-10-17T00:00:00Z,2020-10-18T23:59:59Z',
                            # 'bounding_box': '76.08166,-67.1746,88.19689,21.04862',
                            'page_size': 200,
                        },
                        headers={
                            'Accept': 'application/json'
                        }
                        )
print(response.status_code)

response = requests.get("https://cmr.earthdata.nasa.gov/search/concepts/C1214470488-ASF")
print(response)
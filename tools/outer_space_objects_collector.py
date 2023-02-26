import json
import time

import requests

BASE_URL = "https://www.unoosa.org/oosa/osoindex/waxs-search.json"
start_at = 0


def call(url):
    print(f"calling {url}")
    try:
        response = requests.get(url)
        if response.status_code != 200:
            print()
        response_json = response.json()
        found = response_json['found']
        results = response_json['results']
        return found, results
    except:
        print("An exception occurred")
        time.sleep(10)
        return None, None

outer_space_objects = []
pre_found = 0

while start_at == 0 or start_at < found:
    oso_index_url = f"{BASE_URL}?criteria={{\"startAt\":{start_at}}}"
    found, results = call(oso_index_url)

    if found is None:
        found = pre_found
    else:
        pre_found = found
        for result in results:
            values = result['values']
            name = f"{values['object.nameOfSpaceObjectIno_s1']} - {values['object.nameOfSpaceObjectO_s1']}" if 'object.nameOfSpaceObjectIno_s1' in values else values['object.nameOfSpaceObjectO_s1']
            object = {
                "name": name,
                "launched_at": values['object.launch.dateOfLaunch_s1'] if 'object.launch.dateOfLaunch_s1' in values else None,
                "decayed_at": values['object.status.dateOfDecay_s1'] if 'object.status.dateOfDecay_s1' in values else None,
                "status": values['en#object.status.objectStatus_s1'] if 'en#object.status.objectStatus_s1' in values else None,
                "country": values['object.launch.stateOfRegistry_s1'] if 'object.launch.stateOfRegistry_s1' in values else None,
                "function": values['object.functionOfSpaceObject_s1'] if 'object.functionOfSpaceObject_s1' in values else None
            }
            outer_space_objects.append(object)

        start_at += len(results)


json_object = json.dumps(outer_space_objects, indent=4)
with open('files/outer_space_objects.json', 'w') as file:
    file.write(json_object)

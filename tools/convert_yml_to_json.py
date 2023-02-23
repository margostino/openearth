import yaml
from yaml.loader import SafeLoader
import json

# filename = "/Users/martin.dagostino/workspace/margostino/data/openearth/nasa_rss_feeds.yml"
filename = "/Users/martin.dagostino/workspace/margostino/openearth/tools/files/datasets.yml"

with open(filename, 'r') as yaml_file:
    data = yaml.load(yaml_file, Loader=SafeLoader)
    json_object = json.dumps(data, indent=4)
    with open('files/datasets.json', 'w') as json_file:
        json_file.write(json_object)


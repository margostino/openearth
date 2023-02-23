from python_graphql_client import GraphqlClient
import yaml

# Instantiate the client with an endpoint.
client = GraphqlClient(endpoint="https://graphql.earthdata.nasa.gov/api")

datasets = []
# Create the query string and variables required for the request.
with open("nasa_datasets_query.graphql", 'r') as file:
    query = file.read()
    variables = {"params": {"limit": None}}

    # Synchronous request
    response = client.execute(query=query, variables=variables)
    collections = response['data']['collections']['items']
    # Asynchronous request
    # import asyncio
    # data = asyncio.run(client.execute_async(query=query, variables=variables))
    # print(data)

    for collection in collections:
        dataset = {}

        category = [keyword['category'] for keyword in collection['scienceKeywords']][0]
        keywords = list(dict.fromkeys([keyword['topic'] for keyword in collection['scienceKeywords']] + [keyword['term'] for keyword in collection['scienceKeywords']]))

        dataset['did'] = collection['datasetId']
        dataset['status'] = collection['collectionProgress']
        dataset['format'] = collection['metadataFormat']
        dataset['coordinate_system'] = collection['coordinateSystem']
        dataset['start_at'] = collection['timeStart']
        dataset['end_at'] = collection['timeEnd']
        dataset['version'] = collection['version']
        dataset['category'] = category
        dataset['keywords'] = keywords
        dataset['provider'] = collection['provider']

        dataset['distributors'] = []
        dataset['archivers'] = []
        for data_center in collection['dataCenters']:
            long_name = data_center['longName'] if 'longName' in data_center else None
            short_name = data_center['shortName']
            if 'DISTRIBUTOR' in data_center['roles']:
                dataset['distributors'].append(f"{long_name} ({short_name})")
            if 'ARCHIVER' in data_center['roles']:
                dataset['archivers'].append(f"{long_name} ({short_name})")

        dataset['description'] = f"{collection['title']} - {collection['abstract']}"
        datasets.append(dataset)

with open('files/datasets.yml', 'w') as file:
    yml_dump = yaml.dump(datasets, file, indent=4, sort_keys=False, default_flow_style=False, allow_unicode=True, encoding=None)
    print(yml_dump)
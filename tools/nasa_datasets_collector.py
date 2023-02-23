import json

from python_graphql_client import GraphqlClient

from utils import safe_get_or_default, get_first_not_none

CMR_BASE_URL = "https://cmr.earthdata.nasa.gov/search/concepts"
EARTHDATA_BASE_URL = "https://search.earthdata.nasa.gov"
EARTHDATA_SEARCH_PATH = "/search/granules"

client = GraphqlClient(endpoint="https://graphql.earthdata.nasa.gov/api")

datasets = []

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

        concept_id = collection['conceptId']
        category = [keyword['category'] for keyword in collection['scienceKeywords']][0]
        keywords = list(dict.fromkeys([keyword['topic'] for keyword in collection['scienceKeywords']] + [keyword['term'] for keyword in collection['scienceKeywords']]))

        dataset['dataset_id'] = collection['nativeId']
        dataset['status'] = collection['collectionProgress']
        dataset['version'] = collection['version']
        dataset['start_at'] = collection['timeStart']
        dataset['end_at'] = collection['timeEnd']

        if collection['dataDates'] is not None:
            data_create_dates = [data_date['date'] for data_date in collection['dataDates'] if data_date['type'] == 'CREATE']
            data_update_dates = [data_date['date'] for data_date in collection['dataDates'] if data_date['type'] == 'UPDATE']
            dataset['created_at'] = data_create_dates[0] if len(data_create_dates) > 0 else None
            dataset['updated_at'] = data_update_dates[0] if len(data_update_dates) > 0 else None
        else:
            dataset['created_at'] = None
            dataset['updated_at'] = None

        dataset['dataset_url'] = f"{EARTHDATA_BASE_URL}{EARTHDATA_SEARCH_PATH}?p={concept_id}"
        dataset['metadata_url'] = f"{CMR_BASE_URL}/{concept_id}"
        dataset['format'] = collection['metadataFormat']
        dataset['coordinate_system'] = collection['coordinateSystem']
        dataset['category'] = category
        dataset['keywords'] = keywords
        dataset['provider'] = collection['provider']
        dataset['organizations'] = [organization for organization in collection['organizations']]
        dataset['distributors'] = []

        dataset['archivers'] = []
        for data_center in collection['dataCenters']:
            data_center_long_name = data_center['longName'] if 'longName' in data_center else None
            data_center_short_name = data_center['shortName']
            if 'DISTRIBUTOR' in data_center['roles']:
                dataset['distributors'].append(f"{data_center_long_name} ({data_center_short_name})")
            if 'ARCHIVER' in data_center['roles']:
                dataset['archivers'].append(f"{data_center_long_name} ({data_center_short_name})")

        dataset['platforms'] = []
        for platform in collection['platforms']:
            dataset_platform_instruments = []

            if 'instruments' in platform:
                for instrument in platform['instruments']:
                    instrument_long_name = safe_get_or_default(instrument, 'longName')
                    instrument_short_name = safe_get_or_default(instrument, 'shortName')
                    instrument_name = f"{instrument_long_name} ({instrument_short_name})" if instrument_long_name is not None and instrument_short_name is not None else get_first_not_none(instrument_long_name, instrument_short_name)
                    dataset_platform_instrument = {
                        "name": instrument_name
                    }
                    dataset_platform_instruments.append(dataset_platform_instrument)

            platform_type = safe_get_or_default(platform, 'type')
            platform_long_name = safe_get_or_default(platform, 'longName')
            platform_short_name = safe_get_or_default(platform, 'longName')
            platform_name = f"{platform_long_name} ({platform_short_name})" if (platform_long_name is not None and platform_short_name is not None) else get_first_not_none(platform_long_name,
                                                                                                                                                                            platform_short_name)

            dataset_platform = {
                "platform_type": platform_type,
                "name": platform_name,
                "instruments": dataset_platform_instruments,
            }
            dataset['platforms'].append(dataset_platform)

        if 'useConstraints' in collection and collection['useConstraints'] is not None:
            constrains = collection['useConstraints']
            license_url = constrains['licenseUrl'] if 'licenseUrl' in constrains else None
            constrain_link = license_url['linkage'] if license_url is not None and 'linkage' in license_url else None
            constrain_name = license_url['name'] if license_url is not None and 'name' in license_url else None
            constrain_description = constrains['description'] if 'description' in constrains else (license_url['description'] if license_url is not None and 'description' in license_url else None)

            dataset['constraints'] = {
                "link": constrain_link,
                "name": constrain_name,
                "description": constrain_description
            }
        else:
            dataset['constraints'] = {}

        dataset['description'] = f"{collection['title']} - {collection['abstract']}"
        datasets.append(dataset)

json_object = json.dumps(datasets, indent=4)
with open('files/datasets.json', 'w') as file:
    file.write(json_object)

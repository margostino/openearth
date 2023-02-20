from python_graphql_client import GraphqlClient


# Instantiate the client with an endpoint.
client = GraphqlClient(endpoint="https://graphql.earthdata.nasa.gov/api")

# Create the query string and variables required for the request.
with open("nasa_datasets_query.graphql", 'r') as file:
    query = file.read()
    variables = {"params": {"limit": None}}

    # Synchronous request
    data = client.execute(query=query, variables=variables)
    print(data)

# Asynchronous request
# import asyncio
# data = asyncio.run(client.execute_async(query=query, variables=variables))
# print(data)
from couchbase.cluster import Cluster, ClusterOptions
from couchbase.cluster import PasswordAuthenticator
from dotenv import load_dotenv
import os

# Load environment variables from .env file
load_dotenv()

# Retrieve the Couchbase connection details from environment variables
couchbase_endpoint = os.getenv('COUCHBASE_ENDPOINT')
couchbase_username = os.getenv('COUCHBASE_USER')
couchbase_password = os.getenv('COUCHBASE_PASSWORD')
couchbase_bucket = os.getenv('COUCHBASE_BUCKET')

# Connect to Couchbase using the loaded environment variables
cluster = Cluster(couchbase_endpoint, ClusterOptions(
    PasswordAuthenticator(couchbase_username, couchbase_password)))
bucket = cluster.bucket(couchbase_bucket)
collection = bucket.default_collection()

# Run a simple query to get all documents in the bucket
query_string = "SELECT * FROM `{}` LIMIT 10".format(couchbase_bucket)
query_result = cluster.query(query_string)

# Print the query results
for row in query_result:
    print(row)

# Indicate that the script finished successfully
print("Query executed successfully")
import requests
import time
import matplotlib.pyplot ass plt

nPage = 100
if nPage > 100:
    print("시간이 너무 오래 걸림")
else: 
    t = []
    n = []

    for page in range(1, nPage):

        # curl "https://bitnodes.io/api/v1/snapshots/?limit=10&?page=1" | jq '.'
        url = 'https://bitnodes.io/api/v1/snapshots/?limit=10&?page=1'

        
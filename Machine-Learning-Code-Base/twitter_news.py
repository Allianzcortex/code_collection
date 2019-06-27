# -*- coding:utf-8 -*-
"""
This script can get the trending topic of twitter and related news.

Notice : Sometimes there is no news related with one topic because of the frequency limit of token
"""

import requests
from twitter import Twitter, OAuth


"""config value"""
twitter_access_key = "1644320072-YX7aGGIzXVgnBwM0JquHtRandomValueGenerated"
twitter_access_secret = "UejBVkIWeCN2eF2ogDFXmVqZC4DRandomValueGenerated"
twitter_consumer_key = "2vunqZWJJEtRandomValueGenerated"
twitter_consumer_secret = "bcJyPobhfrknBWoRandomValueGenerated"

news_api_key = "fef960ab-aRandomValueGenerated"


def get_twitter_trends(location_id=23424975):
    """
    get the trending topic based on location
    """
    twitter = Twitter(auth=OAuth(twitter_access_key, twitter_access_secret,
                                 twitter_consumer_key, twitter_consumer_secret))

    # this id is for UK , you can change it to any locations you want
    # TODO : add a translation tool between location name and location code
    results = twitter.trends.place(_id=location_id)

    for location in results:
        for trend in location["trends"]:
            print(
                "The trend topic is {} following is the related news :\n -------".format(trend["name"]))
            # remove # before trend topic
            print(trend["name"][1:] if trend["name"].startswith(
                "#") else trend["name"])
            get_news(trend["name"][1:] if trend["name"].startswith(
                "#") else trend["name"])


def get_news(key_word="SMU"):
    """
    Get the news based on https://eventregistry.org/documentation/examples#event-registry-api-examples-articles-get-articles
    I don't use the client insided but the restful api ,it is more customized
    """
    body = {
        "action": "getArticles",
        "keyword": key_word,
        "articlesPage": 1,
        "articlesCount": 100,
        "articlesSortBy": "date",
        "articlesSortByAsc": False,
        "articlesArticleBodyLen": -1,
        "resultType": "articles",
        "dataType": ["news", "pr"],
        "apiKey": "fef960ab-a667-4217-a42e-fa0e45fe1e53",
        "forceMaxDataTimeWindow": 31
    }

    r = requests.post(
        'http://eventregistry.org/api/v1/article/getArticles', json=body)
    for x in r.json()['articles']['results'][:5]:
        try:
            print("title is : {}".format(x["title"]))
            print("url is : {}".format(x["url"]))
            print("dateTime is : {}".format(x["dateTime"]))
            print("abstract of content is : {}".format(x["body"][:40]))
            print("--------------------------------")
        except UnicodeEncodeError as ex:
            pass


if __name__ == '__main__':
    get_twitter_trends()

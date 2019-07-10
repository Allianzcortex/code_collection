from __future__ import print_function
# -*- coding:utf-8 -*-
#!/usr/bin/env python

"""
This script is used to implement a search engine

It will be very simple and will be a MVP(Minimal Viable Product)

In a formal search engine , we should record the occurence times and display it based on the priority.
But in this case , I just remove it and return all searched result by natural order.

I build the data based on IMDB movie webpage

There is no available dataset so I must crawl it from scratch

I use NLP to remove the stop_words and store the sentences.

I implement all the codes and no copy-paste one-line from Internet.

schedule_spider() -> crawl_page() -> nlp_analyze() -> build index()
                                                           ⬆ keyword 
                                                        search
"""

import re
from collections import deque

import requests 
import nltk
from nltk.corpus import stopwords 
stop_words = set(stopwords.words('english'))


frequence_dict={}
# I prefer to use BFS search algorithm , If you want to use DFS search algorithm , you should use Stack
webpage_list=deque()
crawl_web_page_num=4


# available_link_pattern=re.compile()

content_pattern=re.compile("<meta name=\"description\" content=\"(.*?)\" />")
title_pattern=re.compile("<meta property='og:title' content=\"(.*?)\" />")
url_pattern=re.compile(" <a href=\"/title/tt(\d+)/.*\"")

def crawl_page(url):
    """
    To carawl a single webpage,get content/title/description/related urls.

    Overall It is not a good choice to use regex to parse the website,beautifulsoup/lxml is better
    In this case I will just try to make a simple demo

    @param: url webpage link
    """
    global webpage_list
    content = requests.get(url).text
    # get content
    title = title_pattern.findall(content)[0]
    description = content_pattern.findall(content)[0]
    urls = url_pattern.findall(content)
    # always get new urls and push it to list , so it can be retrieved later
    for new_url in urls[:5]:
        webpage_list.append("https://www.imdb.com/title/tt{}/".format(new_url))
    nlp_analyze(description,title)


def nlp_analyze(content,title):
    """
    Parse the file description
    """
    tags=nltk.word_tokenize(content)
    build_index([x for x in tags if x not in stop_words and len(x)>=3],title)


def schdule_spider(initial_url):
    """
    Decide how many pages a crawler should reach
    @initial_url : seed url
    """
    global webpage_list
    webpage_list.append(initial_url)
    index = 0
    while True:
        # break when exceeding the crawl amount
        if(index>crawl_web_page_num):
            break
        print("-------------Begin to Process {} webpage-------------".format(index+1))
        # get the newest url
        crawl_page(webpage_list.popleft())
        index += 1


def build_index(tags,title):
    """
    Build the search engine based on inverted index technology
    """
    global frequence_dict
    for x in tags:
        if x not in frequence_dict:
            frequence_dict[x]=[title]
        else:
            frequence_dict[x].append(title)


def search():
    """
    search based on user input
    """
    print("Get Search Result is :")
    print(frequence_dict)
    print("-------------------------------------------------------")

    while True:
        keyword=input("Please enter the keyword , press q to quit\n")
        if(keyword=='q'):
            break
        else:
            if(keyword in frequence_dict):
                print(frequence_dict[keyword])
            else:
                print("No Search result")

if __name__=='__main__':
    # Change the following url and you can get a different result
    # I use this link because <The Rock> is the first action movie I watched (:
    schdule_spider("https://www.imdb.com/title/tt0117500/")
    search()
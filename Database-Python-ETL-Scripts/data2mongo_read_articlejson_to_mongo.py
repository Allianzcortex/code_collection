# -*- coding:utf-8 -*-

from __future__ import print_function

import sys

import json
import pymongo
import traceback

def create_mongo_connection(mongo_database,mongo_collection):
    """create connection with mongodb
    """

    myclient = pymongo.MongoClient("mongodb://localhost:27017/")

    mydb = myclient[mongo_database]
    # if(mongo_collection in mydb.list_collection_names()):
    #     mydb[mongo_collection].drop()
    mycol = mydb[mongo_collection]

    return mycol

def read_file(collection):

    except_count=0

    """{"@mdate": "2011-01-11", 
    "@key": "journals/acta/Saxena96",
     "author": {"ftail": "\n", "ftext": "Sanjeev Saxena"}, 
    "title": {"ftail": "\n", "ftext": "Parallel Integer Sorting and Simulation Amongst CRCW Models."}, 
    "pages": {"ftail": "\n", "ftext": "607-619"}, 
    "year": {"ftail": "\n", "ftext": "1996"}, 
    "volume": {"ftail": "\n", "ftext": "33"}, 
    "journal": {"ftail": "\n", "ftext": "Acta Inf."}, 
    "number": {"ftail": "\n", "ftext": "7"}, 
    "url": {"ftail": "\n", "ftext": "db/journals/acta/acta33.htmlfSaxena96"},
     "ee": {"ftail": "\n", "ftext": "http://dx.doi.org/10.1007/BF03036466"}, 
    "ftail": "\n", "ftext": "\n"}"""
    
    with open("articles.json") as input_:
        _doc_index=0
        for line in input_:
            try:
                s = json.loads(line)
            except ValueError:
                # ignore 28087 line , it is "@mdate": "2015-06-18"
                pass
            try:
                # how to get all authors name
                if("author" in s):
                    name_str=s["author"]
                else:
                    name_str=s["editor"]
                # judge whether it is dict
                if(isinstance(name_str,list)):
                    name_str=name_str[0]
                authors=([values.encode('ascii','ignore').decode('utf-8') for values in name_str.values() if values.encode('ascii','ignore').strip()!=""])
            except KeyError:
                authors=" "
            try:
                    if("ftext" in s["title"]):
                        title=s["title"]["ftext"].encode('ascii','ignore')
                    else:
                        try:
                            title=s["title"]["i"]+''.join(x.encode("ascii","ignore") for x in s["title"]["sup"].values())
                        except (TypeError):
                            title=b"except_count"
                        except (AttributeError):
                            title=b"except_count"
            except KeyError:
                try:
                    title=s.values()[0]["ftext"]
                except TypeError:
                    title = b"except_count"
                
            try:
                pages=s["pages"]["ftext"].encode('ascii','ignore')
            except KeyError:
                pages=b"-1"
            try:
                journal=s["journal"]["ftext"].encode('ascii','ignore')
            except KeyError:
                journal=b""
            try:
                if(title==b"except_count"):
                    except_count+=1
                
                article_item = {
                "_id":_doc_index,
                "title":title.decode('utf-8'),
                "journal":journal.decode('utf-8'),
                "volume":s["volume"]["ftext"].encode('ascii','ignore').decode('utf-8'),
                "year":s["year"]["ftext"].encode('ascii','ignore').decode('utf-8'),
                "pages":pages.decode('utf-8'),
                "authors":authors
                }
               
                x=collection.insert_one(article_item)
                print("successful insert {}".format(x.inserted_id))
            except KeyError as ex:
                traceback.print_exc()
                print("------")
                print(line)
                exit(1)
            _doc_index+=1
    
    # print("Total except count is "+str(except_count))



if __name__=='__main__':
    mongo_database = sys.argv[1]
    mongo_collection = "ARTICLE"
    mongo_collection=create_mongo_connection(mongo_database,mongo_collection)
    read_file(mongo_collection)
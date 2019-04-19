# -*- coding:utf-8 -*-

"""
This script will read data from database AUTHOR table
and create corresponding collection to mongodb

"""
from __future__ import print_function

import sys

import pymongo
import pymysql

def create_mongo_connection(mongo_database,mongo_collection):
    """create connection with mongodb
    """

    myclient = pymongo.MongoClient("mongodb://localhost:27017/")

    mydb = myclient[mongo_database]
    # if(mongo_collection in mydb.list_collection_names()):
    #     mydb[mongo_collection].drop()
    mycol = mydb[mongo_collection]

    return mycol


def get_mysql_data(mongo_collection,database,database_username,database_password):
    conn = pymysql.connect(host='localhost', port=3306, user=database_username, passwd=database_password, db=database)

    cur = conn.cursor()
    cur.execute("SELECT * FROM AUTHOR")
    _index_id=0
    for row in cur:
    
        # get sql data
        author={
            "_id":_index_id,
            "id":row[0],
            "fname":row[1],
            "lname":row[2],
            "email":row[3]
        }
        _index_id+=1
        x = mongo_collection.insert_one(author)
        print("successful insert {}".format(x.inserted_id))

    cur.close()
    conn.close()

if __name__=='__main__':
    mongo_database = sys.argv[1]
    mongo_collection = "AUTHOR"
    database = sys.argv[2]
    database_username = sys.argv[3]
    database_password = sys.argv[4]
    mongo_collection=create_mongo_connection(mongo_database,mongo_collection)
    get_mysql_data(mongo_collection,database,database_username,database_password)

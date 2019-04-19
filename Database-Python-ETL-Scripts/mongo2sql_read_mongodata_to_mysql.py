# -*- coding:utf-8 -*-
"""
This script will read data from mongodb collection `AUTHOR` and 
`ARTICLE`;

"""
from __future__ import print_function

import sys
import time

import pymongo
import pymysql

def create_mongo_connection(mongo_database,mongo_collection):
    """create connection with mongodb
    """

    myclient = pymongo.MongoClient("mongodb://localhost:27017/")
    mydb = myclient[mongo_database]
    mycol = mydb[mongo_collection]
    # if(mongo_collection in mydb.list_collection_names()):
    #     mydb[mongo_collection].drop()
    return mycol

def write_collection_to_mysql(mongo_database="h_wang",database_username='root',database_password='root',database="database_project"):
    print("begin to write magazine information into mysql table: ")
    time.sleep(1)

    articles = create_mongo_connection(mongo_database,"ARTICLE")
    # skip(author.count() - 2000):
    # find({},{'journal':1})
    magazine_set = set()
    author_set = set()
    conn = pymysql.connect(host='localhost', port=3306, user=database_username, passwd=database_password, db=database)
    cur = conn.cursor()

    index = 1
    for document in articles.find():
        journal = document['journal']
        if(journal!=""):
            magazine_set.add(journal)
        author = document['authors']
        author = ''.join(x.strip() for x in author )
        if(author!=""):
            author_set.add(author)
        author_set.add(author)
    
        query = ("insert into ARTICLE(_id,title,magazine,volume_number,year,pages,authors) " + 
        "VALUES(%s, %s, %s, %s, %s, %s, %s)")
        cur.execute("insert into ARTICLE(_id,title,magazine,volume_number,year,pages,authors) VALUES(%s, %s, %s, %s, %s, %s, %s)",(document["_id"],document["title"],
             document["journal"],document["volume"],document["year"],document["pages"],"".join(x.strip() for x in document["authors"])))
        
        print("insert article table: "+str(index))
        index += 1
    conn.commit()
    time.sleep(1)

    index = 1
    for x in magazine_set:
        print("insert magazine information {} ".format(index))
        try:
            cur.execute("insert into MAGAZINE(name) values (\"{}\");".format(x)) 
        except Exception:
            pass
        index += 1
    conn.commit()
    print("execute done")
    time.sleep(1)

    print("begin to write author information into mysql table: ")
    time.sleep(1)
    index = 1
    for x in author_set:
        print("insert author {} ".format(index))
        try:
            cur.execute("insert into AUTHOR(lname) values (\"{}\");".format(x)) 
        except Exception:
            pass
        index += 1
    conn.commit()
    print("execute done")

if __name__ == '__main__':
    mongo_database = sys.argv[1]
    database = sys.argv[2]
    database_username = sys.argv[3]
    database_password = sys.argv[4]
    write_collection_to_mysql(mongo_database,database_username,database_password,database)
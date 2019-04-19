#!/bin/bash
# execute mysql scripts first

if [ $# -ne 4 ]
then
    printf "\nUsage: \n\$ $0 <mysql_database> <mysql_user> <mysql_password> <mongo_database>   \n\n"
    exit 1
fi

mysql_databse="$1"
mysql_user="$2"
mysql_password="$3"
mongo_database="$4"

printf "Begin to execute mysql scripts: "
sleep 1
# TODO 
mysql -u $mysql_user -p$mysql_password $mysql_databse < "existing_tables.sql"
mysql -u $mysql_user -p$mysql_password $mysql_databse < "new_tables.sql"

printf "Begin to read mysqldata to mongo: "
sleep 1
pip install pymongo
pip install PyMySQL
mongo $mongo_database --eval "db.AUTHOR.drop()"
cd "$(dirname "$0")" && python data2mongo_read_mysqldata_to_mongo.py $mongo_database  $mysql_databse $mysql_user $mysql_password

printf "Begin to read json file to mongo: "
sleep 1
mongo $mongo_database --eval "db.ARTICLE.drop()"
cd "$(dirname "$0")" && python data2mongo_read_articlejson_to_mongo.py $mongo_database
printf "Done\n"


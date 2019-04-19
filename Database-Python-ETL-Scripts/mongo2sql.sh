#!/bin/bash

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
pip install pymongo
pip install PyMySQL
cd "$(dirname "$0")" && python mongo2sql_read_mongodata_to_mysql.py $mongo_database  $mysql_databse $mysql_user $mysql_password
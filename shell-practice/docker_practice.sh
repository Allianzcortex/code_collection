#!/bin/bash

# currently part of final codes

for (( c=2001;c<=80000; c++ ))
do
   resp=`netstat -tunl | grep ":$c "`
   if [ -z "$resp" ]; then
    echo "$c Port is free"
    avaialble_port=$c
    break
  fi
done
echo "avaialble port is $avaialble_port"


# 8
# https://www.shellhacks.com/check-website-availability-linux-command-line/

    echo "web page is right"
else
    echo "no webpage"
    # exit and clean todo
fi

# 9

# 10
# TODO

# 11



# docker stop first
docker ps | grep your_user_name | awk '{print $1}' | xargs docker stop

docker rm container_name
docker rmi image_name
exit 0
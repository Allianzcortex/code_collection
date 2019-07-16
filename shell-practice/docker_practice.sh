#!/bin/bash

# 1,2,3,6 are exit signals
trap clean_docker 1 2 3 6

show_progress()
{   
    echo "This is task $1 *****************"
    sleep 2
}

user="######"

# 1 & 2 remove the folder if it already exists
show_progress 1and2
if [ -d "out" ]; then rm -Rf out; fi
mkdir out && cd out

# 3 get the target file
show_progress 3
wget /docker/Dockerfile
wget /docker/app.py

# 4 get the target string based on date
show_progress 4
current_day=$(date '+%d')
if (( $current_day % 2 == 0 ))
then
    new_text='Today is an even day'
else
    new_text='Today is an odd day'
fi

echo "The text get is : $new_text"

# replace 
sed -i "s/Hello World!/$new_text/g" app.py

# 5 & 6
show_progress 5and6

docker_image="$user"_a2

# get available port;max port in Linux is 65535
for (( c=2000;c<=65535; c++ ))
do
   resp=`netstat -tunl | grep ":$c "`
   if [ -z "$resp" ]; then
    echo "$c port is available"
    available_port=$c
    break
   else 
    echo "$c port is already used"
  fi
done

# build docker image
docker build -t $docker_image .
docker run -d  -p $available_port:80 $docker_image

# 7
show_progress 7

FS=' ' read -r -a array <<< "$(docker ps | grep $docker_image)"
container_id=${array[0]}
container_ip_address=$(docker inspect $container_id | grep "\"IPAddress\"" | head -1 | sed 's/,//g' | cut -d":" -f2 )

echo "\"Container\'s IP is $container_ip_address"

# very important,or docker image may not be built
sleep 5

# define variables and utils functions used 
available_url="lnx.cs.smu.ca:"$available_port
clean_docker()
{
    docker stop $container_id
    docker rm $container_id
    docker rmi $docker_image
}

# 8 check https://www.shellhacks.com/check-website-availability-linux-command-line/ for more details
# 200 is the HTTP status code that represents the function-well page
show_progress 8
if [[ $(curl -Is $available_url| head -n 1 | grep 200) ]];then
    echo "web page works correctly"
else
    echo "web page doesn't work correctly , will remove the docker container and clean it"
    # exit and clean 
    clean_docker
    exit 1
fi

# 9 save webpage 
show_progress 9
echo "available url is : "
echo $available_url

wget $available_url -O serv.html

# 10 


ssh -R $available_port:localhost:$available_port $user@dev.cs.smu.ca "wget -t 1 smu.ca localhost:$available_port" 

# 11 clean the docker
show_progress 11
clean_docker

exit 0
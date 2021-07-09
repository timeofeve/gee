#!/bin/bash


# curl -XGET '192.168.10.209:9200/_cat/indices?v&pretty' |grep open|grep -v ".mon\|.ki\|trnseqno\|dgsx\|chains" |sort|awk NR!=1'{print $3}' > list.txt

for line in $(cat list.txt)
do  
    # aa=chains
    # echo $aa
    echo '{"source": {"index": "'$aa'","remote": {"host": "http://192.168.10.209:9200","username": "","password": ""}}'
    # echo `curl -i -X POST -H "Content-type:application/json" -d {"source": {"index": "${line}","remote": {"host": "http://192.168.10.209:9200","username": "","password": ""}},"dest": {"index": "${line}"}} http://192.168.10.207:9200/_reindex`
    curl -i -X POST -H "Content-type:application/json" -d '{"source": {"index": "'$line'","remote": {"host": "http://192.168.10.209:9200","username": "","password": ""}},"dest": {"index": "'$line'"}}' http://192.168.10.207:9200/_reindex
    # exit 1
done


for line in $(ls ./);do echo "mysql -uroot -pzjhz2017 -h192.168.10.210 < $line";mysql -uroot -pzjhz2017 -h192.168.10.210 < "$line"; done;
mysql -uroot -pzjhz2017 -h192.168.10.210 < zj.sql
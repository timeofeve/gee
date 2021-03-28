#!/bin/bash

# 获取列
list_num=`sed -n '1p' file.txt | wc -w`
# 逐列转置输出
for((i=1;i<=$list_num;i++))
do
    line=`cat file.txt | awk '{print $'$i'}' | tr -s "\n" " "`
    # echo $line
done

# cat file.txt | tr -s ' ' '\n'|sort|uniq -c|sort -r|awk '{print $2" "$1}'


awk '{
    for(i = 1; i <= NF; i++){
        res[$i] += 1 #以字符串为索引，res[$i]相同的累计
    }
}
END{
    for(k in res){
        print k" "res[k]
    }
}' file.txt | sort -nr -k2 
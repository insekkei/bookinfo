echo -e "请输入你要读取的文件夹路径\n当前路径为${PWD}"  
read InputDir  
echo "你输入的文件夹路径为${InputDir}"  
echo -e "请输入你要将数据输出保存的文件路径n当前路径为${PWD}"  
read OutputFile    
echo "输出保存的文件路径为${OutputFile}"  
: > $OutputFile   #清空OutputFile  
#循环读取文件夹名  
for file_a in ${InputDir}/*; do  
    temp_file=`basename $file_a`  
    echo $temp_file >> $OutputFile  
done  
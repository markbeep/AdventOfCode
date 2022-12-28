 for loop
time=0
n=100
for ((i=0; i<n; i++));
do
 \time -o res.txt -f "%e" node day5.js
 time+=$(cat res.txt)
done

echo $time
echo $n
echo $time/$n

# Read from the file words.txt and output the word frequency list to stdout.
cat words.txt | awk '{ for (i = 1; i <= NF; i++) { print $i } }' | sort | uniq -c | sort -r | awk '{ print $2, $1 }'

cat machines.json | jq ".[]" | grep techniques | awk -F ":" '{print $2}' | sed 's/\\n/\n/g' | tr -d '[:punct:]' | grep -o "[[:alpha:]]*" | sort | uniq -c | sort --numeric-sort

#!/bin/bash

count=0
while [ "$count" -lt "30" ];
do
    status=`kubectl get pods | grep web- | cut -d ' ' -f 9`
    if  [ $status == "Running" ]; then
        echo "POD is up and Running"
        break
    else
        count=`expr $count + 1`
        sleep 1
    fi
    echo $status
done

if [ $count -eq 30 ]; then
    echo "Timeout. Please check your CI configuration"
    exit 1
fi
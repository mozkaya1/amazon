#!/bin/env sh 

## Set alert price threhold. If price below that threhold alarm will be triggered. -- Only Euro. Not cent --
setPrice=50

## Set Amazon url to be watching -- 
url="https://www.amazon.com/TwinGrip-Pliers-Comfort-Grip-8-inch/dp/B097C7W2YK"

## Set reasonable interval for price check 
cooldown=3600

## Infinite loop for checking price ---
while true ;do
timestamp=$(date '+%Y-%m-%d %H:%M:%S')

## Log writer function for better output
log_write() {
    local message="$1"
    local message2="$2"
    echo "[$timestamp] -- $message -- $message2" 
}

## Get Price of the product with Go Code -- 
output=`go run main.go -set $setPrice -url ${url}`


## split variables from output
status=`echo "${output}"|jq -r .status`
price=`echo "${output}"|jq -r .price`
url=`echo "${output}"|jq -r .url`
discount=`printf "%.0f\n" $(echo ${output}|jq -r .discount)`


if [[ $status == 200 && $discount != 0 ]]; then

  message="$timestamp - Amazon Discount Alert - % $discount -- $url" 
  setPrice=$price

  #send Telegram Notification with Python Code
  python /home/musti/python-project-server/telegram_v3.py "${message}" >/dev/null

  ## send mail notification 
  # echo "${message}" | mail -s "Amazon Discount Alert"  -r RECEIVER@hotmail.com
  log_write "Price dropped - % ${discount} --> ${url}"|tee -a amazon.log
  continue
else
log_write "No Discount" ${output}|tee -a amazon.log
fi

sleep $cooldown
done

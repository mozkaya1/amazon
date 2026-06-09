# Get Price Discount Notification for Specific Product on Amazon.com/nl/es  ..

### Code 
[https://github.com/mozkaya1/amazon](https://github.com/mozkaya1/amazon) 

### Source Page
> get amazon.com/nl/es link from product page .

![amazon](./static/pic-selected-260609-1941-40.png) 

### initial setting on shell code (./amazon.sh)
> adjust settings on ./amazon.sh
```bash
set "url" link for product on Amazon 
set "setPrice" for threshold price to be notified below  
set "cooldown"  check interval (seconds)
```
![initial setting](./static/pic-selected-260609-1914-54.png) 


### main GO code to get price on Amazon

> help context 
```bash
go run main.go -h
```

```bash
[musti@musti-deputyp25 amazon]$ go run main.go -h
Usage of /home/musti/.cache/go-build/8f/8fc66cf0ce440f5888751dbcc8d2b1bb42bcea7ce5ee2d1d5a4c243277035125-d/main:
  -set float
        Set Alarm Price Threshold, expected below this threshold (default 10)
  -url string
        Set url to be Checked (default "https://www.amazon.com/TwinGrip-Pliers-Comfort-Grip-8-inch/dp/B097C7W2YK")

```
![go code help](./static/pic-selected-260609-1918-12.png) 

> run raw code  -- only gives output json of the product such as status,price,url ..
```bash
go run main.go

```
```bash
{"url":"https://www.amazon.com/TwinGrip-Pliers-Comfort-Grip-8-inch/dp/B097C7W2YK","price":36,"status":200,"discount":0}
```

### Notification -- (./amazon.sh)
> I have used my own Telegram bot. You can use your own notification system as mail,notify so on ..

![ncode](./static/pic-selected-260609-1929-06.png)


> Preview - Telegram

![telegram](./static/pic-selected-260609-2109-13.png) 

### output log (./amazon.log)

```bash 
[musti@musti-deputyp25 amazon]$ ./amazon.sh 
[2026-06-09 19:13:15] -- No Discount -- {"url":"https://www.amazon.com/TwinGrip-Pliers-Comfort-Grip-8-inch/dp/B097C7W2YK","price":36,"status":200,"discount":0}
[2026-06-09 19:13:27] -- No Discount -- {"url":"https://www.amazon.com/TwinGrip-Pliers-Comfort-Grip-8-inch/dp/B097C7W2YK","price":36,"status":200,"discount":0}
[2026-06-09 19:13:39] -- No Discount -- {"url":"https://www.amazon.com/TwinGrip-Pliers-Comfort-Grip-8-inch/dp/B097C7W2YK","price":36,"status":200,"discount":0}
[2026-06-09 19:13:50] -- No Discount -- {"url":"https://www.amazon.com/TwinGrip-Pliers-Comfort-Grip-8-inch/dp/B097C7W2YK","price":36,"status":200,"discount":0}
[2026-06-09 19:14:02] -- No Discount -- {"url":"https://www.amazon.com/TwinGrip-Pliers-Comfort-Grip-8-inch/dp/B097C7W2YK","price":36,"status":200,"discount":0}
```
![output-log](./static/pic-selected-260609-2107-36.png) 
 

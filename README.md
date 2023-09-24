# stonks

Stonks is a Stock Market Data Processing Engine which leverages the power of cloud native tools like Prometheus and Grafana to provide an in house solution for tracking stock market data

## Architecture

![image](https://github.com/yrs147/stonks/assets/98258627/640d0a0a-2b53-4afe-9ae5-8994962a893f)

## Prerequites (To be installed on the Servers)

1. Golang
2. Docker
3. Docker-Compose 

## Steps to Setup the Project

1. Now Setup a `t2.medium`EC2 Instance named `Kafka`
2. 
3. Connect to the server using `EC2 Instance Connect` or through a `ssh client` (**using the pem**)and run the following command :-

```
sudo apt-get update
```

4. Install the prerequisites on all servers 

5. Now clone the project on all the servers using the following command :- 
``` 
git clone https://github.com/yrs147/stonks
```
6. Go the the kafka server and run the using the following commands :-

```
cd stonks
```
then
```
docker-compose -f kafka.yml -d up
```

7. Now to to the producer servers and run the following command 



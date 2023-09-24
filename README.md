# stonks

Stonks is a Stock Market Data Processing Engine which leverages the power of cloud native tools like Prometheus and Grafana to provide an in house solution for tracking stock market data

## Architecture

![image](https://github.com/yrs147/stonks/assets/98258627/640d0a0a-2b53-4afe-9ae5-8994962a893f)

## Prerequites (To be installed on the Servers)

1. Golang
2. Docker
3. Docker-Compose 

## Steps to Setup the Project

1. Install the prerequisites on your device

2. Now clone the project on all the servers using the following command :- 
``` 
git clone https://github.com/yrs147/stonks
```

3. Open project directory :-
```
cd stonks
```

4. Using an editor open the `docker-compose.yaml` file and add the links of stocks you want to track in `STOCK1` and `STOCK2` under `producer1` and `producer2`

5. Run the following command to start the project
```
docker-compose -f kafka.yml -d up
```



# stonks

## Project Description
![letter_C_trade_marketing_logo_design_template-removebg-preview-removebg-preview](https://github.com/yrs147/stonks/assets/98258627/0f6ce072-963f-4237-863e-301e312d52a6)

In a world where access to real-time stock market data is paramount for making informed investment decisions, traditional stock market trackers often come with exorbitant subscription fees. These fees can place a significant financial burden on both individual investors and businesses alike. Recognizing this issue, I've created `Stonks`  a `Stock Market Data Processing Engine`. It leverages the power of cloud-native tools like Prometheus and Grafana to provide an in-house solution for tracking stock market data, offering an innovative and cost-effective alternative to expensive commercial tracking software.



## Problem 
Stock market trackers are essential tools for investors, financial analysts, and traders. They provide critical information, insights, and analytics to facilitate better investment strategies and decisions. However, many of these trackers come with steep subscription costs, making them inaccessible to a broad range of potential users. The cumulative expenses of subscribing to multiple tracking services can quickly become a significant financial burden, especially for smaller investors and businesses. This creates a substantial barrier to entry for those seeking to gain a foothold in the financial markets.






## Key Features of Stonks:

1. **Real-Time Data** : Stonks provides users with real-time stock market data with help of a custom Prometheus Exporter, ensuring that they can make timely investment decisions.

2. **Custom Dashboards** : Stonks allows users to create custom dashboards using Grafana, tailoring the information they receive to their specific needs.

3. **Cost-Effective** : As an open-source solution, Stonks drastically reduces the cost of tracking stock market data, making it accessible to a broader user base.

4. **Data Control** : With an in-house solution, users have full control over their data, ensuring data security and privacy.


## Technologies Used
1. Golang
2. Apache Kafka
3. Prometheus
4. Grafana
5. MongoDB
6. Docker-Compose

## Architecture

![Screenshot 2023-10-13 151217](https://github.com/yrs147/stonks/assets/98258627/f55e0932-d7e0-4467-94bd-801c4c394ca1)



## Prerequites

1. Docker
2. Docker-Compose
3. A good internet connection

## Screenshots

![Screenshot 2023-10-13 at 1 45 16 PM](https://github.com/yrs147/stonks/assets/98258627/7122c2e7-5181-498e-b846-ba748f93db0b)

![Screenshot 2023-10-13 at 2 41 44 PM](https://github.com/yrs147/stonks/assets/98258627/6f9efed8-b670-4e83-b4db-9e2b98455fa5)


![Screenshot 2023-10-13 at 2 38 43 PM](https://github.com/yrs147/stonks/assets/98258627/70bc539d-2aeb-4cfc-b4db-494cdcda1cb2)


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

4. Using an editor open the `dock.yaml` file and add the URLs of stocks you want to track in `STOCK1` and `STOCK2` under `producer1` and `producer2` (**Remember that URLs should only be from "in.investing.com" due to scraping capabilities**)

5. Run the following command to start the project and wait for 2 min 
```
docker-compose -f dock.yml -d up
```
6. Then open `localhost:3000` in your browser and login to Grafana with username `admin` and password `admin` (Credentials can be changes in the settings)

![Screenshot 2023-10-13 at 1 46 26 PM](https://github.com/yrs147/stonks/assets/98258627/6628fb06-a5a4-4d0b-b98f-6dd8d4e6c7bf)

7. Then Go to **Add Your First Data Source** and choose **Prometheus**

![Screenshot 2023-10-13 at 1 47 09 PM](https://github.com/yrs147/stonks/assets/98258627/8133061a-e875-4bdc-b257-97dafef70eec)

8. Then add `http://prometheus:9090` in the `Prometheus server URL` 

![Screenshot 2023-10-13 at 2 04 13 PM](https://github.com/yrs147/stonks/assets/98258627/6477192c-e88e-4f96-8047-66a5125d6a7e)

and save it

![Screenshot 2023-10-13 at 1 51 04 PM](https://github.com/yrs147/stonks/assets/98258627/bf8ccc55-059e-4020-aa77-dae230cacb27)

10. Then Navigate Back to home and click on `Create Your First Dashboard`

11. Click on `Add visualization` and choose `Prometheus` as your data source

![Screenshot 2023-10-13 at 1 51 53 PM](https://github.com/yrs147/stonks/assets/98258627/a2ae293e-e6ba-4b4b-9de7-734c324f118e)

12.  Add `stock_close` in the **metrics** field , choose `name` under **labels** and the stock you want to analyze as its value and click on `Run queries`

![Screenshot 2023-10-13 at 1 52 55 PM](https://github.com/yrs147/stonks/assets/98258627/f9d80a70-3e21-41c6-9cb1-108900aa882f)

13.  Now that you can see the stock being Tracked , click on `Apply` to save the button

14.  Repeat These Steps for the other stocks as well

15.  Once you have your dashboard set up then select the refresh time as per your requirements

![Screenshot 2023-10-13 at 1 54 03 PM](https://github.com/yrs147/stonks/assets/98258627/ee7afa6d-ffb8-40e2-8f39-52edefcb8363)

16.  And congrats!! you've successfully set up the project

## Future Prosepects 

1. Adding alerting support through which users can get real time notifications on their devices through slack , email and discord
2. Adding Support to track more than 2 stocks
    
 


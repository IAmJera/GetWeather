# GetWeather
This program gets the weather from https://weatherapi.com/
site and enters the temperature in Celsius in the weather.txt file.

weatherapi.com only allows you to send requests to its API 1 million times a month for free. 
Therefore, the program receives currency values every 10 minutes.

__By default,  file is stored in the data directory.__

__You may change the `volumes` section in the `docker-compose.yaml` file.__
## Getting started
1. Clone this repository
2. Add api key and city to `config.env` file
3. Run `docker-compose up -d --build` to start the project

## Environment variables
| Name     | Description        | Example                            |
|----------|--------------------|------------------------------------|
| `APP_ID` | weatherapi api key | `N8DKXqKPP9t7ptmmFNxWireHpmtUxK8e` |
| `CITY`   | city name          | `London`                           |

### Example output:
```
22°
```
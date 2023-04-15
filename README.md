## Warning 
Sometimes "depends on" in docker-compose.yml does not work causing bot to running before postgres container that is why sometimes "connection refused" error occures, 
```
docker start weather_bot
```
solves the problem
## Weather bot written in Golang
    This weather bot is a telegram bot tool that provides users with up-to-date weather information for a given location.

# Running
    To run the weather bot, you will need to have Docker installed on your machine. Once you have Docker installed,
    you can use the following commands to build and run the bot:
```
docker compose build
```
```
docker compose up
```

## Usage
To use the weather bot, simply run send the name of the location you want to get weather information for.

For example:
    London

This will return the current temperature for London.

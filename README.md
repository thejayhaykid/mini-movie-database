# Mini movie database

Read [this article](https://towardsdatascience.com/step-by-step-guide-to-build-your-own-mini-imdb-database-fc39af27d21b) about how to write a scraper in Python that will create a personal movie database and decided to create it in Go.

I have also decided to implement this as an API that I'm setting up for my local network only as part of my home intranet. 

You will need to [create an account at themoviedb.org](https://www.themoviedb.org/account/signup) and [get an API key](https://developers.themoviedb.org/3/getting-started/introduction). The take that key and create a `keys.json` file with the following:

```json
{
    "api-key": "YOURAPIKEY",
    "sql-db": {
        "host": "localhost",
        "port": "3306",
        "password": "YOURPASSWORDHERE",
        "dbname": "moviedb"
    },
    "mongo-db": {
        "host": "localhost",
        "password": "YOURPASSWORDHERE",
        "port": "8080"
    },
    "host": "localhost"
}
```
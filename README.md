# Mini movie database

Read [this article](https://towardsdatascience.com/step-by-step-guide-to-build-your-own-mini-imdb-database-fc39af27d21b) about how to write a scraper in Python that will create a personal movie database and decided to create it in Go.

You will need to [create an account at themoviedb.org](https://www.themoviedb.org/account/signup) and [get an API key](https://developers.themoviedb.org/3/getting-started/introduction). The take that key and create a `keys.json` file with the following:

```
{
    "api-key": "YOUR API KEY HERE"
}
```

*Note: There is already an [awesome Go wrapper for themoviedb.org](https://github.com/ryanbradynd05/go-tmdb/) that I have not used. This project is more for me learning how to do it than anything else so I did not use that.*
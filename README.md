# Sauce
A web scraping Discord bot that extracts developer-related articles from select developer
blog websites

## Setup
1. Install Go. See [here](https://golang.org/doc/install) for instructions.
2. Clone this repository.
3. Install dependencies.
    ```bash
    go mod download
    ```
4. Create a `.env` file from the `.env.dist` template and provide the necessary tokens.
5. If the bot is to be run locally, create a Discord server and invite the bot to it.
   [Here](https://www.ionos.com/digitalguide/server/know-how/creating-discord-bot/) are
   instructions on how to create a Discord bot and invite it to a server.
6. Once the bot is set up with the tokens provided in the respective `.env` file,
   run the program.
    ```bash
    go run main.go
    ```
   
Note: This is an ongoing project. Some features are yet to be implemented and
functionality is limited for the time being. The application is also only capable of
running in a local environment for now

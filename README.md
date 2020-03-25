# Random Useless Facts (RUF)

Discord which send random useless facts using https://uselessfacts.jsph.pl/

## Usage

Invite the bot by clicking on [this link](https://discordapp.com/oauth2/authorize?client_id=692343582727602197&permissions=0&scope=bot "Invire RUF")

On your discord server you'll be able to use the three commands `°help`, `°random` and `°today`

- `°help` → Get a help message from the bot
- `°random` or `°fact` → Get a random useless fact (A new one each request)
- `°today` → Get useless fact of today (Updates every 24 hours)

## Implementation

The bot use [Random Useless Facts](https://uselessfacts.jsph.pl/ "Random Useless Facts API") API.

The bot have two commands:

- `°random` or `°fact` → call the url [https://uselessfacts.jsph.pl/random.json?language=en](https://uselessfacts.jsph.pl/random.json?language=en)
- `°today` → call the url [https://uselessfacts.jsph.pl/today.json?language=en](https://uselessfacts.jsph.pl/today.json?language=en)

In the documentation of the api you can choose the language.  
Here, the bot make all his requests with the `language=en` query parameter.

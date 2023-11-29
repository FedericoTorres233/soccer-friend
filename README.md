# Soccer friend

Soccer friend is a telegram bot that keeps you up to date with soccer, written in Go. Uses the REST API from allsportsapi.com. My own bot is called [`@soccer_friend_bot`](https://t.me/soccer_friend_bot). Check its features by sending `/start` to it.

## Getting started

### Installing locally

1. Install go at [go.dev](https://go.dev/dl/)

2. Clone repository
```bash
git clone https://github.com/FedericoTorres233/soccer-friend && cd soccerfriend
```

3. Create a `.env` file in the project folder and add the following environmental variables:
* `APITOKEN`="Your api token from allsportsapi.com"
* `TOKEN`="Your telegram bot's token"

4. Run the following command
```bash
make run
```
5. Enjoy!

### Install using Docker

1. Clone repository
```bash
git clone https://github.com/FedericoTorres233/soccer-friend && cd soccerfriend
```

2. Set the environmental variables in the Dockerfile

3. Run the following commands:
```bash
docker build -t soccer-friend:latest .
docker run soccer-friend:latest
```

4. Enjoy!

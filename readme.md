#  LYTDT

Last year this day's tweet

## Prerequisite

Twitter developer account and keys

## How it works

1. Setup the account 

- copy `.env_template` to `.env`
- Add your twitter access token and consumer keys

2. Setup [developer sandbox](https://developer.twitter.com/en/account/environments) for `Search Tweets Full Archive`, with value `searchText`

2. Complie the code

```
go build main.go
```

2. Setup the Account

```
./main add <username>
```

> This will help you list the tweets next time without hassle

3. List the tweets from last year

```
./main list
```

> Currently it lists all tweets with more than 10 likes
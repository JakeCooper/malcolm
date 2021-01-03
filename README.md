# Malcolm

Malcolm is a very small, dynamic reverse proxy backed by Redis and written in Go

You tell Malcolm a mapping, it'll store it and a proxy'd connection

## Quickstart

1. Git clone
2. make build
3. REDISHOST=x REDISPORT=y ./malcolm (Optionally provide REDISPASSWORD)

Malcom is now running on port 1337 (Or PORT if you pass it in)

## API

GET /rule?id=x

* Returns all rules, or a specific rule if query id=x is set

POST /rule

* Create a rule

PUT /rule?id=x

DELETE /rule?id=x

* Delete a rule by id

## FAQ

* Q: Why did you make this?
* A: We use an industry standard proxy at Railway. The control plane/config language leaves to be desired IMO. Wanted to see if I could whip one up in a few hours. This is the result.

* Q: Is this production ready?
* A: Absolutely not

* Q: Why is it called Malcolm?
* A: https://en.wikipedia.org/wiki/Malcolm_in_the_Middle

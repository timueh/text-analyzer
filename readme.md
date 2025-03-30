# Analyze text as you type

This proof-of-concept project allows you to enter text that is being analyzed while you type.

![An demo of the running frontend](img.png 'Frontend demo')

If you want to spin up the application yourself, here's what you'll have to do:

- Clone the repository.
- Run `make up`.
- Visit `http://localhost` in the browser of your choice.

This project came about as an idea to combine a few things:

- Implement a rudimentary backend in `go`.
- Toy around with interfaces in `go`.
- Implement a rudimentary frontend with React using `typescript`.
- Dockerize the entire application.
- Use a reverse proxy to isolate the frontend from the backend.

This project makes no claims of being productive code.
Instead, the project aims to show the swiss-army knife character that today's engineers need to be able to provide:
It's not just about figuring out the actual business logic (_How, again, do I count the number of occurrences of letters in a given string?_), but also about dockerizing the application to make it OS-agnostic, and to package it all up.
Oh, did I mention that an appealing frontend is expected too?

Hence, this project provides a minimal setup from which further explorations are possible.

Enjoy!

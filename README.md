Current Test Coverage: 62.2%

# Pokecache

Pokecache is a memory-based, high-performance, web-scale Pokemon cache.

## Usage

See the `Makefile` and/or run `make` in your terminal to see how to run Pokecache as well as run tests.

Helpful commands:

Run all tests with go's built-in race detector enabled (this can be slow):

```
make test
```

Run all tests with go's built-in race detector disabled:

```
make quicktest
```

Run linting tools:
```
make lint
```

Run the server on port `8080`

```
make server
```

Once the server is running you can run commands such as:

```
curl -s -i --json '{"id":12,"name":"spot","height":17}' -X POST http://127.0.0.1:8080/add
curl -s -i http://127.0.0.1:8080/id/12
curl -s -i http://127.0.0.1:8080/name/spot
curl -s -i  -X DELETE http://127.0.0.1:8080/id/12
```

## TODO

Remaining items to implement in order to secure VC-backing and properly disrupt the Pokemon Cache Industry:

- Runtime-configuration with go's stdlib `flag` package
- Data-validation for submitted pokemon
- Better test coverage in HTTP handlers
- Edge-cases in HTTP handlers

## Test Coverage

See `test-coverage.html` and `test-coverage.txt` for HTML and plain-text
reports of current test-coverage details. These files are automatically updated
by the `make test` and `make quicktest` targets.

### Coding Challenge Guidelines

Hello and thanks for your interest in joining `Weave`! We are excited to see what you can do.
The goal of this challenge is to assess your skills and approach to solving a problem. We want to see how you write code and how you solve problems.

If you're familiar with Golang, we'd love to see you complete this assignment in Go. That being said, we're much more interested in grading your technical ability than
your familiarity with Go, so use whichever language you feel will best show that ability.

Write your code as if it were going into production. We will be looking at the structure/architecture of your code, how you test your code, and how you document it.

### The Challenge
#### Part 1
Implement a Cache (key/value) for storing Pokemon data. The cache should have a maximum capacity and should evict Pokemon data when it reaches the maximum capacity.
The cache should be able to store and retrieve Pokemon data by name. It should be able to store the following data for each Pokemon:
- Id
- Name
- Type
- Height
- Weight
- Abilities

The cache should support the following operations: Set, Get and Delete Pokemon data.

#### Part 2
Create endpoints to expose the cache functionality as a REST API. The API should support the following operations:
1. Get Pokemon data by Id
2. Get Pokemon data by Name
3. Delete Pokemon data by Id
4. Add Pokemon data

The API should be able to handle concurrent requests and should be able to handle errors gracefully.

Use the cache implementation from Part 1 to store and retrieve Pokemon data.

### Evaluation Criteria
1. The problems are solved efficiently and correctly, the application should work as expected.
2. You demonstrate the knowledge of how to test the critical parts of your code. We don't require 100% coverage, but we want to see how you approach testing.
3. Code should be easy to read and understand, organized, and well-documented.
4. The answers to the questions are clear and concise.
5. Following the industry best practices and standards.
6. (Plus points) if you are aware of git best practices.

### Notes
Documentation is important, don't forget to include a README file with instructions on how to run your code and tests.

### Useful Links
1. Cache: https://en.wikipedia.org/wiki/Cache_replacement_policies
2. REST API: https://restfulapi.net/
3. Golang: https://golang.org/

### CodeSubmit

Please organize, design, test, and document your code as if it were
going into production - then push your changes to the master branch.

Have fun coding! 🚀

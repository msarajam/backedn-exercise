# backend-exercise

This project will serve as the basis for your coding exercise. It is a simple web service
written in Golang and supports the creation and retrieval of "apps".

An "app" is a simple entity that defines the attributes of a software application. 

Here is an example of a valid "app" in JSON format:

```json
{
	"title": "Valid App 1",
	"version": "0.0.1",
	"maintainers": [
		{
			"name": "firstmaintainer app1",
			"email": "firstmaintainer@hotmail.com"
		},
		{
			"name": "secondmaintainer app1",
			"email": "secondmaintainer@gmail.com"
		}
	],
	"company": "Random Inc.",
	"website": "https://website.com",
	"source": "https://github.com/random/repo",
	"license": "Apache-2.0",
	"description": "### Interesting Title\nSome application content, and description\n"
}
```

You have been tasked with extending this project by adding one or more features.

When you are done run `make bundle` and send us the generated file `backend-exercise.bundle`.

## Advice

This exercise is an opportunity to show off your passion and the craftsmanship of your solution. Optimize your solution for quality and reliability. If you feel your solution is missing a cool feature and you have time, have fun and add it. Make the solution your own, and show off your skills.

## Features

#### "A"

- Add the ability to submit and retrieve "apps" in YAML format while retaining JSON support.

#### "B"

- Add an endpoint to search "apps" and retrieve a list that matches the query parameters.
- Search functionality must be written from scratch.

#### "C"

- Add per-request rate limiting. 

## Submitting finished exercise

1. Commit your changes.
2. Run `make bundle`.
3. Upload `backed-exercise.bundle` to the Google Drive folder that you were invited to by Upbound.

## Questions?

Please do not hesitate to reach out to your Upbound recruiting contacts if you have any questions.

## The base project

### Required software

- make
- Go v1.12.*
    - https://golang.org/
- Dep v0.5.3
    - https://github.com/golang/dep/releases
- GolangCI-Lint
    - https://github.com/golangci/golangci-lint
- GoMock and `mockgen`
    - https://github.com/golang/mock        

### Running

```bash
make run
```

### Testing

```bash
make test
```

### Linting

```bash
make lint
```

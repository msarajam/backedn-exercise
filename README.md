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

 
## Features

#### YAML

- Add the ability to submit and retrieve "apps" in YAML format while retaining JSON support.

#### Search

- Add a new RESTful server endpoint for searching stored "apps".

#### Rate limiting

- Add per-request rate limiting. 


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


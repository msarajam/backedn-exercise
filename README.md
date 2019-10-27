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
3. Upload `backed-exercise_completed.bundle` to the Google Drive folder that you were invited to by Upbound.

## Advice

This exercise is an opportunity to show off your passion and the craftsmanship of your solution. Optimize your solution for quality and reliability. If you feel your solution is missing a cool feature and you have time, have fun and add it. Make the solution your own, and show off your skills.

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

### Run

```bash
make run
```

### Test

```bash
make test
```

### Lint

```bash
make lint
```

### Bundle

```bash
make bundle
```


### Testing the Application 
##### Send json msg :
```
curl --data '{
"title":"Valid App 1",
"version":"0.0.1",
"maintainers":[
   {
      "name":"firstmaintainer app1",
      "email":"firstmaintainer@hotmail.com"
   },
   {
      "name":"secondmaintainer app1",
      "email":"secondmaintainer@gmail.com"
   }
],
"company":"Random Inc.",
"website":"https://website.com",
"source":"https://github.com/random/repo",
"license":"Apache-2.0",
"description":"### Interesting Title\nSome application content, and description\n"
}' http://localhost:8080/apps
```

##### Send yaml msg :
```
curl --data '---
company: "Random Inc."
description: |-
    ### Interesting Title
    Some application content, and description
license: Apache-2.0
maintainers: 
  - 
    email: firstmaintainer@hotmail.com
    name: "firstmaintainer app1"
  - 
    email: secondmaintainer@gmail.com
    name: "secondmaintainer app1"
source: "https://github.com/random/repo"
title: "Valid App 1"
version: "123"
website: https://website.com' http://localhost:8080/apps
```

##### Get json msg :
```
http://localhost:8080/apps/json/{id}
```

##### Get yaml msg :
```
http://localhost:8080/apps/yaml/{id}
```

##### Search example :
```
http://localhost:8080/apps/search/company=Random
```


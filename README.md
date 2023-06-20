# Sample Stripe Scripts

This is a sample project to demonstrate functional architecture in go.

**Table of Contents**

- [Why Clean Architecture?](#why-clean-architecture)
- [Application Structure](#application-structure)
- [Testing Strategy](#testing-strategy)
- [Building and Running the application](#building-and-running-the-application)
- [API request Doc ](#the-example-domain)

## Why We Need Clean Architecture?

Clean architecture helps us solve, or at least mitigate, these common problems with architecture:

- **Decisions are taken too early**, often at the beginning of a project, when we know the least about the problem that we have to solve
- **It's hard to change**, so when we discover new requirements we have to decide if we want to hack them in or go through an expensive and painful re-design. We all know which one usually wins. _The best architectures are the ones that allow us to defer commitment to a particular solution and let us change our mind_
- **It's centered around frameworks**. Frameworks are tools to be used, not architectures to be conformed to. Frameworks often require commitments from you, but they don’t commit to you. They can evolve in different directions, and then you’ll be stuck following their rules and quirks
- **It's centered around the database**. We often think about the database first, and then create a CRUD system around it. We end up using the database objects everywhere and treat everything in terms of tables, rows and columns
- **We focus on technical aspects** and when asked about our architecture we say things like “it’s servlets running in tomcat with an oracle db using spring”
- **It's hard to find things** which makes every change longer and more painful
- **Business logic is spread everywhere**, scattered across many layers, so when checking how something works our only option is to debug the whole codebase. Even worse, often it's duplicated in multiple places
- **Forces/Encourages slow, heavy tests**. Often our only choice for tests is to go through the GUI, either because the GUI has a lot of logic, or because the architecture doesn't allow us to do otherwise. This makes tests slow to run, heavy and brittle. It results in people not running them and the build beind broken often
- **Infrequent deploys** because it's hard to make changes without breaking existing functionalities. People resort to long-lived feature branches that only get integrated at the end and result in big releases, rather than small incremental ones

## Application Structure

### /app

##### app

- define and intialize modules for main service

### /cmd

##### cmd

- Main applications for this service
- Use cli module for config args

### /internal

##### internal

- Private application and library code.
- Define Core business logic entities
- Handler endpoint module
- Custom http request and response and middleware
- Custom request type
- Custom response type
- Define route endpoint structure

### /model

##### repositories or model

- an abstraction that allows to access entities by using a collection like interface

### /pkg

##### pkg

- an abstraction that allows to access entities by using a collection like interface
- Library code that's ok to use by external applications (e.g., /pkg/mypubliclib). Other projects will import these libraries expecting them to work, so think twice before you put something here :-) Note that the internal directory is a better way to ensure your private packages are not importable because it's enforced by Go. The /pkg directory is still a good way to explicitly communicate that the code in that directory is safe for use by others. The I'll take pkg over internal blog post by Travis Jeffery provides a good overview of the pkg and internal directories and when it might make sense to use them.
- loading config file with viper package
- define constant for loading configuration setting
- define database interface and module for using database provider
- define global errors

### /service

##### service

- an abstraction that allows to access logic systems as single responsibility logic layer

## Testing Strategy

Use Package gomock framework for Testing Purpose

##### Unit Tests

- Unit Test For Service
- Unit Test For Model
- Unit Test For Entities & Handler
- Unit Test For Database pkg

##### End-to-end Tests

- End-to-end test for handler

## Building and Running the application

#### build for windows

```
make build-windows
```

#### build for linux

```
make build-linux
```

#### build for mac

```
make build-mac
```

## Run application

make run

## Run tests

make test

## API request Doc

### Create Customer

```
curl -X POST -H "Content-type: application/json" \
-d '{
	"name":"test name of user",
	"email":"sample@gmail.com"
}' 'http://localhost:8080/customer'

```

### Create Payment Method

```
curl -X POST -H "Content-type: application/json" \
-d '{
	"CustomerID":"customer sample id",
	"Number":"4242424242424242",
	"Cvc":"576",
	"ExpMonth":12,
	"ExpYear":34
}' 'http://localhost:8080/paymentmethods'

```

### Payment Intent With Confirm Refund Flow

```
curl -X POST -H "Content-type: application/json" \
-d '{
	"customerID":"customer sample id",
	"description":"sample description",
    "amount":300
}' 'http://localhost:8080/pi_with_confirm_refund_flow'

```

### Payment Intent With Cancel Flow

```
curl -X POST -H "Content-type: application/json" \
-d '{
	"customerID":"customer sample id",
	"description":"sample description",
    "amount":300
}' 'http://localhost:8080/pi_with_cancel_flow'

```

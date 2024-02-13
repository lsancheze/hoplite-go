# Hoplite - Go

Boilerplate to start working in APIs just right ahead, using Go and Gin Framework.

![API diagram.](/img/api-diagram.png)

This approach takes care of critical points in the lifecycle of an API request, where the Router and Business Layers are the focus.

## Route Middleware

It offers a streamlined way to perform operations across all the endpoints (or some of them, it's your choice!) using the Route Middleware. You can add the Route Middleware on-demand, to each endpoint that you would like.

## Validation Middleware

Validating requests is sometimes a pain, but with the Validation Middleware, you can validate all kinds of data using regular expressions. You just have to create your schemas inside the entities folder and then add them in the Router layer. It's not necessary to use regular expressions, you may also use the [Gin validator](https://github.com/go-playground/validator). If you will use regular expressions, it's good to know that some people may want to move the regex folder to be placed in each entity, I would encourage that as well, it's your call!

## Error Middleware

It's not always easy to handle errors, but with Error Middleware, you can have peace of mind. The Error Middleware will catch any error triggered by the request execution, and using the custom error class ServiceError, it will quickly know about the nature of the error. The ServiceError class is made to include several data points, like code, message, traceId, and flightRecorder. You don't have to use them all, but I highly encourage you to use at least the code and message data points to identify errors.

![Hoplite Framework.](/img/hoplite.png)

## Use

### Prerequisites

* Go: any one of the three latest major [releases](https://go.dev/doc/devel/release). Please visit the official documentation to know more about how to [install](https://go.dev/doc/install) Go.

### Running Hoplite

Use the Go command to run the server:

```
$ go run src/main.go
```

Now, you can start playing with Postman, Insomnia, or any other tool to send requests using `0.0.0.0:8088`

## Let's improve!

Feel free to adapt this code to your necessities, you can also use another Go framework if you want.

Let's keep in touch, don't hesitate to contact me about questions, improvements, or ideas.
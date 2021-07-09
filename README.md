# Dockerized Signal Messenger gRPC API
This project is forked from [signal-cli-rest-api](https://github.com/datti-to/signal-cli-grpc-api) to create a gRPC based version.

This project creates a small dockerized gRPC API around [signal-cli](https://github.com/AsamK/signal-cli).

At the moment, the following functionality is exposed via gRPC:

- Register a number
- Verify the number using the code received via SMS
- Send message (+ attachments) to multiple recipients (or a group)
- Receive messages
- Link devices
- Create/List/Remove groups
- List/Serve/Delete attachments
- Update profile

and many more

## Examples

Sample `docker-compose.yml`file:

```sh
version: "3"
services:
  signal-cli-grpc-api:
    image: dattito/signal-cli-grpc-api:latest
	environment:
	  - USE_NATIVE=0
      #- AUTO_RECEIVE_SCHEDULE=0 22 * * * #enable this parameter on demand (see description below)
	ports:
      - "9090:9090" #map docker port 9090 to host port 9090.
    volumes:
      - "./signal-cli-config:/home/.local/share/signal-cli" #map "signal-cli-config" folder on host system into docker container. the folder contains the password and cryptographic keys when a new number is registered

```

## Auto Receive Schedule

[signal-cli](https://github.com/AsamK/signal-cli), which this gRPC API wrapper is based on, recommends to call `receive` on a regular basis. So, if you are not already calling the `receive` endpoint regularily, it is recommended to set the `AUTO_RECEIVE_SCHEDULE` parameter in the docker-compose.yml file. The `AUTO_RECEIVE_SCHEDULE` accepts cron schedule expressions and automatically calls the `receive` endpoint at the given time. e.g: `0 22 * * *` calls `receive` daily at 10pm. If you are not familiar with cron schedule expressions, you can use this [website](https://crontab.guru).

**WARNING** Calling `receive` will fetch all the messages for the registered Signal number from the Signal Server! So, if you are using the gRPC API for receiving messages, it's _not_ a good idea to use the `AUTO_RECEIVE_SCHEDULE` parameter, as you might lose some messages that way.


## Native Image (EXPERIMENTAL)

On Systems like the Raspberry Pi, some operations like sending messages can take quite a while. That's because signal-cli is a Java application and a significant amount of time is spent in the JVM (Java Virtual Machine) startup. signal-cli recently added the possibility to compile the Java application to a native binary (done via GraalVM).

By adding `USE_NATIVE=1` as environmental variable to the `docker-compose.yml` file the native mode will be enabled.

In case there's no native binary available (e.g on a 32 bit Raspian OS), it will fall back to the signal-cli Java application. The native mode only works on a 64bit system, when the native mode is enabled on a 32bit system, it falls back to the Java application.

## Clients & Libraries

|     Name    | Client           | Library  | Language | Maintainer |
| ------------- |:-------------:| :-----:|:-----:|:-----:|
| [Shell Client](https://github.com/florian-h05/shell-script_collection/blob/main/signal-cli-rest-api_client.bash)      | X | | Shell | [@florian-h05](https://github.com/florian-h05)
| [pysignalclirestapi](https://pypi.org/project/pysignalclirestapi/)      | | X | Python | [@bbernhard](https://github.com/bbernhard)

In case you need more functionality, please **file a ticket** or **create a PR**.

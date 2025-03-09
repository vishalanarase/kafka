# Kafka Golang Streaming Example 📡

> This repository contains a simple example of streaming data using Kafka with a Golang application. It demonstrates how to create a Kafka producer and consumer using the **Confluent Kafka Docker images** and the **`kafka-go` Go client**.

The project includes:
- A **Kafka producer** that sends messages to a Kafka topic 📩.
- A **Kafka consumer** that reads messages from the Kafka topic 📥.
- A **Docker Compose** setup for running Kafka and Zookeeper locally 🐳.

## 📚 **Table of Contents**
- [Prerequisites](#Prerequisites)
- [Installation](#Installation)
- [Usage](#Usage)
  - [Running the Producer](#running-the-producer)
  - [Running the Consumer](#running-the-consumer)
- [Docker Setup](#docker-setup)
- [Contributing](#contributing)
- [Code of Conduct](#code-of-conduct)
- [License](#license)

## Prerequisites ⚙️

To run this project, you'll need the following installed on your system:

- **Go (1.18+)**: [Download Go](https://golang.org/dl/)
- **Docker**: [Download Docker](https://www.docker.com/get-started) (to run Kafka and Zookeeper in containers)
- **Docker Compose**: [Install Docker Compose](https://docs.docker.com/compose/install/) (if not included with Docker)

## Installation 🛠️

### Clone the Repository

Start by cloning this repository to your local machine:

```bash
git clone https://github.com/vishalanarase/kafka.git
cd kafka
```

### Install Go Dependencies

This project uses Go modules for dependency management. To install the necessary dependencies, run:

```bash
go mod tidy
```

## Usage

This project contains two main components:
1. **Producer**: Sends messages to a Kafka topic.
2. **Consumer**: Reads messages from the Kafka topic.

Before running the producer and consumer, you need to start **Kafka** and **Zookeeper** using Docker Compose. This will ensure that the necessary Kafka services are running.

### Docker Setup 🐳

#### Starting Kafka and Zookeeper with Docker Compose

To set up Kafka and Zookeeper locally using Docker, run the following command:

```bash
docker-compose up -d
```

This will start the required services:
- **Zookeeper** on port `2181`
- **Kafka** on port `9093`

Kafka and Zookeeper will be running in detached mode. You can verify their status with:

```bash
docker ps
```

#### Stopping Kafka and Zookeeper

To stop the services and remove the containers, use the following command:

```bash
docker-compose down
```

### Running the Producer 🎯

To run the Kafka **producer**, execute the following command:

```bash
go run producer/producer.go
```

The producer will send a message to the Kafka topic `purchases`. You should see an output message confirming that the message was sent.

### Running the Consumer 👀

To start the **consumer** and begin reading messages from the Kafka topic `purchases`, run:

```bash
go run consumer/consumer.go
```

You should see the consumer output the message sent by the producer.

---

## 🙋‍♂️ Contributing

Contributions are welcome! Please feel free to open issues and pull requests. Here are some ways you can help:

1. **Report Bugs**: If you encounter any issues with the Kafka producer/consumer or the Docker setup, please open an issue.
2. **Fix Bugs**: If you're able to fix bugs or add new features, open a pull request with your changes.
3. **Enhance Documentation**: If you think any part of the documentation could be clearer, feel free to submit improvements.

### How to Contribute:
1. Fork the repository.
2. Create a new branch (`git checkout -b feature-xyz`).
3. Make your changes and commit them (`git commit -am 'Add feature xyz'`).
4. Push to the branch (`git push origin feature-xyz`).
5. Open a pull request.

## 🧑‍🤝‍🧑 **Code of Conduct**

This project and everyone participating in it is governed by the [Contributor Covenant Code of Conduct](https://www.contributor-covenant.org/). By participating, you are expected to uphold this code. Please report unacceptable behavior to the project team.

The goal of this code of conduct is to ensure a welcoming, respectful, and productive environment for all contributors. The expectations are:

- **Be respectful**: Treat everyone with kindness and empathy.
- **Be inclusive**: Embrace diversity of thought, background, and perspective.
- **Be responsible**: Take ownership of your contributions and interactions.
- **Be collaborative**: Foster a cooperative environment where knowledge sharing is encouraged.

For more details, please refer to the full [Code of Conduct](https://www.contributor-covenant.org/version/2/0/code_of_conduct.html).

## License 📜

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ⭐ **Please Star the Repo**

If you find this project useful, please consider giving it a ⭐️ on GitHub. Your star helps us reach more developers and contributors and supports the continued improvement of the project. It’s a great way to show appreciation! 😊

---
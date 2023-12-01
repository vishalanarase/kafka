# Kafka

> Running Kafka locally using Docker is straightforward. Docker makes it easy to set up and run Kafka in a single-node configuration. Here are the steps to run Kafka locally using Docker:

## Step 1: Install Docker

Ensure that Docker is installed on your system. You can download and install Docker Desktop from the [official Docker website](https://www.docker.com/products/docker-desktop).

### Step 2: Create a Docker Compose File for Kafka

Create a `docker-compose.yml` file with the Kafka and ZooKeeper configurations:

```yaml
version: '3'
services:
  zookeeper:
    image: confluentinc/cp-zookeeper:latest
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000

  kafka:
    image: confluentinc/cp-kafka:latest
    depends_on:
      - zookeeper
    ports:
      - "9092:9092"
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka:29092,PLAINTEXT_HOST://localhost:9092
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
```

### Step 3: Start Kafka and ZooKeeper Services

Navigate to the directory containing the `docker-compose.yml` file in your terminal and run the following command:

```bash
docker-compose up -d
```

This command will start Kafka and ZooKeeper containers in the background based on the configurations specified in the `docker-compose.yml` file.

### Step 4: Verify Kafka is Running

Check if Kafka is running by inspecting the Docker containers:

```bash
docker ps
```

You should see containers named `docker_kafka_1` and `docker_zookeeper_1` running.

### Step 5: Access Kafka Services

- To interact with Kafka, you can use tools like `kafkacat` or other Kafka clients.
- Connect to Kafka using `localhost:9092` for your Kafka-related operations.
- You can also create topics, produce messages, and consume messages using command-line tools like `kafka-topics.sh`, `kafka-console-producer.sh`, and `kafka-console-consumer.sh`.

### Step 6: Stop Kafka and ZooKeeper Services

To stop Kafka and ZooKeeper containers, run:

```bash
docker-compose down
```

This command will stop and remove the running Kafka and ZooKeeper containers.

This setup provides a basic local single-node Kafka instance for development and testing purposes. Adjust configurations or expand the setup for production scenarios and more complex requirements.

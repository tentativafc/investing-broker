# Investing Broker

Provides a complete tool to control personal investments.

```sh
cd ./app
docker-compose up
```

### Services and port numbers

| Service                 | Port Number | Type/Tech   |
| ----------------------- | ----------- | ----------- |
| Postgres                | 5432        | DB          |
| Postgres STS            | 5433        | DB          |
| Mongo-IbovespaAssets    | 27017       | DB          |
| Mongo-Currency          | 27018       | DB          |
| Mongo-Selic             | 27019       | DB          |
| Mongo-Sts               | 27020       | DB          |
| Sts-Service             | 50051       | gRpc        |
| User-Service            | 8081        | Rest        |
| Ibovespa-Assets-Scraper | 8082        | Rest        |
| Currency-Scraper        | 8083        | Rest        |
| Selic-Scraper           | 8084        | Rest        |
| Frontend                | 8080        | Vue         |
| Nginx                   | 80          | HTTP Server |

Postman collection of project: investing.postman_collection.json

## Analytics

### Environment

1. Install [Metatrader 5](https://www.metatrader5.com)

1. Edit home project directory on config module:

   ```ini
   [APP]
   home=C:/Users/tenta/Documents/Projetos/investing-broker/analytics
   ```

1. Run commands:

   ```sh
   conda create -n investing-broker python=3.6.12
   conda activate investing-broker
   pip install -r requirements.txt
   jupyter notebook notebooks
   docker-compose up -d
   ```

1. Install ortisan-investing-lib:

   ```sh
   pip install -e ./ortisan-ta --user
   ```

1. Run your Analysis: http://localhost:8888

### Docs

- [Metatrader5 Python lib](https://www.mql5.com/pt/docs/integration/python_metatrader5)

- [TAlib Python](https://technical-analysis-library-in-python.readthedocs.io/en/latest/ta.html#trend-indicators)

### Commands:

Delete all containers:

```sh
docker rm -f $(docker container  ls -aq)
```

Create docker image:

```sh
docker build -t sts-service -f Dockerfile .
```

Run docker image:

```sh
docker run --rm -d --network host --name my_sts_service sts-service:latest
```

Generating GRPC stub:

```sh
protoc sts.proto --go_out=plugins=grpc:.
```

Start cluster K8s:

```sh
minikube start
```

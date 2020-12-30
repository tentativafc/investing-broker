# Investing Broker

Provides a complete tool to control personal investments.

## Frontend

```sh
cd ./app/frontend/app

npm run serve
```

## Backend

```sh
cd ./app/backend/user-rest
docker-compose up
go run main.go
```

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
* [Metatrader5 Python lib](https://www.mql5.com/pt/docs/integration/python_metatrader5)

* [TAlib Python](https://technical-analysis-library-in-python.readthedocs.io/en/latest/ta.html#trend-indicators)




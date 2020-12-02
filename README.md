# Investiment Broker

### Environment

1. Install [Metatrader 5](https://www.metatrader5.com)

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
    pip install -e ./ortisan-investing-lib --user
   ```

1. Run your Analysis: http://localhost:8888

### Docs
* [Metatrader5 Python lib](https://www.mql5.com/pt/docs/integration/python_metatrader5)

* [TAlib Python](https://technical-analysis-library-in-python.readthedocs.io/en/latest/ta.html#trend-indicators)




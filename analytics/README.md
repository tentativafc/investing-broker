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
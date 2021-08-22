# Bulbasaur
This is Pokemon 😄. (WIP)

Using [godog](https://github.com/DATA-DOG/godog) and [Gherkin](https://docs.cucumber.io/gherkin/), make behavioral tests with scripts easier to understand for developers, QA or others in the product team.
### List of Resources
* HTTP Client - to send http request
* Database 
    * PostgreSQL  (*coming soon)*
    * MySQL (*coming soon)*
    
* Redis (*coming soon)*

## Getting Started

### Set up your configuration
Tomato integrates your app and its test dependencies using a simple configuration file `tomato.yml`.

Create a `config.yml` file with your application's required test [resources]():
```yml
---

# Stops on the first failure
stop_on_failure: false

# All feature file paths
features_path:
    - ./features
    - feature-name.feature

# List of resources for application dependencies
resources:
    - name: your-resource-name
      type: postgres
      options:
        datasource: postgres://USERNAME:PASSWORD@localhost:5432/DB_NAME

    - name: your-resource-name
      type: httpclient
      options:
        base_url: http://127.0.0.1:8080
```

### Write your first feature test
Write your own [Gherkin](https://docs.cucumber.io/gherkin/) feature (or customize the check-status.feature example below) and place it inside ./features/check-status.feature:

```gherkin
Feature: Check my application's status endpoint

  Scenario: My application is running and active
    Given "your-resource-name" send request to "GET /healthCheck"
    Then "your-resource-name" response code should be 200
```

### How to Run
```sh
bulbasaur config.yml
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

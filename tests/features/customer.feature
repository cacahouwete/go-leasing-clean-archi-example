Feature: CRUD Customer API
  Scenario: GET Collection
    When I send "GET" request to "/api/v1/customers"
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "totalItems" should be a number
    And the JSON node "totalItems" should be equal to 3
    And the JSON node "member" should be an array of length 3
    And the JSON node "member.0.name" should be equal to "Jean Dupont"
    And the JSON node "member.1.name" should be equal to "Banana"

  Scenario: GET single entity
    When I send "GET" request to "/api/v1/customers/cust1"
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "name" should be equal to "Jean Dupont"

  Scenario: POST new entity
    When I send "POST" request to "/api/v1/customers" with body:
    """
    {
      "name": "Toto"
    }
    """
    Then the response status code should be 201
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "name" should be equal to "Toto"

    When I send "GET" request to "/api/v1/customers"
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "totalItems" should be a number
    And the JSON node "totalItems" should be equal to 4

  @db:clean
  Scenario: PUT an entity
    When I send "PUT" request to "/api/v1/customers/cust1" with body:
    """
    {
      "name": "Toto"
    }
    """
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "name" should be equal to "Toto"

    When I send "GET" request to "/api/v1/customers"
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "totalItems" should be a number
    And the JSON node "totalItems" should be equal to 3

    When I send "GET" request to "/api/v1/customers/cust1"
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "name" should be equal to "Toto"

  Scenario: DELETE entity
    When I send "DELETE" request to "/api/v1/customers/cust2"
    Then the response status code should be 500

    When I send "DELETE" request to "/api/v1/customers/cust3"
    Then the response status code should be 204

    When I send "DELETE" request to "/api/v1/customers/cust3"
    Then the response status code should be 404

Feature: CRUD Car API
  Scenario: GET Collection
    When I send "GET" request to "/api/v1/cars"
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "totalItems" should be a number
    And the JSON node "totalItems" should be equal to 2
    And the JSON node "member" should be an array of length 2
    And the JSON node "member.0.name" should be equal to "My Super Car"
    And the JSON node "member.1.name" should be equal to "My hippest Car"

  Scenario: GET single entity
    When I send "GET" request to "/api/v1/cars/car1"
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "name" should be equal to "My Super Car"

  Scenario: POST new entity
    When I send "POST" request to "/api/v1/cars" with body:
    """
    {
      "name": "Toto"
    }
    """
    Then the response status code should be 201
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "name" should be equal to "Toto"

    When I send "GET" request to "/api/v1/cars"
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "totalItems" should be a number
    And the JSON node "totalItems" should be equal to 3

  @db:clean
  Scenario: PUT an entity
    When I send "PUT" request to "/api/v1/cars/car1" with body:
    """
    {
      "name": "Toto"
    }
    """
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "name" should be equal to "Toto"

    When I send "GET" request to "/api/v1/cars"
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "totalItems" should be a number
    And the JSON node "totalItems" should be equal to 2

    When I send "GET" request to "/api/v1/cars/car1"
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "name" should be equal to "Toto"

  Scenario: DELETE entity
    When I send "DELETE" request to "/api/v1/cars/car2"
    Then the response status code should be 204

    When I send "DELETE" request to "/api/v1/cars/car2"
    Then the response status code should be 404

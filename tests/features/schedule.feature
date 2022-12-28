Feature: CRUD Schedule API
  Scenario: GET Collection
    When I send "GET" request to "/api/v1/schedules"
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "totalItems" should be a number
    And the JSON node "totalItems" should be equal to 2
    And the JSON node "member" should be an array of length 2
    And the JSON node "member.0.id" should be equal to "schedule1"

  Scenario: GET single entity
    When I send "GET" request to "/api/v1/schedules/schedule1"
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "customerId" should be equal to "cust1"
    And the JSON node "carId" should be equal to "car1"

  Scenario: POST new schedule without overlap with an other on the same car
    When I send "POST" request to "/api/v1/schedules" with body:
    """
    {
      "beginAt": "2021-01-02T08:00:00Z",
      "endAt": "2021-01-03T18:00:00Z",
      "customerId": "cust3",
      "carId": "car1"
    }
    """
    Then the response status code should be 201
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "customerId" should be equal to "cust3"
    And the JSON node "carId" should be equal to "car1"

    When I send "POST" request to "/api/v1/schedules" with body:
    """
    {
      "beginAt": "2021-01-02T08:00:00Z",
      "endAt": "2021-01-03T18:00:00Z",
      "customerId": "cust2",
      "carId": "car1"
    }
    """
    Then the response status code should be 422
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "violations.0.code" should be equal to "timewindow_overlap"

    When I send "GET" request to "/api/v1/schedules"
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "totalItems" should be a number
    And the JSON node "totalItems" should be equal to 3

  @db:clean
  Scenario: PUT a schedule without overlap with an other on the same car
    When I send "PUT" request to "/api/v1/schedules/schedule1" with body:
    """
    {
      "beginAt": "2022-01-06T08:00:00Z",
      "endAt": "2022-01-08T18:00:00Z"
    }
    """
    Then the response status code should be 422
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "violations.0.code" should be equal to "timewindow_overlap"

    When I send "PUT" request to "/api/v1/schedules/schedule1" with body:
    """
    {
      "beginAt": "2021-01-06T08:00:00Z",
      "endAt": "2021-01-08T18:00:00Z"
    }
    """
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "beginAt" should be equal to "2021-01-06T08:00:00Z"
    And the JSON node "endAt" should be equal to "2021-01-08T18:00:00Z"

    When I send "GET" request to "/api/v1/schedules"
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "totalItems" should be a number
    And the JSON node "totalItems" should be equal to 2

    When I send "GET" request to "/api/v1/schedules/schedule1"
    Then the response status code should be 200
    And the response payload should be JSON
    And the JSON should be valid
    And the JSON node "beginAt" should be equal to "2021-01-06T08:00:00Z"
    And the JSON node "endAt" should be equal to "2021-01-08T18:00:00Z"

  Scenario: DELETE entity
    When I send "DELETE" request to "/api/v1/schedules/schedule1"
    Then the response status code should be 204

    When I send "DELETE" request to "/api/v1/schedules/schedule1"
    Then the response status code should be 404

Feature: Ready and health check
  In order to know when app is alive
  As anonyme user
  I need have api routes for ready and health check

  @reset
  Scenario: Health check
    When I send "GET" request to "/healthz"
    Then the response status code should be 200

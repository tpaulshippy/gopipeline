Feature: SQL

  Scenario: Connect to bad port
    Given I provide server of "(local)" with port '1'
    When I connect
    Then I should not be connected
    And I should see error "Unable to open tcp connection with host 'localhost:1': dial tcp 127.0.0.1:1: connectex: No connection could be made because the target machine actively refused it."

  Scenario: Connect to bad server
    Given I provide server of "non-existent-server.net" with port '1433'
    When I connect
    Then I should not be connected
    And I should see error "lookup non-existent-server.net: no such host"


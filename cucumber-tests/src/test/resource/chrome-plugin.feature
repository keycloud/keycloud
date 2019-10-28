@WithPlugin
Feature: KeyCloud - Chrome Plugin

  Scenario Outline: Get the password for a opened login page
    Given I am on "<site>"
    When I click on the autofill - icon
    And I successfully log in
    Then the login for "<site>" should be filled
    Examples:
      | site|
      | https://www.google.com/ |
@WithoutPlugin
Feature: Dashboard

  Scenario Outline: Register a new Account in KeyCloud
    Given I am on the landing page
    When I type in <user> as my username and click register
    Then I will be on the settings page of a new created Account
    Examples:
      | user |
      | "ich" |

  Scenario: Add a new password to my list
    Given I am on my home page in the keycloud dashboard
    When I press the add button
    And I fill out the popup
    Then I will see a new password added to the list

  Scenario Outline: Remove a password from my list
    Given I am on my home page in the keycloud dashboard
    When  I press the remove button for the <password> password
    Then The password <password> entry is removed from the list
    Examples:
      | password |
      | "https://google.com"   |

  Scenario Outline: Get password
    Given I am on my home page in the keycloud dashboard
    When I copy the password for <url> to clipboard
    Then I have the password for <url> in my clipboard
    Examples:
      | url |
      | "https://google.com"   |

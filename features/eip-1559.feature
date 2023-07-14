Feature: EIP-1559
  In order to save money
  As a wallet holder
  I want to use dynamic fee transactions

  Scenario: Create a dynamic fee transaction
    Given I send a dynamic fee transaction to 0x3816756E807ADC276c2fa820D78AEF02C2C11A38
    When I call eth_getTransactionByHash
    Then I get all the details for my transaction
    And the transaction receipt should have a status of 1

#   Scenario: Create a dynamic fee transaction
#     Given I bridge 7500000000000000000 IMX to 0xA887D80C6746f8d6B586767871c5cD57E080E2E3 on L2
#     When 0xA887D80C6746f8d6B586767871c5cD57E080E2E3 sends 10000000000000000 IMX to 0x3816756E807ADC276c2fa820D78AEF02C2C11A38 on L2 using EIP-1559
#     Then the transaction on chain should match what was submitted

#     Given I bridge 7500000000000000000 IMX to 0xA887D80C6746f8d6B586767871c5cD57E080E2E3 on L2
#     When 0xA887D80C6746f8d6B586767871c5cD57E080E2E3 sends 10000000000000000 IMX to 0x3816756E807ADC276c2fa820D78AEF02C2C11A38 on L2 using EIP-1559
#     And we wait for the transaction to be mined
#     Then the receipt should have a status of 1


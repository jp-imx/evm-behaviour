Feature: EIP-1559
  In order to save money
  As a wallet holder
  I want to use dynamic fee transactions

  Scenario: Bridge and transfer with EIP-1559
    Given receiver 0x3816756E807ADC276c2fa820D78AEF02C2C11A38 has a certain balance on L2
    And sender bridges 7500000000000000000 IMX from L1 to receiver on L2
    And sender transfers 10000000000000000 IMX to receiver on L2
    And sender withdraws 3000000000000000000 IMX to L1 
    Then the receiver balance on L2 should increase by 10000000000000000
    Then sender IMX balance on L1 should increase by 3000000000000000000


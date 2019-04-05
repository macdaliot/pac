Feature: ACME API Features

  @UseCaseThreeAPI
  Scenario Outline: POST Request: enter acme manager
    Then the <manager> information will be returned

    Examples: 
      | manager      |
      | managerTest1 |

  @UseCaseThreeAPITODO
  Scenario Outline: POST Request: enter partner organization information
    Then the <partner> information will be returned

    Examples: 
      | partner      |
      | partnerTest1 |

  #only thing that will be updated are the approvers
  #list of approver ids
  #need get endpoint
  #Approver Name, email, isApproved (boolean)
  @UseCaseThreeAPITODO
  Scenario Outline: Update Request: update partner organization information
    Then the <approver1> will be updated to <approvers2>

    Examples: 
      | approvers1   | approvers2   |
      | managerTest1 | managerTest2 |

  @UseCaseThreeAPITODO
  Scenario Outline: Get Request: get partner organization information
    Then the <partner> information will be returned

    Examples: 
      | partner      |
      | partnerTest1 |

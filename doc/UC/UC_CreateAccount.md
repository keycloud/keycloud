## Revision History
Date | Version | Description | Author
--- | --- | --- | ---
21.10.19 | 1.0 | First release of the UC | Lukas Priester

## Table of Contents
- [1. Use Case Create Account](#1-use-case-create-account)
  - [1.1 Brief Description](#11-brief-description)
- [2. Flow of Events](#2-flow-of-events)
  - [2.1 Basic Flow](#21-basic-flow)
  - [2.2 Alternative Flows](#22-alternative-flows)
- [3. Sepcial Requirements](#3-special-requirements)
- [4. Preconditions](#4-preconditions)
- [5. Postconditions](#5-postconditions)
- [6. Extension Points](#6-extension-points)

## 1. Use-Case Create Account
### 1.1 Brief Description
This use case allows the user to create an account. For the creation following credentials must be provided:
- username

## 2. Flow of Events
### 2.1 Basic Flow
- User enters username
- User clicks on register
- User is prompted with a password
#### 2.1.1 Activity Diagram
![UC_CreateAccount](images/UC/UC_CreateAccount.png) 
 
.feature  
![FeatureFile](images/featureFileScreenshots/Featurefile_UC_RegisterAccount.PNG)
#### 2.1.2 Mock up
##### Register
![Mockup_Register](images/mockups/Mockup_register.PNG)
##### Create Password
![Mockup_CreatePassword](images/mockups/Mockup_createPassword.PNG)
### 2.2 Alternative Flows
n/a

## 3. Special Requirements
n/a

## 4. Preconditions
n/a

## 5. Postconditions
If a user completes this workflow, the users profile must have been created on the server. 

## 6. Extension Points
n/a
# KeyCloud - Software Requirements Specifiction  

## Table of Contents
- [1. Introduction](#1-introduction)
    - [1.1 Purpose](#11-purpose)
    - [1.2 Scope](#12-scope)
    - [1.3 Definition, Acronyms and Abbreviations](#13-definitions-acronyms-and-abbreviations)
    - [1.4 References](#14-references)
    - [1.5 Overview](#15-overview)
- [2. Overall Description](#2-overall-description)
    - [2.1 Vision](#21-vision)
    - [2.2 Use Case Diagram](#22-use-case-diagram)
- [3. Specific Requirements](#3-specific-requirements)
    - [3.1 Functionality](#31-functionality)
    - [3.2 Usability](#32-usability)
    - [3.3 Reliability](#33-reliability)
    - [3.4 Performance](#34-performance)
    - [3.5 Supportability](#35-supportability)
    - [3.6 Design Constraints](#36-design-constraints)
    - [3.7 Online User Documentation and Help System Requirements](#37-on-line-user-documentation-and-help-system-requirements)
    - [3.8 Purchased Components](#38-purchased-components)
    - [3.9 Interfaces](#39-interfaces)
    - [3.10 Licensing Requirements](#310-licensing-requirements)
    - [3.11 Legal, Copyright and Other Notices](#311-legal-copyright-and-other-notices)
    - [3.12 Applicable Standards](#312-applicable-standards)
- [4. Supporting Information](#4-supporting-information)

## 1. Introduction
### 1.1 Purpose  
This Software Requirements Specification (SRS) was created to collect and organize the requirements for the KeyCloud Application and all its components. It includes a thorough description of the expected functionality for the project, as well as the nonfunctional requirements. These requirements are crucial as they minimize the risks of not meeting customer's expectations and establish a clear understanding of what the application is capable of doing and what not. This document will furthermore provide the basis for costs-estimation and later validation of the results achieved.
### 1.2 Scope
This SRS applies to the entire KeyCloud project. KeyCloud is an open source password manager.  
  
  **ACTORS:**  
  - **user:** Person who creates, stores, retrieves, deletes passwords.  
  
  **SUBSYSTEMS:**  
  - **Dashboard:** Allows the actors to create an account, add passwords, retrieve passwords, get lists of his passwords
  - **Chrome-Plugin:** Allows the actors to retrieve passwords
  - **Android-AutoFiller-Service:** Allows the actors to retrieve passwords
  
### 1.3 Definitions, Acronyms and Abbreviations
Abbreviation | |
--- | --- 
SRS | Software Requirements Specification
UC | Use Case
n/a | not applicable
tbd | to be determined
FAQ | Frequently Asked Questions

Definition | |  
--- | ---  
Software Requirements Specification | a document, which captures the complete software requirements for the system, or a portion of the system
### 1.4 References
Title | Date | Publishing organization |  
--- | :---:  | ---
[KeyCloud Blog](https://keycloud.zeekay.dev/) | 12.10.2019 | KeyCloudTeam  
[YouTrack Instance](https://keycloud-dev.zeekay.dev:7000/issues) | 12.10.2019 | KeyCloudTeam  
[SRS](../doc/SRS.md) | 12.10.2019 | KeyCloudTeam  

### 1.5 Overview
The remainder is structured in the following way: In the next chapter, the overall description, an overview of the functionality and an use-case-diagram is given.
The third chapter, the requirements specification, provides a more detailed description of the requirements.
Further requirements like usability and supportability are listed in chapters 3.2 through 3.12. 
Supporting information is given in the fourth chapter.

## 2. Overall Description
### 2.1 Vision
When it comes to modern authentication 2FA and complex passwords are indispensable. There are many solutions to store your passwords encrypted but all of them lack one core feature – encrypted and secure access with your phone.

KeyCloud offers free cloud space for your account data. You can access you data with our Chrome plugin for your desktop and mobile device.
Don’t worry your data is bio metric encrypted locally and thus only transferred and stored encrypted.
### 2.2 Use Case Diagram


## 3. Specific Requirements

### 3.1 Functionality
#### 3.1.1 Registration
#### 3.1.2 etc

### 3.2 Usability
#### 3.2.1 Unterpunkt

### 3.3 Reliability
#### 3.3.1 Uptime etc

### 3.4 Performance
#### 3.4.1 Response time

### 3.5 Supportability
#### 3.5.1 Coding standards
#### 3.5.2 Maintenance Utilities

### 3.6 Design Constraints

### 3.7 On-line User Documentation and Help System Requirements

### 3.8 Purchased Components
A server hosted at [zap-hosting.com](https://zap-hosting.com/de/) is used to run the frontend application as well as the backend server.  
Currently there are no other purchased components.
### 3.9 Interfaces
#### 3.9.1 User Interfaces
#### 3.9.2 Hardware Interfaces
#### 3.9.3 Software Interfaces
#### 3.9.4 Communications Interfaces

### 3.10 Licensing Requirements
Under [MIT license](https://github.com/zkdev/keycloud/blob/master/LICENSE.md).
### 3.11 Legal, Copyright, and Other Notices

### 3.12 Applicable Standards

## 4. Supporting Information
**For more information contact:**  
- [Florian Drescher](https://github.com/Dudeldu)
- [Philipp Heil](https://github.com/zkdev)
- [Lukas Priester](https://github.com/hottek)
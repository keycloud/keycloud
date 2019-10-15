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
API | Application Programming Interface
CI | Continuous integration
FAQ | Frequently Asked Questions
https | Hypertext Transfer Protocol Secure
n/a | not applicable
SRS | Software Requirements Specification
tbd | to be determined
UC | Use Case

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
![Use Case Diagram](./images/UseCases.png)

## 3. Specific Requirements

### 3.1 Functionality
This section will list all functional requirements for KeyCloud and explain their functionality.
Each of the following subsections represents a subsystem of our application.
#### 3.1.1 Dashboard
By using the dashboard several activities can be managed.  
First of all there will be the option to create a user account. Furthermore the passwords of a user are managed here. Settings can also be changed.  
Additional features might include the generation of passwords.
#### 3.1.2 Chrome-Plugin
The chrome-plugin will enable the user to get his passwords in the Google Chrome browser.
Additional features might include the function to add and remove individual passwords.
#### 3.1.3 Android-AutoFiller-Service
The android-auto-filler-service will enable the user to get his passwords on his android device. There should be a possibility to see licence and security info.
Additional features might include the function to add and remove individual passwords as well as the change of settings.

### 3.2 Usability
We plan on designing any interface as simple and intuitive as possible while still providing maximum functionality. Of course there will be documents available explaining the usage, but we aim at making them redundant.
#### 3.2.1 No training time needed
Our goal is that any user can use our application without any training.
#### 3.2.2 Natural and easy workflow
We want to build an easy to use application that isn't overloaded with features.

### 3.3 Reliability
#### 3.3.1 Availability
While we cannot 100% guarantee it, we want to ensure that the server is available at least 95% of the time, equivalent to approx. 1 hour of downtime per day.
The time to repair bugs etc. should be as low as possible.
#### 3.3.2 Defect Rate and Security
Data should be stored in the most secure way possible. The data transfer is to be made via a secure connection.

### 3.4 Performance
#### 3.4.1 Response time
Should be as low as possible. Maximum response time is 5 seconds.
#### 3.4.2 Capacity
The system should be able to manage multiple requests per second without noticeable latency. The system is to be build on a scalable basis.
#### 3.4.3 Connection bandwidth
The size of data received through one request should be as low as possible, e.g. sending duplicates of data should be avoided.

### 3.5 Supportability
#### 3.5.1 Coding standards
In order to maintain supportability and readability of our code, we will adopt the clean code standard, common naming conventions uniform formatting and best practices throughout the project.
#### 3.5.2 Maintenance Utilities
In order to test and integrate newer versions of the application, a continuous integration service is required.

### 3.6 Design Constraints
Our goal is to provide a modern design in both code and application.  
Server-side programs ought to be compiled using Golang. A REST-API will be used to communicate between client and server.  
The dashboard will be implemented using a MVC architecture.  
The chrome-plugin and the android autofill service will be implemented using their respective technologies and languages.  
#### 3.6.1 Development tools
- Version control system: Git (Github)
- Backend development: JetBrains GoLand
- Frontend development: tbd
- Project planning tool: JetBrains YouTrack
- Build management: tbd
- CI: tbd

### 3.7 On-line User Documentation and Help System Requirements
As stated our goal is to make our application as intuitive as possible. However, we will provide a FAQ and documentation. This will be especially helpful for users who want to know how our application works in depth.

### 3.8 Purchased Components
A server hosted at [zap-hosting.com](https://zap-hosting.com/de/) is used to run the frontend application as well as the backend server.  
Currently there are no other purchased components.

### 3.9 Interfaces
#### 3.9.1 User Interfaces
There will be following interfaces available:
- Dashboard: manage (add, remove), retrieve passwords, change settings
- Chrome-Plugin: manage, retrieve passwords
#### 3.9.2 Hardware Interfaces
n/a
#### 3.9.3 Software Interfaces
As the dashboard is a web-application it should be running on the most common browsers. Nevertheless, we want to primarily support the Google Chrome browser and its features.
#### 3.9.4 Communications Interfaces
The client will connect to the server `https`. An unencrypted connection wil not be supported. 

### 3.10 Licensing Requirements
Under [MIT license](https://github.com/zkdev/keycloud/blob/master/LICENSE.md).

### 3.11 Legal, Copyright, and Other Notices
The KeyCloud team will not take any responsibility for the loss of personal data. The KeyCloud logo may only be used for the official KeyCLoud applications.

### 3.12 Applicable Standards
The following Clean Code standards are going to be applied to the code as far as possible:
- Intuitive names of variables and methods
- Comply with coding conventions of the respective languages
- Comments used to navigate the code but not polluting it
- Design patterns integration
- No premature optimization
- Each method does one thing and does it well

## 4. Supporting Information
**For more information contact:**  
- [Florian Drescher](https://github.com/Dudeldu)
- [Philipp Heil](https://github.com/zkdev)
- [Lukas Priester](https://github.com/hottek)
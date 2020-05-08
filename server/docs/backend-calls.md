# Backend calls
| Method | Route | Description | Header | Body | Requires Cookie | Return | Additional Information
|---|---|---|---|---|---|---|---|
| GET | `/user` | retrieves username and password | - | - | ✔️ | `{"Name": "johndoe", "MasterPassword": "my-master-passwd"}` | - |
| DELETE | `/user` | deletes user | - | - | ✔️ | `{"Status": "REMOVED", "Error": ""}`| [remove passwords before user](https://github.com/keycloud/keycloud/issues/25) |
| PUT | `/user` |  updates username | - | - | ✔️ | - | [user updating not working](https://github.com/keycloud/keycloud/issues/26) |
| GET | `/password` | retrieves specific password | - | `{"username": "johndoe", "url": "john.doe"}` | ✔️ | - | [GET password not working](https://github.com/keycloud/keycloud/issues/28) |
| GET | `/password-url` | retrieves all passwords and usernames according to provided url | - | `{"url": "john.doe"}` | ✔️ | `[{"Password":"doejohn","Id":"9","Url":"john.doe","Username":"johndoe"},{...}}]` | - |
| POST | `/password` | creates new password entry | - | `{"username": "johndoe", "password": "doejohn", "url": "john.doe"}` | ✔️ | `{"Status": "CREATED", "Error": ""}` | - |
| DELETE | `/password` | deletes specific password | - | `{"username": "johndoe", "url": "john.doe"}` | ✔️ | `{"Status": "REMOVED", "Error": ""}` | - |
| GET | `/passwords` | retrieves list of passwords | - | - | ✔️ | `[{"Password": "doejohn", "Id": "3", "Url": "john.doe"}, ...]` | - |
| POST | `/logout` | clears session cookie | - | - | ✔️ | - | - |
| POST | `/webauthn/login/start` | - | - | - | ❌ | - | - |
| POST | `/webauthn/login/finish` | - | - | - | ❌ | - | - |
| POST | `/standard/login` | authenticates user, sets session | - | `{"username": "johndoe", "password": "my-master-passwd"}` | ❌ | cookie: `keycloud-main` | - |
| POST | `/standard/register` | creates new user | `'Accept': 'application/json'` <br> `'Content-Type': 'application/json'` | `{"username": "johndoe", "mail": "john@doe.com"}` | ❌ | generated masterpassword | - |
| POST | `/webauthn/registration/start` | - | - | - | ✔️ | - | - |
| POST | `/webauthn/registration/finish` | - | - | - | ✔️ | - | - |

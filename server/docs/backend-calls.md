# Backend calls
| Method | Route | Description | Parameters | Body | Requires Cookie | Return | Additional Information
|---|---|---|---|---|---|---|---|
| GET | `/user` | retrieves username and password | - | - | ✔️ | `{"username": "johndoe", "masterpassword": "my-master-passwd", "mail": "john@doe.com"}` | - |
| DELETE | `/user` | deletes user | - | - | - | ✔️ | `{"Status": "REMOVED", "Error": ""}`| - |
| PUT | `/user` |  updates username | - | `{"username": "newjohndoe"}` | ✔️ | - | - |
| GET | `/password` | retrieves specific password | - | `{"username": "johndoe", "url": "john.doe"}` | ✔️ | - | - |
| GET | `/password-by-url` | retrieves all passwords and usernames according to provided url | `url=john.doe` | - | ✔️ | `[{password": "doejohn", "id": "3", "url": "john.doe", "username": "johndoe"}, ...]` | - |
| POST | `/password` | creates new password entry | - | `{"username": "johndoe", "password": "doejohn", "url": "john.doe"}` | ✔️ | `{"Status": "CREATED", "Error": ""}` | - |
| DELETE | `/password` | deletes specific password | - | `{"username": "johndoe", "url": "john.doe"}` | ✔️ | `{"Status": "REMOVED", "Error": ""}` | - |
| GET | `/passwords` | retrieves list of passwords | - | - | ✔️ | `[{password": "doejohn", "id": "3", "url": "john.doe", "username": "johndoe"}, ...]` | - |
| POST | `/logout` | clears session cookie | - | - | ✔️ | - | - |
| POST | `/webauthn/login/start` | - | - | - | ❌ | - | - |
| POST | `/webauthn/login/finish` | - | - | - | ❌ | - | - |
| POST | `/standard/login` | authenticates user, sets session | - | `{"username": "johndoe", "masterpassword": "my-master-passwd"}` | ❌ | cookie: `keycloud-main` | - |
| POST | `/standard/register` | creates new user | - | `{"username": "johndoe", "mail": "john@doe.com"}` | ❌ | generated masterpassword | - |
| POST | `/webauthn/registration/start` | - | - | - | ✔️ | - | - |
| POST | `/webauthn/registration/finish` | - | - | - | ✔️ | - | - |


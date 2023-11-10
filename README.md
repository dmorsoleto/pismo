# PISMO TEST

Project developed for a test to PISMO company.

### RUN PROJECT

To run project will be necessary the following libs:

* [docker](https://www.docker.com)
* [golang 1.19](https://go.dev)

Create a .env file on root folder application. 

Only for this test project I added a file called .example.env in ./internal/tests. You can copy all content and paste on .env file that you have create on root folder.

- Run command ```make run``` to run project on terminal

- Local api service up run command ```make local-api``` will start de docker services and will be running application

### API

Call routes with a basic auth, follow credentials:
username: pismo
pass: 12345678

##### Routes

###### Accounts
(GET) - http://localhost:3000/accounts/:accountId
(POST) - http://localhost:3000/accounts

###### Transactions
(POST) - http://localhost:3000/transactions

We have same examples in ./internal/tests/http. This tests can be used directly on vscode getting the addon called "Rest Client". Follow link to download it:

https://marketplace.visualstudio.com/items?itemName=humao.rest-client


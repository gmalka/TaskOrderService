<h1 align="center">
  📚 TaskOrderService
  </h1>
  
  ## 💡 About the project:
  	Simple project designed to emulate the interaction of microservices.
    Implemented two services:
      - UserService allows the user to register or login. Also it allows view and buy tasks.
        provides HTTP api.
      - TaskService provides gRPC and HTTP api for interaction

  ## 🛠 Testing and Running:
  1. To clone project use:
  ```shell
  git clone https://github.com/gmalka/TaskOrderService.git
  ```
  2. Modify .env files if necessary.
  3. Use makefile commands to start:

     - To start app use:
      ```shell
      make
      ```
     - To run tests:
      ```shell
      make test
      ```
     - To rebuild containers:
      ```shell
      make rebuild
      ```
     - To delete containers:
      ```shell
      make down
      ```
     - To down containers and delete images:
      ```shell
      make clean
      ```

  ## 🚀 Usage:
  To try usage project you can go by default:
  
  For UserService:
  
    - localhost:8080/swagger 
  For TaskService:
  
    - localhost:8082/swagger 


  <br/>


  Then you need to register with your login and password:
  
  <img width="1381" alt="image" src="https://github.com/gmalka/TaskOrderService/assets/94842625/5978ffab-e27e-4945-905b-7b1de99d0760"/>
  
  > ❗ **Note:** If you want to user admin handlers you need to login with root/root

  Login with same login and password:
  
  <img width="1389" alt="image" src="https://github.com/gmalka/TaskOrderService/assets/94842625/78b31391-34e1-4a4f-a9d2-ddc52f2b1859"/>

  Your will get jwt token for get access:
  
  <img width="1368" alt="image" src="https://github.com/gmalka/TaskOrderService/assets/94842625/2bd66c43-0fc5-4851-b8c3-83a9e5dbb5fe"/>


  Put it into authorization field like this:
  
  <img width="654" alt="image" src="https://github.com/gmalka/TaskOrderService/assets/94842625/57785020-77f1-4c3b-bff4-72358fa74572"/>
  <br/>
  <br/>


> ❗ **Примечание:** Не было понятно, с помощью чего должна быть реализованна связь между сервисами,
> так что я решил в проекте использовать gGRPC.


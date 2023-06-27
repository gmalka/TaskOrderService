<h1 align="center">
  📚 Mp3 Player
  </h1>
  
  ## 💡 About the project:
  	Simple project designed to emulate the interaction of microservices.
    Implemented two services:
      - UserService allows the user to register or login. Llso it allows view and buy tasks.
        provides HTTP api.
      - TaskService provides gRPC and HTTP api for interaction

  ## 🛠 Testing and Running:
  1. To clone project use:
  ```shell
  git@github.com:gmalka/TaskOrderService.git
  ```
  2. Modify .env files if necessary.
  3. Use makefile commands to start:

     - To start app use:
      ```shell
      make
      ```
     - To test app use:
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
    


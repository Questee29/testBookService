<div id="top"></div>


<!-- PROJECT LOGO -->
<br />

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">EndPoints</a></li>

  </ol>
</details>

<!-- GETTING STARTED -->
## Getting Started

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/Questee29/testBookService
   ```
2. Set your environment (app.env file)
   ```sh
   	DB_DRIVER=mysql
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_USER=root
    DB_PASSWORD=example
    DB_NAME=books_db
 
    SERVER_HOST=localhost
    SERVER_PORT=:8080
  	GRPCSERVER_PORT=:9081
   ```
3. run docker-compose
   ```js
   docker-compose up -d
   ```
4. send grpcCurl request(use postman for example)


### EndPoints

  1. get book using  first_name and last_name
    
    localhost:urRestPort/get-book
    
  2. get author using book_title
    
    localhost:urRestPort/get-author
   
 
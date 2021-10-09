![Pinjem](https://socialify.git.ci/sultanfariz/Pinjem/image?description=1&descriptionEditable=A%20self-initiated%20project%20to%20help%20people%20borrow%20books%20from%20library.%20Developed%20using%20Go%20Echo%20and%20implementing%20Clean%20Architecture.&font=Inter&language=1&logo=https%3A%2F%2Favatars.githubusercontent.com%2Fu%2F55394860%3Fv%3D4&owner=1&pattern=Circuit%20Board&theme=Dark)

<div align="center">
    <img alt="GitHub top language" src="https://img.shields.io/github/languages/top/sultanfariz/Pinjem?style=for-the-badge">
    <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/sultanfariz/Pinjem?style=for-the-badge">
    <img alt="GitHub repo size" src="https://img.shields.io/github/repo-size/sultanfariz/Pinjem?style=for-the-badge">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/sultanfariz/Pinjem?style=for-the-badge">
</div>

### Table of Contents

| [Tech Stack](#tech-stack) | [Entity Relationship Diagram](#entity-relationship-diagram) | [Architectural Pattern](#architectural-pattern) | [Unit Testing](#unit-testing) | [API Postman Docs](#api-postman-docs) | [Swagger OpenAPI Docs](#swagger-openapi-docs) | [Server URL](#server-url) |
| :-----------------------: | :---------------------------------------------------------: | :---------------------------------------------: | :---------------------------: | ------------------------------------- | --------------------------------------------- | ------------------------- |

## Tech Stack

[`^ kembali ke atas ^`](#table-of-contents)

- **Language:** [Go](https://golang.org/)
- **Framework:** [Echo](https://echo.labstack.com/)
- **Database:** [MySQL](https://www.mysql.com/), [RDS](https://aws.amazon.com/rds/)
- **ORM:** [GORM](https://gorm.io/)
- **Unit Testing:** [Mockery](https://github.com/vektra/mockery), [Testify](https://github.com/stretchr/testify)
- **API Testing:** [Postman](https://www.getpostman.com/)
- **API Docs:** [Swagger](https://swagger.io/), [Postman](https://www.getpostman.com/)
- **Container:** [Docker](https://www.docker.com/)
- **Deployment:** [EC2](https://aws.amazon.com/ec2/)
- **CI/CD:** [GitHub Actions](https://github.com/features/actions)
- **Code Editor:** [Visual Studio Code](https://code.visualstudio.com/)

## Entity Relationship Diagram

[`^ kembali ke atas ^`](#table-of-contents)

![ERD](./assets/Pinjemin!-FixERD.png)

## Architectural Pattern

[`^ kembali ke atas ^`](#table-of-contents)

Architectural pattern yang digunakan adalah Clean Architecture, dimana aplikasi terbagi atas 4 layer, antara lain: **Domain/Entity**, **Use Case**, **Controller**, dan **Repository**. Dengan pola seperti ini, semua komponen aplikasi dapat dibuat secara independen sehingga mengurangi dependensi antar komponen dan dapat dikembangkan secara berkelanjutan.

![Clean Architecture](./assets/CleanArch.png)

## Unit Testing

[`^ kembali ke atas ^`](#)

Hasil dari unit test yang telah dilakukan menunjukkan rerata nilai coverage yang berada di atas 80% pada tiap domain bisnis yang ada (6 domain bisnis).
[Link](http://ec2-18-223-110-242.us-east-2.compute.amazonaws.com:8080/api/coverage)

```

ok Pinjem/businesses/book_orders (cached) coverage: 81.2% of statements
? Pinjem/businesses/book_orders/mocks [no test files]
ok Pinjem/businesses/books (cached) coverage: 100.0% of statements
? Pinjem/businesses/books/mocks [no test files]
ok Pinjem/businesses/deposits (cached) coverage: 81.2% of statements
? Pinjem/businesses/deposits/mocks [no test files]
ok Pinjem/businesses/orders (cached) coverage: 84.2% of statements
? Pinjem/businesses/orders/mocks [no test files]
ok Pinjem/businesses/shipping_details (cached) coverage: 100.0% of statements
? Pinjem/businesses/shipping_details/mocks [no test files]
ok Pinjem/businesses/users (cached) coverage: 94.1% of statements
? Pinjem/businesses/users/mocks [no test files]

```

## API Postman Docs

[`^ kembali ke atas ^`](#table-of-contents)

https://documenter.getpostman.com/view/14458184/UUy4e66j

## Swagger OpenAPI Docs

[`^ kembali ke atas ^`](#table-of-contents)

https://app.swaggerhub.com/apis/sultanfariz/Pinjem/1.0.0

## Server URL

[`^ kembali ke atas ^`](#table-of-contents)

http://ec2-18-223-110-242.us-east-2.compute.amazonaws.com:8080/api/v1

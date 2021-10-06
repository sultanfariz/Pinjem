![Pinjem](https://socialify.git.ci/sultanfariz/Pinjem/image?description=1&descriptionEditable=A%20self-initiated%20project%20to%20help%20people%20borrow%20books%20from%20library.%20Developed%20using%20Go%20Echo%20and%20implementing%20Clean%20Architecture.&font=Inter&language=1&logo=https%3A%2F%2Favatars.githubusercontent.com%2Fu%2F55394860%3Fv%3D4&owner=1&pattern=Circuit%20Board&theme=Dark)

<div align="center">
    <img alt="GitHub top language" src="https://img.shields.io/github/languages/top/sultanfariz/Pinjem?style=for-the-badge">
    <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/sultanfariz/Pinjem?style=for-the-badge">
    <img alt="GitHub repo size" src="https://img.shields.io/github/repo-size/sultanfariz/Pinjem?style=for-the-badge">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/sultanfariz/Pinjem?style=for-the-badge">
</div>

### Table of Contents

| [Ruang Lingkup Pengembangan](#ruang-lingkup-pengembangan) | [Entity Relationship Diagram](#entity-relationship-diagram) | [Architectural Patterm](#architectural-pattern) | [Unit Testing](#unit-testing) | [API Postman Docs](#api-postman-docs) | [Swagger OpenAPI Docs](#swagger-openapi-docs) |
| :-------------------------------------------------------: | :---------------------------------------------------------: | :---------------------------------------------: | :---------------------------: | ------------------------------------- | --------------------------------------------- |

## Ruang Lingkup Pengembangan

[`^ kembali ke atas ^`](#table-of-contents)

**Hardware :**

```
- AMD Ryzen 5 2500u
- RAM 4GB DDR4
- AMD Radeon Vega Graphics
- 1TB HDD 128GB SSD
```

**Software :**

```
- MySQL + GORM
- Echo Go
- Docker
- Postman
- Swagger
- Visual Studio Code
```

**Lainnya :**

```
- Waktu pengerjaan +-3 minggu
```

## Entity Relationship Diagram

[`^ kembali ke atas ^`](#table-of-contents)

![ERD](./assets/Pinjemin!-FixERD.png)

## Architectural Pattern

[`^ kembali ke atas ^`](#table-of-contents)

![Clean Architecture](./assets/CleanArch.png)

## Unit Testing

[`^ kembali ke atas ^`](#)

Hasil dari unit test yang telah dilakukan menunjukkan rerata nilai coverage yang berada di atas 80% pada tiap domain bisnis yang ada (6 domain bisnis).

```
ok  	Pinjem/businesses/book_orders	(cached)	coverage: 81.2% of statements
?   	Pinjem/businesses/book_orders/mocks	[no test files]
ok  	Pinjem/businesses/books	(cached)	coverage: 100.0% of statements
?   	Pinjem/businesses/books/mocks	[no test files]
ok  	Pinjem/businesses/deposits	(cached)	coverage: 81.2% of statements
?   	Pinjem/businesses/deposits/mocks	[no test files]
ok  	Pinjem/businesses/orders	(cached)	coverage: 84.2% of statements
?   	Pinjem/businesses/orders/mocks	[no test files]
ok  	Pinjem/businesses/shipping_details	(cached)	coverage: 100.0% of statements
?   	Pinjem/businesses/shipping_details/mocks	[no test files]
ok  	Pinjem/businesses/users	(cached)	coverage: 94.1% of statements
?   	Pinjem/businesses/users/mocks	[no test files]
```

## API Postman Docs

[`^ kembali ke atas ^`](#table-of-contents)

https://documenter.getpostman.com/view/14458184/UUy4e66j

## Swagger OpenAPI Docs

[`^ kembali ke atas ^`](#table-of-contents)

https://app.swaggerhub.com/apis/sultanfariz/Pinjem/1.0.0

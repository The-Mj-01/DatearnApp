```
📦 
├─ .env.example
├─ .gitignore
├─ Dockerfile
├─ README.md
├─ cmd
│  └─ server
│     └─ main.go
├─ go.mod
├─ go.sum
├─ internal
│  ├─ config
│  │  ├─ config.go
│  │  └─ config_test.go
│  ├─ domains
│  │  ├─ bio
│  │  │  ├─ bio.go
│  │  │  ├─ errors.go
│  │  │  ├─ models.go
│  │  │  ├─ repository.go
│  │  │  ├─ repository_test.go
│  │  │  ├─ service.go
│  │  │  ├─ service_test.go
│  │  │  ├─ usecase.go
│  │  │  └─ usecase_test.go
│  │  ├─ image
│  │  │  ├─ errors.go
│  │  │  ├─ image.go
│  │  │  ├─ models.go
│  │  │  ├─ repository.go
│  │  │  ├─ repository_test.go
│  │  │  ├─ service.go
│  │  │  ├─ service_test.go
│  │  │  ├─ usecase.go
│  │  │  └─ usecase_test.go
│  │  ├─ interest
│  │  │  ├─ errors.go
│  │  │  ├─ interest.go
│  │  │  ├─ models.go
│  │  │  ├─ repository.go
│  │  │  ├─ repository_test.go
│  │  │  ├─ service.go
│  │  │  ├─ service_test.go
│  │  │  ├─ usecase.go
│  │  │  └─ usecase_test.go
│  │  ├─ location
│  │  │  ├─ errors.go
│  │  │  ├─ location.go
│  │  │  ├─ models.go
│  │  │  ├─ repository.go
│  │  │  ├─ repository_test.go
│  │  │  ├─ service.go
│  │  │  ├─ service_test.go
│  │  │  ├─ usecase.go
│  │  │  └─ usecase_test.go
│  │  ├─ socialMedia
│  │  │  ├─ errors.go
│  │  │  ├─ models.go
│  │  │  ├─ repository.go
│  │  │  ├─ repository_test.go
│  │  │  ├─ service.go
│  │  │  ├─ service_test.go
│  │  │  ├─ socialMedia.go
│  │  │  ├─ usecase.go
│  │  │  └─ usecase_test.go
│  │  ├─ swipe
│  │  │  ├─ errors.go
│  │  │  ├─ models.go
│  │  │  ├─ repository.go
│  │  │  ├─ repository_test.go
│  │  │  ├─ service.go
│  │  │  ├─ service_test.go
│  │  │  ├─ swipe.go
│  │  │  ├─ usecase.go
│  │  │  └─ usecase_test.go
│  │  └─ user
│  │     ├─ errors.go
│  │     ├─ models.go
│  │     ├─ repository.go
│  │     ├─ repository_test.go
│  │     ├─ service.go
│  │     ├─ service_test.go
│  │     ├─ usecase.go
│  │     ├─ usecase_test.go
│  │     └─ user.go
│  ├─ events
│  │  └─ events.go
│  ├─ middleware
│  │  └─ auth
│  │     └─ auth.go
│  └─ server
│     ├─ api
│     │  └─ api.go
│     └─ http
│        └─ server.go
├─ pkg
│  ├─ advancedError
│  │  ├─ errors.go
│  │  ├─ observer.go
│  │  └─ subject.go
│  ├─ authorization
│  │  ├─ authorization.go
│  │  └─ authorization_test.go
│  ├─ bioHandler
│  │  └─ bioHandler.go
│  ├─ database
│  │  ├─ migrator.go
│  │  └─ mysql.go
│  ├─ hasher
│  │  └─ hasher.go
│  ├─ logger
│  │  ├─ internal
│  │  │  ├─ file
│  │  │  │  ├─ file_logger.go
│  │  │  │  └─ file_logger_test.go
│  │  │  └─ syslog
│  │  │     └─ sys_logger.go
│  │  └─ logger.go
│  ├─ tokenizer
│  │  ├─ auth.go
│  │  ├─ internal
│  │  │  └─ JWT
│  │  │     ├─ jwt.go
│  │  │     └─ jwt_test.go
│  │  └─ tokenUtils.go
│  ├─ userHandler
│  │  └─ userHandler.go
│  ├─ uuid
│  │  └─ uuid.go
│  └─ validation
│     └─ validation.go
└─ scripts
   └─ migrations
      ├─ 00_cities.sql
      ├─ 01_countries.sql
      ├─ 02_sexs.sql
      ├─ 03_images.sql
      ├─ 04_social_media.sql
      ├─ 05_interest.sql
      ├─ 06_bio.sql
      ├─ 07_bio_Interest.sql
      ├─ 08_users.sql
      ├─ 09_likes.sql
      ├─ 10_dislike.sql
      └─ 11_bio_social_media.sql
```
©generated by [Project Tree Generator](https://woochanleee.github.io/project-tree-generator)
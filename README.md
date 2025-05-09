## Requirement
Install go version >= 1.24
```bash
https://go.dev/doc/install
```

## Setup Project

```bash
# project ini dijalankan menggunakan docker compose, service dan database sudah otomatis berjalan

make docker-up

# atau

docker compose up --build
```

## Import Collection
```bash
# project ini menggunakan aplikasi insomnia untuk pengujian api
# import file Insomnia.yaml yang ada didalam folder project
```
![Image](https://github.com/user-attachments/assets/e081997c-dc9a-42b9-a11e-47b2d31ece45)


## API Documentation
```bash
# project ini menggunakan swagger untuk api documentation, silahkan akses 

http://localhost:8080/swagger/index.html
```
![Image](https://github.com/user-attachments/assets/f049a66b-de07-46a8-a72a-7d0af26f7249)
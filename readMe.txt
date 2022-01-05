para configurar heroku:
1) si no funciona el comando en la terminal de Vs Code hacerlo en cmd 
heroku login
2) localizarme en la carpeta de mi proyecto
    C:\Users\Meki\Desktop\WorkSpace>cd socialTwittorBackEnd
3) git init
4) heroku git:remote -a socialtwittorbackend (esto lo saco del proyecto en heroku)

5) antes de poder hacer el deply en heroku
    crear main.go
    https://dashboard.heroku.com/apps/socialtwittorbackend/settings - add buildpack que sea go
    crear go.mod de github con "go mod init github.com/msolivera/socialTwittorBackEnd"
6) ejecutar el siguiente comando para hacer el deploy "git push heroku HEAD:master"

7) instalar otras dependencias que vamos a usar:
    go get go.mongodb.or/mongo-driver/mongo
    go get go.mongodb.org/mongo-driver/mongo/options
    go get go.mongodb.org/mongo-driver/bson
    go get go.mongodb.org/mongo-driver/bson/primitive
    go get golang.org/x/crypto/bcrypt
    go get https://github.com/gorilla/mux
    go get github.com/dgrijalva/jwt-go
    go get github.com/rs/cors

IMPORTANTE PARA QUE FUNCIONEN LAS IMPORTACIONES DE MODULOS LOCALES:
el proyecto debe estar guardado en la carpeta de instalacion de go
En mi caso: C:\Program Files\Go\src\github.com\msolivera\socialTwittorBackEnd
sino hacemos esto NO FUNCIONA

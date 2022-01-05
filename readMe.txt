para configurar heroku:
1) si no funciona el comando en la terminal de Vs Code hacerlo en cmd 
heroku login
2) localizarme en la carpeta de mi proyecto
    C:\Users\Meki\Desktop\WorkSpace>cd socialTwittorBackEnd
3) git init
4) heroku git:remote -a socialtwittorbackend (esto lo saco del proyecto en heroku)

5) antes de poder hacer git push heroku main
    crear main.go
    https://dashboard.heroku.com/apps/socialtwittorbackend/settings - add buildpack que sea go
    crear go.mod de github con "go mod init github.com/msolivera/socialTwittorBackEnd"
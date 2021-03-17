# Things I need to buy

This repository contain my efforts to learn GO. Its a project with the object to build a Telegram bot when you can make list what wanna buy to consult and maintain.

### Things to pay attention on deploy

1. Its a good pratice keep te directory struct with all files configuration outside ```/src``` dir. So, all files ```.go``` inside a dir ```scr/*```
2. To the server know what GO Compiler and Provide need to run, you need to make the go.mod by the command ```go mod init <NameOfModYouCreate>``` (* This a name you want to giv to your module[Module its the name of you project]) on the root path the application.
3. To the server know what vendor lib's need to get to your application we need to run the comand ```go mod vendor``` in the root path of your application and to all of intelicence work in the VSCODE and another IDE's need to run this same command ```go mod vendor```.
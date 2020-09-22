1、如果使用了-F参数，curl会以multipart/form-data的方式发送POST请求。-F以key=value的形式指定要上传的参数，如果是文件，则需要使用key=@file的形式。
curl -F "key=value" -F "filename=@file.tar.gz" http://localhost/upload

2、如果使用-d命令，curl会以application/x-www-url-encoded格式上传参数。
curl -d "action=del" -d "id=12" http://localhost/test

3、如果要以json格式上传参数，需要使用-H在命令中指定。
curl -H "Content-Type: application/json" -X POST -d '{"username":"xyz","password":"xyz"}' http://localhost/api/login
curl 是常用的命令行工具，用来请求 Web 服务器。它的名字就是客户端（client）的 URL 工具的意思。熟练的话，完全可以取代 Postman 这一类的图形界面工具。
curl https://www.example.com，不带有任何参数时，curl 就是发出 GET 请求$ 
curl -X POST https://www.example.com，-X参数指定 HTTP 请求的方法。
curl -u 'bob:12345'，https://google.com/login，-u参数用来设置服务器认证的用户名和密码。
curl -o example.html https://www.example.com，-o参数将服务器的回应保存成文件，等同于wget命令，将www.example.com保存成example.html
curl -I https://www.example.com，-I参数向服务器发出 HEAD 请求，然会将服务器返回的 HTTP 标头打印出来。
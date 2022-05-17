# -*- coding: utf-8 -*-
# 參考資料:趙英傑 超圖解 python 物聯網實作入門 https://swf.com.tw/?p=1129
import socket
#server端
HOST = '0.0.0.0'    #若ip設為0.0.0.0，代表監聽所有連結到本機(控制器)的IP
PORT = 8080         #埠號

s = socket.socket()     # socket()：建立 socket 與設定使用哪種通訊協定，AF_INET是設置地址格式爲IPV4、AF_INET6是設置地址格式爲IPV6，SOCK_STREAM參數表示使用TCP的傳輸方式、SOCK_DGRAM參數表示使用UDP協定的傳輸方式，若不填參數，預設採IPv4.TCP協定
s.bind((HOST, PORT))    #socket.bind(address)：將 socket 綁定到地址，在 AF_INET 下，以 tuple(host, port) 的方式傳入
s.listen(5)             #socket.listen(backlog)：開始監聽 TCP 傳入連接，backlog 指定在拒絕連線前，操作系統可以掛起的最大連接數，該值最少為1，通常設為5就夠用了
print('{}伺服器在{}埠開通了！'.format(HOST, PORT))
client, addr = s.accept()   # socket.accept()：等待前端連線，接受到 TCP 連線後回傳（conn, address），回傳的 conn 是新的 socket 類型，可以用來接收和發送資料，address 是連線客戶端的地址。
print('用戶端位址：{}，埠號：{}'.format(addr[0], addr[1]))

#有一種情況下，字符串一定要轉成二進制，在socket網路編程中，網路傳輸都是以二進制傳輸,所以務必要把str轉成byte(二進制),否則會產生錯誤a bytes-like object is required, not 'str'
while True:
    msg = client.recv(100).decode('utf8')   #socket.recv(bufsize[, flag])：接收 TCPsocket(客戶端) 的資料(每次最多1024位元組)並用UTF-8解碼成字串，Python3 回傳資料為 byte 類型，Python2 回傳資料為 str 類型，buffsize 指定要接收的最大資料量，flag 提供有關消息的其他訊息，通常可以忽略
    print ('收到訊息：' + msg)
    reply = ''

    if msg == '你好':
        reply = b'Hello!'   #回應的訊息也必須是位元組格式
    elif msg == 'bye':      #若訊息是「bye」，則回應b'quit'並跳出迴圈
        client.send(b'quit')#socket.send(string[, flag])：發送 TCP 資料，Python3 是將 byte 發送到已連線的 socket，Python2 是將 str 發送到連線的socket，回傳值是發送的 byte 數
        break
    else:
        reply = b'what??'

    client.send(reply)


client.close()

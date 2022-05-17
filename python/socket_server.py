# -*- coding: utf-8 -*-
# 參考資料:趙英傑 超圖解 python 物聯網實作入門 https://swf.com.tw/?p=1129
import socket
#server端
HOST = '0.0.0.0'
PORT = 8080

s = socket.socket()
s.bind((HOST, PORT))
s.listen(5)
print('{}伺服器在{}埠開通了！'.format(HOST, PORT))
client, addr = s.accept()
print('用戶端位址：{}，埠號：{}'.format(addr[0], addr[1]))

while True:
    msg = client.recv(100).decode('utf8')
    print ('收到訊息：' + msg)
    reply = ''

    if msg == '你好':
        reply = b'Hello!'
    elif msg == 'bye':
        client.send(b'quit')
        break
    else:
        reply = b'what??'

    client.send(reply)


client.close()
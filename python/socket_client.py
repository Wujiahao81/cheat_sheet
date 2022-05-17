import socket
# client端
s = socket.socket()
s.connect(('192.168.1.94', 8080))
while True:
    msg = input('請輸入訊息：')
    s.send(msg.encode('utf8'))
    reply = s.recv(128)
    if reply == b'quit':
        print('關閉連線')
        s.close()
        break
    print(str(reply))


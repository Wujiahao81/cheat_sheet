import socket
# client端
s = socket.socket() # socket()：建立 socket 與設定使用哪種通訊協定，AF_INET是設置地址格式爲IPV4、AF_INET6是設置地址格式爲IPV6，SOCK_STREAM參數表示使用TCP的傳輸方式、SOCK_DGRAM參數表示使用UDP協定的傳輸方式，若不填參數，預設採IPv4.TCP協定
s.connect(('192.168.1.94', 8080))
while True:
    msg = input('請輸入訊息：')   #讀取終端機的文字
    s.send(msg.encode('utf8'))   #把輸入文字轉成UTF-8編碼，再從socket發送出去。
    reply = s.recv(128)          #socket.recv(bufsize[, flag])：接收 TCPsocket(客戶端) 的資料(每次最多128位元組)
    if reply == b'quit':         #若回應為'quit'，則跳出while迴圈(結束程式)。
        print('關閉連線')         #關閉socket
        s.close()
        break
    print(str(reply))           #在終端機顯示回應內容，並把回應的byte類型轉成字串


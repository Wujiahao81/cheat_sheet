import socket
# client端
s = socket.socket()              # socket()：建立 socket 與設定使用哪種通訊協定，AF_INET是設置地址格式爲IPv4、AF_INET6是設置地址格式爲IPv6，SOCK_STREAM參數表示使用TCP的傳輸方式、SOCK_DGRAM參數表示使用UDP協定的傳輸方式，若不填參數，預設採IPv4位址和TCP協定通訊
s.connect(('192.168.1.94', 8080))#連接伺服器，前後端的主機名稱和埠號必須一致，否則無法相連。
while True:
    msg = input('請輸入訊息：')   #讀取終端機的文字
    s.send(msg.encode('utf8'))   #因為在socket網路編程中，網路傳輸都是以二進制傳輸(位元組格式)，所以要先把輸入文字(msg)用encode()轉成UTF-8編碼，再從socket發送出去。
    reply = s.recv(128)          #socket.recv(bufsize[, flag])：接收 TCPsocket(客戶端) 的資料(每次最多128位元組)
    if reply == b'quit':         #若回應為'quit'，則跳出while迴圈(結束程式)。
        print('關閉連線')        
        s.close()                #關閉socket
        break                    #跳出迴圈
    print(str(reply))            #在終端機顯示回應內容，並把回應的byte類型轉成字串


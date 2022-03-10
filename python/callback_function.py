# python回呼函式(callback function)練習
# reference : https://www.youtube.com/watch?v=iUrUf5ijiWs&list=PL-g0fdC5RMboYEyt6QS2iLb_1m7QcgfHk&index=33
def add(n1,n2,cb):
    cb(n1+n2)

def handle(result):
    print("結果是",result)

def handle1(sum):
    print("answer = ",sum)

add(94,87,handle)
add(3,5,handle1)

# 結果是 181
# answer =  8
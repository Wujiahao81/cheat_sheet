# 解法2 function寫法
# 參考來源 : https://gist.github.com/uranusjr/581ba170cc5a42bdd3ff56ede01994ae
max = eval(input("max : "))
min = eval(input("min : "))

def is_prime(max):
    for i in range(min, max):
        if max % i == 0:  # 整除，i 是 max 的因數，所以 max 不是質數。
            return False
    return True     # 都沒有人能整除，所以 max 是質數。

for i in range(min, max + 1):   # 產生 2 到 max 的數字。
    prime = is_prime(i)    # 判斷 i 是否為質數。
    if prime:              # 如果是，印出來。
        print(i)
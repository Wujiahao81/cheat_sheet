# 寫一程式，輸入兩個數字做為範圍最小值和最大值(兩數字介於2~500)，輸出指定範圍內的質數。
# (Hint:質數的定義為因數只有1和自己)
# for 迴圈寫法
max = eval(input("max : "))
min = eval(input("min : "))

for num in range(min,max+1):
    # 質數大於 1
    # if num > 1:
    if max>min and max<=500 and min>=2:
        for i in range(2,num):
            if (num % i) == 0:
                # print(num,"不是質數")
                break
        else:
            print(num,"是質數")
    else:
        print("不在2~500範圍內")
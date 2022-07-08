# 區網內用MAC address找樹莓派的IP地址
很多剛接觸樹莓派的朋友都不知道怎麼在區域網中查找樹莓派的IP地址（一般的內網IP由DHCP隨機分配），但如果你知道樹莓派的MAC地址，這就好辦多了。   
ref : https://codeantenna.com/a/phHQP5Mjbr

```python
import os
res = os.popen('arp -a').readlines() 
for I in res:
    if "b8-27-xx-xx-xx-xx" in I: #windows MAC地址中間用-來分隔
        print(I[:20])
    # if "b8:27:xx:xx:xx:xx" in I: #linux MAC地址中間用:來分隔
    #     print(I[:20])
```

[arp -a](https://blog.downager.com/2013/07/03/%E7%B6%B2%E8%B7%AF-%E6%B7%BA%E8%AB%87-ARP-Address-Resolution-Protocol-%E9%81%8B%E4%BD%9C%E5%8E%9F%E7%90%86/)指令可以在CLI查看IP位址和實體位址的對應關係   
可以用ip link或是ifconfig指令在樹莓派中查詢mac address (wlan0是無線，eth0是有線)
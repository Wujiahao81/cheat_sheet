# Python 定時執行任務 (ASPcheduler)
## What is APScheduler?
以下是來自[官方文檔](https://pypi.org/project/APScheduler/#id1)的介紹
>Advanced Python Scheduler ([APScheduler](https://apscheduler.readthedocs.io/en/3.x/)) is a Python library that lets you schedule your Python code to be executed later, either just once or periodically. You can add new jobs or remove old ones on the fly as you please.  
[Flask-APScheduler](https://viniciuschiele.github.io/flask-apscheduler/index.html) is a Flask extension which adds support for the APScheduler.

- 安裝套件
```python
pip install flask-apscheduler
```
- ASPcheduler中的重要概念  
`Job` 任務  
Job 作為 APScheduler 最小執行單位。創建 Job 時指定執行的函數，函數中所需參數，Job 執行時的一些設置信息。  

- 範例程式
```python
from flask import Flask
# 引用 APSchedule
from flask_apscheduler import APScheduler

# 定時任務
class Config(object):
    JOBS = [
        {
            'id': 'job1',# Job 的唯一ID，注意命名不要衝突
            'func': 'app.main:aps_test', # Job 執行的function 名稱，app.test 就是 app下面的`main.py` 文件，`asp_test` 是方法名稱。文件module和method之間用冒號":"，而不是用英文的"."
            'args': (1, 2), # 如果function需要參數，就在這里添加
            'trigger': {
                'type': 'cron', # 類型，有date、interval、cron三種
                # 'day_of_week': "0-6", # 可定義具體哪幾天要執行
                # 'hour': '*',  # 小時數
                # 'minute': '1',
                'second': '3' # "*/3" 表示每3秒執行一次，單獨一个"3" 表示每分鐘的3秒。現在就是每一分鐘的第3秒時循環執行。
            }
        }
    ]

    # APS(調度器)的API的開關
    SCHEDULER_API_ENABLED = True

# 將a和b相加，用來測試APS的函式
def aps_test(a, b):
    print(a+b)


if __name__ == '__main__':
    app = Flask(__name__)
    
    # 定時任務，導入配置
    app.config.from_object(Config())

    # 進行flask-apscheduler的初始化
    scheduler = APScheduler()
    scheduler.init_app(app)
    scheduler.start()

    app.run()

```

- trigger觸發器設置  
當你開始定時任務時，需要為定時策略選擇一個trigger（設置 class Config 中 trigger 的值）。flask_apscheduler 提供了三種類型的trigger。
```python
"""
date  一次性指定固定時間，只執行一次
interval   間隔調度，隔多長時間一次
cron  指定相對時間執行，比如：每月1號、每星期一執行
"""
# date  最基本的一種調度，指定固定時間，只執行一次
# run_date（str）– 精確時間
class Config(object):
    JOBS = [
        {
            'id': 'job1',                
            'func': '__main__:job1',          
            'args': (1, 2),              
            'trigger': 'date',                        # 指定任務觸發器 date 
            'run_date': '2020-7-23 16:50:00'          # 指定時间 2020-7-23 16:50:00 執行
        }
    ]

    SCHEDULER_API_ENABLED = True
```

```python
"""
interval  通過設置 時間間隔 来運行定時任務

weeks (int) – 間隔幾周
days (int) – 間隔幾天
hours (int) – 間隔幾小時
minutes (int) – 間隔幾分鐘
seconds (int) – 間隔多少秒
start_date (datetime|str) – 間隔計算的起點
end_date (datetime|str) – 結束日期
"""
class Config(object):
    JOBS = [
        {
            'id': 'job1',                
            'func': '__main__:job1',          
            'args': (1, 2),              
            'trigger': 'interval',                        # 指定任務觸發器 interval
            'hours': 5                                     # 每間隔5h執行一次
        }
    ]

    SCHEDULER_API_ENABLED = True
```

```python 
"""
cron 通過設置 相對時間 来運行定時任務

year (int|str) – 年，4位數字
month (int|str) – 月 (範圍1-12)
day (int|str) – 日 (範圍1-31)
week (int|str) – 周 (範圍1-53)
day_of_week (int|str) – 周内第幾天或者星期幾 (範圍0-6 或者 mon,tue,wed,thu,fri,sat,sun，使用名稱可能會報錯，建議使用數字)
hour (int|str) – 時 (範圍0-23)
minute (int|str) – 分 (範圍0-59)
second (int|str) – 秒 (範圍0-59)
start_date (datetime|str) – 最早開始日期/時間(包含)
end_date (datetime|str) – 最晚結束日期/時間(包含)
"""
class Config(object):
    JOBS = [
        {
            'id': 'job1',                
            'func': '__main__:job1',          
            'args': (1, 2),              
            'trigger': 'cron',                            # 指定任務觸發器 cron
            'day_of_week': 'mon-fri',                     # 每周1至周5早上6點執行 
            'hour': 6,
            'minute': 00                                    
        }
    ]
    
    SCHEDULER_API_ENABLED = True
```

## 參考資料
[APScheduler官方文檔](https://apscheduler.readthedocs.io/en/3.x/)  
[Flask-APScheduler官方文檔](https://viniciuschiele.github.io/flask-apscheduler/index.html)  
[在Flask中執行定時任務(作者:迷途小书童的Note)](https://juejin.cn/post/6999817649291001887)  
[flask 使用 flask_apscheduler 做定時循环任務](https://segmentfault.com/a/1190000021245279)  
[Flask-APScheduler 爬坑指南(module not found的解決方案)](https://www.jianshu.com/p/2628f566b31c?u_atoken=594769e5-a129-419d-a665-7e398dd512de&u_asession=01a9YruDu67N1BGgI2H5xEPw3ZduJ-3IVGGJM6e1Z_82r52bXMU6piWNKRiYn-1raKX0KNBwm7Lovlpxjd_P_q4JsKWYrT3W_NKPr8w6oU7K9DZOSciAnCsYLlLE271k_jnHmbkqVcEgdObpAroqY1_GBkFo3NEHBv0PZUm6pbxQU&u_asig=059tO55TWQmGNgVVdK2i_mh2HSglvAvpMC9n4MucpPwQLQ6IJCbJkKc9lZVTK3ymH66M5fQeI5e173iq-EKFnn0j1r84ySwhkBWmrDFRQRp5xDZhAkQEeazvLX9VUIHDls7jAayolQNNF_IcOYAg4jY8_CfRYmg1oidYL1CnPkseT9JS7q8ZD7Xtz2Ly-b0kmuyAKRFSVJkkdwVUnyHAIJzSoDFm_nCmPKOfNsu88Pkj-bzyryXDnJcK7ScIOO55GJ6xbSxAaWh9ph0bRUFW-6vO3h9VXwMyh6PgyDIVSG1W8mfiLo-yz9Xoo-KQpZ55AnW9p7hkxq5Nnkc59BnFNJ_IsEeOrvhD-TkgKw9AK9KUzF_np1q_4Yu5yEbxvnu6z2mWspDxyAEEo4kbsryBKb9Q&u_aref=7NlKP9mrlEqYOndwDn53FO4L9HM%3D)  
[Python 實現定時任務的八種方案！(裡面有關於trigger"觸發器"參數的詳細說明)](https://www.readfog.com/a/1648168828734115840)  
[Python任務调度模块 – APScheduler，Flask-APScheduler实现定時任務(多個APScheduler程式範例)](https://cloud.tencent.com/developer/article/1172218)  
[flask定時框架flask_apscheduler的使用"內有三種trigger設置範例"(轉載)](https://blog.51cto.com/u_15127518/4523657)  
[Flask中使用定时任务以及注意事項(可添加mysql作為定時任務存儲庫)](https://blog.csdn.net/study_in/article/details/106888201)
## 相關連結
這些問題目前沒碰到，但將來可能會碰到所以先放著~  
[APScheduler中两种调度器(BackgroundScheduler和BlockingScheduler)的区别及使用过程中要注意的问题](https://blog.csdn.net/ybdesire/article/details/82228840)  
[Python定時任務神器-APScheduler](https://cloud.tencent.com/developer/article/1506679?from=article.detail.1172218)

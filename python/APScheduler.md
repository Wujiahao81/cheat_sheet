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
            'func': 'app.main:asp_test', # Job 執行的function 名稱，app.test 就是 app下面的`main.py` 文件，`asp_test` 是方法名稱。文件module和method之間用冒號":"，而不是用英文的"."
            'args': (1, 2), # 如果function需要參數，就在这里添加
            'trigger': {
                'type': 'cron', # 类型
                # 'day_of_week': "0-6", # 可定義具體哪幾天要執行
                # 'hour': '*',  # 小時數
                # 'minute': '1',
                'second': '3' # "*/3" 表示每3秒执行一次，单独一个"3" 表示每分鐘的3秒。現在就是每一分鐘的第3秒时循環執行。
            }
        }
    ]

    SCHEDULER_API_ENABLED = True

# 將a和b相加
def asp_test(a, b):
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


## 參考資料
[APScheduler官方文檔](https://apscheduler.readthedocs.io/en/3.x/)  
[Flask-APScheduler官方文檔](https://viniciuschiele.github.io/flask-apscheduler/index.html)  
[在Flask中執行定時任務(作者:迷途小书童的Note)](https://juejin.cn/post/6999817649291001887)  
[flask 使用 flask_apscheduler 做定时循环任务](https://segmentfault.com/a/1190000021245279)  
[Flask-APScheduler 爬坑指南(module not found的解決方案)](https://www.jianshu.com/p/2628f566b31c?u_atoken=594769e5-a129-419d-a665-7e398dd512de&u_asession=01a9YruDu67N1BGgI2H5xEPw3ZduJ-3IVGGJM6e1Z_82r52bXMU6piWNKRiYn-1raKX0KNBwm7Lovlpxjd_P_q4JsKWYrT3W_NKPr8w6oU7K9DZOSciAnCsYLlLE271k_jnHmbkqVcEgdObpAroqY1_GBkFo3NEHBv0PZUm6pbxQU&u_asig=059tO55TWQmGNgVVdK2i_mh2HSglvAvpMC9n4MucpPwQLQ6IJCbJkKc9lZVTK3ymH66M5fQeI5e173iq-EKFnn0j1r84ySwhkBWmrDFRQRp5xDZhAkQEeazvLX9VUIHDls7jAayolQNNF_IcOYAg4jY8_CfRYmg1oidYL1CnPkseT9JS7q8ZD7Xtz2Ly-b0kmuyAKRFSVJkkdwVUnyHAIJzSoDFm_nCmPKOfNsu88Pkj-bzyryXDnJcK7ScIOO55GJ6xbSxAaWh9ph0bRUFW-6vO3h9VXwMyh6PgyDIVSG1W8mfiLo-yz9Xoo-KQpZ55AnW9p7hkxq5Nnkc59BnFNJ_IsEeOrvhD-TkgKw9AK9KUzF_np1q_4Yu5yEbxvnu6z2mWspDxyAEEo4kbsryBKb9Q&u_aref=7NlKP9mrlEqYOndwDn53FO4L9HM%3D)  
[APScheduler中两种调度器(BackgroundScheduler和BlockingScheduler)的区别及使用过程中要注意的问题](https://blog.csdn.net/ybdesire/article/details/82228840),沒碰到但將來可能會碰到先放著
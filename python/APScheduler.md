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
Job 任務  
Job 作為 APScheduler 最小執行單位。創建 Job 時指定執行的函數，函數中所需參數，Job 執行時的一些設置信息。  

- 範例程式
```python
from flask import Flask
# 引用 APSchedule
from flask_apscheduler import APScheduler


class Config(object):
    JOBS = [
        {
            'id': 'job1',# Job 的唯一ID，注意命名不要衝突
            'func': 'app.test:add', # Job 執行的function 名稱


            # 执行任务的function名称，app.test 就是 app下面的`test.py` 文件，`shishi` 是方法名称。文件模块和方法之间用冒号":"，而不是用英文的"."


            'args': (1, 2), # 如果function需要参数，就在这里添加
            'trigger': {
                'type': 'cron', # 类型
                # 'day_of_week': "0-6", # 可定义具体哪几天要执行
                # 'hour': '*',  # 小时数
                # 'minute': '1',
                'second': '3' # "*/3" 表示每3秒执行一次，单独一个"3" 表示每分钟的3秒。现在就是每一分钟的第3秒时循环执行。
            }
        }
    ]

    SCHEDULER_API_ENABLED = True


def add(a, b):
    print(a+b)


if __name__ == '__main__':
    app = Flask(__name__)
    app.config.from_object(Config())

    # 進行flask-apscheduler的初始化
    scheduler = APScheduler()
    scheduler.init_app(app)
    scheduler.start()

    app.run()

```


## 參考資料
[在Flask中執行定時任務(作者:迷途小书童的Note)](https://juejin.cn/post/6999817649291001887)  
[APScheduler —— Python化的Cron](http://www.dannysite.com/blog/73/)  
[Flask-APScheduler 爬坑指南(module not found的解決方案)](https://www.jianshu.com/p/2628f566b31c?u_atoken=594769e5-a129-419d-a665-7e398dd512de&u_asession=01a9YruDu67N1BGgI2H5xEPw3ZduJ-3IVGGJM6e1Z_82r52bXMU6piWNKRiYn-1raKX0KNBwm7Lovlpxjd_P_q4JsKWYrT3W_NKPr8w6oU7K9DZOSciAnCsYLlLE271k_jnHmbkqVcEgdObpAroqY1_GBkFo3NEHBv0PZUm6pbxQU&u_asig=059tO55TWQmGNgVVdK2i_mh2HSglvAvpMC9n4MucpPwQLQ6IJCbJkKc9lZVTK3ymH66M5fQeI5e173iq-EKFnn0j1r84ySwhkBWmrDFRQRp5xDZhAkQEeazvLX9VUIHDls7jAayolQNNF_IcOYAg4jY8_CfRYmg1oidYL1CnPkseT9JS7q8ZD7Xtz2Ly-b0kmuyAKRFSVJkkdwVUnyHAIJzSoDFm_nCmPKOfNsu88Pkj-bzyryXDnJcK7ScIOO55GJ6xbSxAaWh9ph0bRUFW-6vO3h9VXwMyh6PgyDIVSG1W8mfiLo-yz9Xoo-KQpZ55AnW9p7hkxq5Nnkc59BnFNJ_IsEeOrvhD-TkgKw9AK9KUzF_np1q_4Yu5yEbxvnu6z2mWspDxyAEEo4kbsryBKb9Q&u_aref=7NlKP9mrlEqYOndwDn53FO4L9HM%3D)  
[]()
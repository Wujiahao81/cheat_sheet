# Python 定時執行任務 (ASPcheduler)

- 安裝套件
```python
pip install flask-apscheduler
```
- ASPcheduler中的重要概念  
Job 任務  
Job 作為 APScheduler 最小執行單位。創建 Job 時指定執行的函數，函數中所需參數，以及Job 執行時的一些設置信息。
- 範例程式
```python
from flask import Flask
from flask_apscheduler import APScheduler


class Config(object):
    JOBS = [
        {
            'id': 'job1',# Job 的唯一ID
            'func': 'run:add', # Job 執行的函數
            'args': (1, 2),
            'trigger': 'interval',
            'seconds': 3
        }
    ]

    SCHEDULER_API_ENABLED = True


def add(a, b):
    print(a+b)


if __name__ == '__main__':
    app = Flask(__name__)
    app.config.from_object(Config())

    scheduler = APScheduler()
    scheduler.init_app(app)
    scheduler.start()

    app.run()

作者：迷途小书童的Note
链接：https://juejin.cn/post/6999817649291001887
来源：稀土掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
```


## 參考資料
[在Flask中執行定時任務
](https://juejin.cn/post/6999817649291001887)  
[]()  
[Flask-APScheduler 爬坑指南(module not found的解決方案)](https://www.jianshu.com/p/2628f566b31c?u_atoken=594769e5-a129-419d-a665-7e398dd512de&u_asession=01a9YruDu67N1BGgI2H5xEPw3ZduJ-3IVGGJM6e1Z_82r52bXMU6piWNKRiYn-1raKX0KNBwm7Lovlpxjd_P_q4JsKWYrT3W_NKPr8w6oU7K9DZOSciAnCsYLlLE271k_jnHmbkqVcEgdObpAroqY1_GBkFo3NEHBv0PZUm6pbxQU&u_asig=059tO55TWQmGNgVVdK2i_mh2HSglvAvpMC9n4MucpPwQLQ6IJCbJkKc9lZVTK3ymH66M5fQeI5e173iq-EKFnn0j1r84ySwhkBWmrDFRQRp5xDZhAkQEeazvLX9VUIHDls7jAayolQNNF_IcOYAg4jY8_CfRYmg1oidYL1CnPkseT9JS7q8ZD7Xtz2Ly-b0kmuyAKRFSVJkkdwVUnyHAIJzSoDFm_nCmPKOfNsu88Pkj-bzyryXDnJcK7ScIOO55GJ6xbSxAaWh9ph0bRUFW-6vO3h9VXwMyh6PgyDIVSG1W8mfiLo-yz9Xoo-KQpZ55AnW9p7hkxq5Nnkc59BnFNJ_IsEeOrvhD-TkgKw9AK9KUzF_np1q_4Yu5yEbxvnu6z2mWspDxyAEEo4kbsryBKb9Q&u_aref=7NlKP9mrlEqYOndwDn53FO4L9HM%3D)
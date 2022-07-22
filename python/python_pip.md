# pip 指令全記錄

- 安裝套件： 可以將 flask 換成任何想安裝的 Python 套件
```
pip install flask
```
- 更新套件(這方法如果還沒安裝套件會自動安裝，已經安裝的話則對該套件進行更新)：
```
pip install -U flask
```
- 移除安裝過的套件：
```
pip uninstall flask
```
- 安裝並且指定套件版本：
```
pip install -v flask==1.0
```
- 查看目前安裝過的清單：
```
pip list
```
- pip 安裝 requirements.txt 內的清單(一次安裝多個套件)：
```
pip install -r requirements.txt 
```
- 將安裝過的套件建立成 requirements.txt 文件清單，專案若打算要傳給別人使用，這個指令可以表示我們這個專案用哪些模組：
```
pip freeze > requirements.txt
```
- python更新 pip 套件 (可以使用Python命令参数-m選項来安裝，-m的意思是用Python解釋器來運行pip再更新):
```
python -m pip install --upgrade pip
python -m pip install -U pip
```

# 打包python成exe檔
### 教學網址
- [使用auto-py-to-exe完美打包python程序(簡體中文教學)](https://zhuanlan.zhihu.com/p/130328237)
### 問題排解
- [Issues When Using auto-py-to-exe](https://nitratine.net/blog/post/issues-when-using-auto-py-to-exe/?utm_source=auto_py_to_exe&utm_medium=application_link&utm_campaign=auto_py_to_exe_help&utm_content=top)  
- [Bundling data files with PyInstaller (--onefile)](https://stackoverflow.com/questions/7674790/bundling-data-files-with-pyinstaller-onefile/13790741#13790741)

```python
def resource_path(relative_path):
    """Get absolute path to resource, works for dev and for PyInstaller (獲取程序中所需文件資源的絕對路徑)"""
    try:
        # PyInstaller creates a temp folder and stores path in _MEIPASS (PyInstaller創建臨時文件夾，將路徑儲存於_MEIPASS)
        base_path = sys._MEIPASS
    except Exception:
        base_path = os.path.abspath(".")

    return os.path.join(base_path, relative_path)
```
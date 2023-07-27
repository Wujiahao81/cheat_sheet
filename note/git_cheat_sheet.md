# Git cheat sheet

### Tutorial
- [深入淺出講解及實戰GIT教學!](https://www.bilibili.com/video/BV1FE411P7B3)
- [廖雪峰的git教程](https://www.liaoxuefeng.com/wiki/896043488029600)
- [快速上手Git(作者:王铭东)](https://doc.itprojects.cn/0001.zhishi/python.0011.git/index.html#/README)
- [【git】如何寫好的commit訊息](https://youtu.be/g9-X6fR2eYA)
- [初心者懶人包git入門](https://www.maxlist.xyz/2020/05/10/git-tutorial/)
- [Git必懂指令](https://git-tutorial.readthedocs.io/zh/latest/commands.html)
- [Git 還原](https://w3c.hexschool.com/git/1b44e53)
- [【git教學 #1】15分鐘學會git & github（附實例）](https://youtu.be/Zd5jSDRjWfA)(走歪的工程師)
- [拯救資工系學生的基本素養—Github 的 Pull Request 教學](https://youtu.be/MYVMeqj9spw)
- [Learn-Git-in-30-days](https://github.com/doggy8088/Learn-Git-in-30-days)
- [git commit message怎麼寫?](https://chi_gitbook.gitbooks.io/personal-note/content/xue_xi_zi_yuan.html)[Jim Huang / 2015年嵌入式系統課程]
- **GitBook**教學: [深入淺出 GitBook 寫作與自助出版，電子書也能多人協作](https://www.codedata.com.tw/social-coding/gitbook-self-publishing/)

### Step by step.
- Create a repository on [Github](https://github.com).
- Clone the repository to my computer.
```
git clone <https://github.com/account/repo.git>
```
- Start coding...
- Check out what you've done.
```
git status
```
- Check which lines of codes you've changed.
```
git diff
```
- Add file to check list one by one.
```
git add *file1* *file2*
```
Or add all of them.
```
git add .
```
- Add a note when you finished a portion.
```
git commit -m "What you've done."
```
- Upload to github. They will ask for your account/password.
```
git push
```

### When you have a partner who worked at mid night with a cup of coffee...

- Make sure you update the source code in the morning.
```
git pull
```

### Problem fix
- LF will be replaced by CRLF in git - [What is that and is it important? [duplicate]](https://stackoverflow.com/a/5834094)
```
git config core.autocrlf true
```

## git修改未push和已经push的注释信息

1、修改还未push的注释：
git commit --amend
修改后保存退出。

2、刚刚push到远端还没有人其他人下载或改动的：
git commit --amend
进入修改页面修改注释信息，修改后:wq保存退出。
再使用git push --force-with-lease origin master

3、如果其他人已经下载或改动：
git fetch origin
git reset --hard origin/master
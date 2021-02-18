# Git是版本控制系统，Github是在线的基于Git的代码托管服务。

## SVN 和 Git
	- SVN（集中式版本控制系统）：版本库集中存放在中央服务器，需要修改版本的时候，首先从中央服务器哪里得到最新的版本到自己的电脑上，做完的东西再推
	送到中央服务器。【缺点】必须联网工作，如果在局域网还可以（带宽大，速度快），在互联网下如果网速慢的话，就会影响开发交付团队的工作效率；
	- Git（分布式版本控制系统）：没有中央服务器，每个人的电脑就是一个完整的版本库。
	
## Merge 和 Rebase
	- Merge(合并)：用来合并分支，命令会根据具体情况有两种不同的合并方式，快进合并，三方合并。
	- Rebase(变基)：一个经过变基的分支的提交历史是一条没有分叉的直线，尽管实际的开发工作是并行的，但它们看上去就像是先后串行的一样，使得提交历史
	更加整洁，便于查看。
	-| 无论是通过变基，还是通过三方合并，整合的最终结果所指向的快照始终是一样的，只不过提交历史不同罢了。变基是将一系列提交按照原有次序依次应用到
	另一分支上，而合并是把最终结果合在一起。


		

## Github
	- 配置公钥
	- git init 
		- 初始化上传仓库目录
	- git remote add origin https://github.com/ZhaoXinqing/Go_Folder.git
	- 配置远程仓库  
    	- git config --global user.name "ZhaoXinqing"
    	- git config --global user.email 1375565592@qq.com

	- 提交有问题
	- git push origin master -f （舍弃线上的文件，强制推送）
 
	- 分支切换
	- git branch -a  ------------查看所有分支
	- git branch -----------查看当前分支
	- git checkout -b 分支名称  分支路径   --------------切换当前分支


## git commit之后取消的操作使用reset指令进行。
    1.git reset --soft HEAD^，撤销commit，但是不撤销add动作。
    2.git reset --hard HEAD^，撤销commit，并且撤销add动作。
    3.git reset HEAD <文件名>，撤回add动作。
    4.git checkout .，丢弃本次修改内容


## error：failed to push some refs to。 
    这个问题是因为远程库与本地库不一致造成的，那么我们把远程库同步到本地库就可以了。 
    使用指令：git pull --rebase origin master


## 修改还未push的注释：
    git commit --amend
    修改后保存退出。


## commit message 的规范写作：
    1、更好的表达代码意图
    2、加速 code review 速度
    3、基于一定准则的情况下，可以快速筛选特定的 commit message
    4、帮助新同学或者未来的自己，快速建立上下文
    
    fix, fix bug, update,
    http://www.ruanyifeng.com/blog/2016/01/commit_message_change_log.html

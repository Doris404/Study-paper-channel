## Django study note 2: model

#### 数据库

数据库使得一个网页功能更加强大，这也是我们在这篇笔记中提到它的原因，Django
自带默认的数据库是SQLite，如果想用其他的数据库你需要做以下的工作：

* 安装合适的database bindings
* 改变**mysite/settings.py**中的**DATABASES**的一些键值

在这里不过多讲解这部分的知识，具体细节看官方文档https://docs.djangoproject.com/zh-hans/2.2/intro/tutorial02/

#### 模型

在Django里写以一个数据库驱动的web应用的第一步是定义模型，这部分的工作我们会在
**polls/models.py**中完成，我们将在其中完成模型的构造工作

**polls/models.py**
from django.db import models


class Question(models.Model):
    question_text = models.CharField(max_length=200)
    pub_date = models.DateTimeField('date published')


class Choice(models.Model):
    question = models.ForeignKey(Question, on_delete=models.CASCADE)
    choice_text = models.CharField(max_length=200)
    votes = models.IntegerField(default=0)
```
然后我们在命令行中执行下面的代码，来让整个网站知道我们对polls的改变
```python 
python manage.py makemigrations polls
```
结果会是这个样子的
```
Migrations for 'polls':
  polls\migrations\0001_initial.py
    - Create model Question
    - Create model Choice
```

下面我们将利用这个模型制作第一个poll，在命令行运行
```python
python manage.py sqlmigrate polls 0001
```
我们将得到
```
BEGIN;
--
-- Create model Question
--
CREATE TABLE "polls_question" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "question_text" varchar(200) NOT NULL, "pub_date" datetime NOT NULL);
--
-- Create model Choice
--
CREATE TABLE "polls_choice" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "choice_text" varchar(200) NOT NULL, "votes" integer NOT NULL, "question_id" integer NOT NULL REFERENCES "polls_question" ("id") DEFERRABLE INITIALLY DEFERRED);
CREATE INDEX "polls_choice_question_id_c5b4b260" ON "polls_choice" ("question_id");
COMMIT;
```
```python
python manage.py migrate
```
```python
Operations to perform:
  Apply all migrations: admin, auth, contenttypes, polls, sessions
Running migrations:
  Applying polls.0001_initial... OK
```
可以看到我们创建的第一个具体模型0001已经建立成功

改变模型需要以下三步：
* 编辑model.py 改变模型
* 命令行下运行**python manage.py makemigrations** 为模型的改变生成迁移文件
* 命令行下运行**python manage.py migrate**来应用数据库迁移





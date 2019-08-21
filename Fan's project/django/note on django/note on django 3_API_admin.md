## Django study note 3: API 与 Django 管理页面

#### API

我们可以利用API来创建符合模型的对象，命令行下输入
```
$ python manage.py shell
```
进入命令行后我么可以交互式地创建对象

```
In [1]: from polls.models import Choice,Question

In [2]: Question.objects.all()
Out[2]: <QuerySet []>

In [3]: from django.utils import timezone

In [4]: q = Question(question_text="What's new?",pub_date = timezone.now())

In [5]: q.save()

In [6]: q.id
Out[6]: 1

In [7]: q.question_text
Out[7]: "What's new?"

In [8]: q.pub_date
Out[8]: datetime.datetime(2019, 8, 20, 8, 51, 45, 726246, tzinfo=<UTC>)
```
> 更多信息详见： https://docs.djangoproject.com/zh-hans/2.2/topics/db/queries/

#### Django 管理页面

```
$ python manage.py createsuperuser
```
按照指令输入，即可创建一个超级用户。创建一个超级用户之后，即可在网页上进行加入模型
的操作，这简化了网页的建设过程。

**polls/admin.py**
```
from django.contrib import admin

from .models import Question

admin.site.register(Question)
```
**注意**：一定要在polls/admin.py中注册模型，这个模型才可以使用

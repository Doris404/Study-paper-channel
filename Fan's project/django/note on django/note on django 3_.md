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

## Django study note 3: API �� Django ����ҳ��

#### API

���ǿ�������API����������ģ�͵Ķ���������������
```
$ python manage.py shell
```
���������к���ô���Խ���ʽ�ش�������

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
> ������Ϣ����� https://docs.djangoproject.com/zh-hans/2.2/topics/db/queries/

#### Django ����ҳ��

```
$ python manage.py createsuperuser
```
����ָ�����룬���ɴ���һ�������û�������һ�������û�֮�󣬼�������ҳ�Ͻ��м���ģ��
�Ĳ������������ҳ�Ľ�����̡�

**polls/admin.py**
```
from django.contrib import admin

from .models import Question

admin.site.register(Question)
```
**ע��**��һ��Ҫ��polls/admin.py��ע��ģ�ͣ����ģ�Ͳſ���ʹ��

## Note on Django 4 视图模板


视图是网页中最核心的部分，我们可以再views.py中写入代码来更改页面的效果
先看一个简单的例子：

**polls/views.py**
```python
def detail(request, question_id):
    return HttpResponse("You're looking at question %s." % question_id)

def results(request, question_id):
    response = "You're looking at the results of question %s."
    return HttpResponse(response % question_id)

def vote(request, question_id):
    return HttpResponse("You're voting on question %s." % question_id)
```
**polls/urls.py**
```python
from django.urls import path

from . import views

urlpatterns = [
    # ex: /polls/
    path('', views.index, name='index'),
    # ex: /polls/5/
    path('<int:question_id>/', views.detail, name='detail'),
    # ex: /polls/5/results/
    path('<int:question_id>/results/', views.results, name='results'),
    # ex: /polls/5/vote/
    path('<int:question_id>/vote/', views.vote, name='vote'),
]
```
在这个例子中，页面的输出都是固定好的字符串，这显然无法满足需求，为此我们引入模板
来使得输出的页面更为灵活多变

下面看一个更有用的例子：

**polls/views.py**
```python
from django.http import HttpResponse

from .models import Question


def index(request):
    latest_question_list = Question.objects.order_by('-pub_date')[:5]
    output = ', '.join([q.question_text for q in latest_question_list])
    return HttpResponse(output)

# Leave the rest of the views (detail, results, vote) unchanged
```
**polls/templates/polls/index.html**

```
{% if latest_question_list %}
    <ul>
    {% for question in latest_question_list %}
        <li><a href="/polls/{{ question.id }}/">{{ question.question_text }}</a></li>
    {% endfor %}
    </ul>
{% else %}
    <p>No polls are available.</p>
{% endif %}
```
在html文件中，我们定义了页面的模式，通过给这个模型输入具体的数值，可以得到最终的
页面

> 模板语言详见：https://docs.djangoproject.com/zh-hans/2.2/topics/templates/

Django 提供了一个快捷函数render()来完成载入模板，填充上下文，再返回由它生成的
HttpResponse对象这一系列常规操作

**polls/views.py**
```python
from django.shortcuts import render

from .models import Question


def index(request):
    latest_question_list = Question.objects.order_by('-pub_date')[:5]
    context = {'latest_question_list': latest_question_list}
    return render(request, 'polls/index.html', context)
```



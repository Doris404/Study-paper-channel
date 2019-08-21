## Note on Django 4 ��ͼģ��


��ͼ����ҳ������ĵĲ��֣����ǿ�����views.py��д�����������ҳ���Ч��
�ȿ�һ���򵥵����ӣ�

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
����������У�ҳ���������ǹ̶��õ��ַ���������Ȼ�޷���������Ϊ����������ģ��
��ʹ�������ҳ���Ϊ�����

���濴һ�������õ����ӣ�

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
��html�ļ��У����Ƕ�����ҳ���ģʽ��ͨ�������ģ������������ֵ�����Եõ����յ�
ҳ��

> ģ�����������https://docs.djangoproject.com/zh-hans/2.2/topics/templates/

Django �ṩ��һ����ݺ���render()���������ģ�壬��������ģ��ٷ����������ɵ�
HttpResponse������һϵ�г������

**polls/views.py**
```python
from django.shortcuts import render

from .models import Question


def index(request):
    latest_question_list = Question.objects.order_by('-pub_date')[:5]
    context = {'latest_question_list': latest_question_list}
    return render(request, 'polls/index.html', context)
```



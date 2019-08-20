## django study note 

#### ����

����һ���µ���վ
``` 
django-admin startproject mysite
```
��ʱ��õ�һ���Ѿ�����õ��ļ��У�����ֻ��Ҫ��������ļ���Ԥ���趨�Ľṹ��
����������������ǵĴ��룬��������ʵ��һ����վ�Ĺ������������ǵõ����ļ���
�Ľṹ���£�
``` python
mysite/
	manage.py
	mysite/
		__init__.py
		settings.py
		urls.py
		wsgi.py
```
����һ���µ�Ӧ��
```
python manage.py startapp <your app's name>
```
���ǿ�����views.py�����һ����ҳ����ͼ��Ȼ����ƺ���ͼ��������Ҫ�������ӵ�·���ϡ�
�ⲿ������Ҫͨ��urls.py����ɡ�ֵ��ע����ǣ���վ��mywebsite)�������д�����Ŀ����
ʱ�Ѿ���Ȼ��һ��urls.py�ļ��ˣ�������һ���µ�Ӧ��ʱ������ͬ����Ҫ������ļ���������
һ���µ�urls.py��

�������������µ�һ�����Ӱɣ���������������ǵ���վֻ��һ��Ӧ��polls��������ǵ��ļ�
�ṹ���£�
```python
mysite\
	__pycache__\
		__init__.cpython-37.pyc
		settings.cpython-37.pyc
		urls.cpython-37.pyc
		wsgi.cpython-37.pyc
	__init__.py
	settings.py
	urls.py
	wsgi.py

polls\
	__pycache__\
		...
	migrations\
		...
	__init__.py
	admin.py
	apps.py
	models.py
	tests.py
	urls.py
	views.py
db.sqlite3
manage.py
```
���ǿ��Դ������Ƕ��������������ͨ������������һ���µ���ҳ�ģ�
* mysite�ļ�����Ĵ��룺
* polls�ļ�����Ĵ��룺

���ȿ�һ��mysite�ļ��еĴ����������ģ�֮ǰ�ᵽ˵views.py��urls.py�����ǽ�Ϊ��ע���ļ�
������views.py������ͼ��Ч����urls.py������Щ��ҳ���ӵ�һ�������������ֱ�һ����
�����ļ��������������ġ�
\mysite\urls.py
```python 
from django.contrib import admin
from django.urls import path,include

urlpatterns = [
    path('admin/', admin.site.urls),
    path('polls/',include('polls.urls')),
]
```
���Կ��õ���������```path('polls/,include('polls.urls'))``����������ҳ���ӵĹ�����

����������һ��polls�ļ�����Ĵ����������ġ�

polls\urls.py
```
from django.urls import path
from . import views 

urlpatterns = [
    path('',views.index,name = 'index'),#����������views.py�е�index����
]
```
polls\views.py
```
from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.
def index(request):
    return HttpResponse("Hello, world. You are at the polls index.")
```


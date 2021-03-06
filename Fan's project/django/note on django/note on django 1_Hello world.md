## Django study note 1: the first page

#### 入门

创建一个新的网站，在命令行中运行下面的代码（前提是要先下载django的包，可以参考django
官方文档的方法：https://docs.djangoproject.com/zh-hans/2.2/intro/install/）
``` 
django-admin startproject mysite
```
这时会得到一个已经构造好的文件夹，我们只需要按照这个文件夹预先设定的结构，
继续向里面加入我们的代码，即可轻松实现一个网站的构建。现在我们得到的文件夹
的结构如下：
``` python
mysite/
	manage.py
	mysite/
		__init__.py
		settings.py#这个文件里可以更改主站应用的app，每当你加入新的应用时不要忘了改变这里
		urls.py
		wsgi.py
```
创建一个新的应用，依旧是在命令行中输入下面代码
```
python manage.py startapp <your app's name>
```
我们可以在views.py中设计一个网页的视图，然而设计好视图后，我们需要将起连接到路径上。
这部分我们要通过urls.py来完成。值得注意的是，主站（mysite)中在运行创建项目代码
时已经自然有一个urls.py文件了，在生成一个新的应用时，我们同样需要在这个文件夹中设立
一个新的urls.py。

#### 第一个例子

让我们来看以下第一个例子吧，在这个例子中我们的网站只有一个应用polls，因此我们的文件
结构如下：
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
我们可以从两个角度来看我们是如何通过代码来创造一个新的网页的：
* mysite文件夹里的代码：
* polls文件夹里的代码：

首先看一下mysite文件夹的代码是怎样的，之前提到说views.py与urls.py是我们较为关注的文件
，并且views.py负责视图的效果，urls.py负责将这些网页连接到一起。下面我们来分别看一下这
两个文件的内容是怎样的。

**\mysite\urls.py**
```python 
from django.contrib import admin
from django.urls import path,include

urlpatterns = [
    path('admin/', admin.site.urls),
    path('polls/',include('polls.urls')),
]
```
可以看得到，我们用```path('polls/,include('polls.urls'))``来进行了网页连接的工作。

我们再来看一下polls文件夹里的代码是怎样的。

**polls\urls.py**
```
from django.urls import path
from . import views 

urlpatterns = [
    path('',views.index,name = 'index'),#这里引用了views.py中的index函数
]
```
**polls\views.py**
```
from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.
def index(request):
    return HttpResponse("Hello, world. You are at the polls index.")
```

**只进行上述的工作是不够的，还需要在正式运行之前在命令行中执行以下代码**
```python
python manage.py migrate
```
migrate命令会检查**mysite/settings.py**中**INSTALLED_APPS**的设置，当一个
新的应用被定义出来，Django不会立即应用，需要使用这个语句来使得整个网站知道此事。

到目前为止我们实现的仅仅是在django框架下将一个最简单的页面显示出来。
![](1.png "result")




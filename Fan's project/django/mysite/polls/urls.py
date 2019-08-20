from django.urls import path
from . import views 

urlpatterns = [
    path('',views.index,name = 'index'),#这里引用了views.py中的index函数
]
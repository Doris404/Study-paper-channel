## Django study note 2: model

#### ���ݿ�

���ݿ�ʹ��һ����ҳ���ܸ���ǿ����Ҳ����������ƪ�ʼ����ᵽ����ԭ��Django
�Դ�Ĭ�ϵ����ݿ���SQLite������������������ݿ�����Ҫ�����µĹ�����

* ��װ���ʵ�database bindings
* �ı�**mysite/settings.py**�е�**DATABASES**��һЩ��ֵ

�����ﲻ���ི���ⲿ�ֵ�֪ʶ������ϸ�ڿ��ٷ��ĵ�https://docs.djangoproject.com/zh-hans/2.2/intro/tutorial02/

#### ģ��

��Django��д��һ�����ݿ�������webӦ�õĵ�һ���Ƕ���ģ�ͣ��ⲿ�ֵĹ������ǻ���
**polls/models.py**����ɣ����ǽ����������ģ�͵Ĺ��칤��

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
Ȼ����������������ִ������Ĵ��룬����������վ֪�����Ƕ�polls�ĸı�
```python 
python manage.py makemigrations polls
```
�������������ӵ�
```
Migrations for 'polls':
  polls\migrations\0001_initial.py
    - Create model Question
    - Create model Choice
```

�������ǽ��������ģ��������һ��poll��������������
```python
python manage.py sqlmigrate polls 0001
```
���ǽ��õ�
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
���Կ������Ǵ����ĵ�һ������ģ��0001�Ѿ������ɹ�

�ı�ģ����Ҫ����������
* �༭model.py �ı�ģ��
* ������������**python manage.py makemigrations** Ϊģ�͵ĸı�����Ǩ���ļ�
* ������������**python manage.py migrate**��Ӧ�����ݿ�Ǩ��





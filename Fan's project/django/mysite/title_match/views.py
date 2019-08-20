from django.shortcuts import render
from django.http import HttpResponse
from django.template import loader
#from .models import Question
# Create your views here.
def detail(request,question_id):
    return HttpResponse("You are looking at question %s." % question_id)
def results(request,question_id):
    response = "You are looking at the results of question %s."
    return HttpResponse(response % question_id)
def vote(request,question_id):
    return HttpResponse("You are voting on question %s." %question_id)



def index(request):
    latest_question_list = Question.objects.order_by('-pub_date')[:5]
    output = ', '.join([q.question_text for q in latest_question_list])
    return HttpResponse(output)

# Leave the rest of the views (detail, results, vote) unchanged
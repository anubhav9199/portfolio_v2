from . import views
from rest_framework import routers
from django.urls import path, include


router = routers.DefaultRouter()

urlpatterns = [
    path('', include(router.urls)),
    path("v1/resume_data/", views.Portfolio.as_view({"post": "get_resume_data"})),
]
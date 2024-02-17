from . import views
from rest_framework import routers
from django.urls import path, include


router = routers.DefaultRouter()

urlpatterns = [
    path('', include(router.urls)),
    path(
        "api/v1/portfolio",
        views.Portfolio.as_view(),
        name="portfolio_data"
    ),
]

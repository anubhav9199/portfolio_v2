from rest_framework import viewsets
from rest_framework.response import Response
from rest_framework import status
from . import controllers

# Create your views here.
class Portfolio(viewsets.GenericViewSet):

    def get_resume_data(self, request):
        _data = controllers.get_resume_data_controller()
        return Response({"data": _data, "message": "Data fetch Successfully"}, status=status.HTTP_200_OK)

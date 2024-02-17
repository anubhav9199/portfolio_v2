from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.permissions import AllowAny
from rest_framework import status
from . import services

# Create your views here.
class Portfolio(APIView):
    permission_classes = [AllowAny]

    def get(self, request):
        _data = services.get_resume_data_controller()
        return Response({"data": _data, "message": "Data fetch Successfully"}, status=status.HTTP_200_OK)

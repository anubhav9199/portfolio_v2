from fastapi_server.views.portfolio_views import PortfolioView
from main import app


# Write all the route down below
@app.get("/")
def route_read_root():
    return PortfolioView.read_root()


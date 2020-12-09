

from flask import Blueprint, make_response
from flask import current_app as app
from . import service

healthz_bp = Blueprint(
    'healthz_bp',
    __name__,
    template_folder='templates',
    static_folder='static',
)
@healthz_bp.route('/healthz', methods=['GET'])
def healthz():
    result = service.backend_status(app)
    response = make_response()
    if result["content"] != None:
        response = make_response("database_status: {}".format(result["content"]))
    response.content_type = 'application/json'
    response.status_code = result["status_code"]
    return response



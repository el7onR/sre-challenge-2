from flask import Blueprint, render_template, request
from flask import current_app as app
from . import service

users_bp = Blueprint(
    'users_bp',
    __name__,
    template_folder='templates',
    static_folder='static',
)

@users_bp.route('/users', methods=['GET', 'POST'])
def users():
    result = service.get_all_users(app)
    if result != None:
        return render_template(
            'users.html',
            result=result
        )
    if request.method == 'POST':
        user = {}
        user.update({'full_name': request.form.get('full_name')})
        user.update({'username': request.form.get('username')})
        user.update({'email': request.form.get('email')})

        result = service.create_user(app, user)
    
    return render_template(
        'users.html',
        result=result
    )

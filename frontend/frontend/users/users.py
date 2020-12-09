from flask import Blueprint, render_template, request, redirect, url_for
from flask import current_app as app
from . import service

users_bp = Blueprint(
    'users_bp',
    __name__,
    template_folder='templates',
    static_folder='static',
)

@users_bp.route('/users', methods=['GET'])
def users():
    users = service.get_all_users(app)
    if users != None:
        return render_template(
            'users.html',
            users=users,
            new_user=None
        )

@users_bp.route('/users', methods=['POST'])
def new_user():
    user = {}
    user.update({'full_name': request.form.get('full_name')})
    user.update({'username': request.form.get('username')})
    user.update({'email': request.form.get('email')})

    new_user = service.create_user(app, user)
    
    return redirect(url_for('users_bp.users', new_user=new_user))


@users_bp.route('/users/delete', methods=['POST'])
def delete_user():
    user_id = request.form.get('id')
    
    user_deleted = service.delete_user(app, user_id)

    return redirect(url_for('users_bp.users'))

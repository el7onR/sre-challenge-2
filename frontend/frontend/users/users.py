import ast
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
    result = service.get_all_users(app)
    if result["status_code"] == 500 and result["content"] == None:
        return render_template(
            'service_unavailable.html',
            error=result["content"]
        )
    return render_template(
        'users.html',
        users=result,
    )
    

@users_bp.route('/users/create', methods=['POST', 'GET'])
def create_user():    
    if request.method == 'GET':
        return render_template('new_user.html', new_user=None)
    
    user = {}
    user.update({'full_name': request.form.get('full_name')})
    user.update({'username': request.form.get('username')})
    user.update({'email': request.form.get('email')})

    result = service.create_user(app, user)

    if result["status_code"] == 500 and result["content"] == None:
        return render_template(
            'service_unavailable.html',
            error=result["content"]
    )
    
    if result["status_code"] != 201:

        return render_template(
            'new_user.html',
            new_user=result,
        )

    return redirect(url_for('users_bp.users'))


@users_bp.route('/users/delete', methods=['POST'])
def delete_user():
    user_id = request.form.get('id')
    
    result = service.delete_user(app, user_id)

    if result["status_code"] == 500:
        return render_template(
            'service_unavailable.html',
            error=result["content"]
    )

    return redirect(url_for('users_bp.users'))


@users_bp.route('/users/edit', methods=['POST'])
def edit_user():
    user = ast.literal_eval(request.form.get('user'))
    return render_template(
        'edit_user.html',
        user=user,
        error=None
        )

@users_bp.route('/users/update', methods=['POST'])
def update_user():
    user = {}
    user.update({'full_name': request.form.get('full_name')})
    user.update({'username': request.form.get('username')})
    user.update({'email': request.form.get('email')})
    user.update({'id': request.form.get('id')})
    result = service.update_user(app,user)

    if result["status_code"] == 500 and result["content"] == None:
        return render_template(
            'service_unavailable.html',
            error=result["content"]
    )
    
    if result["status_code"] != 200:
        return render_template(
            'edit_user.html',
            user=user,
            error=result
    )

    return redirect(url_for('users_bp.users'))

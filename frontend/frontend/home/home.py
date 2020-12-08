from flask import Blueprint, render_template

home_bp = Blueprint(
    'home_bp',
    __name__,
    template_folder='templates',
    static_folder='static',
    url_prefix="/"
)

@home_bp.route('/', methods=['GET'])
def home():
    return render_template(
        'home.html'
    )

@home_bp.route('/contact', methods=['GET'])
def contact():
    return render_template(
        'contact.html'
    )
    
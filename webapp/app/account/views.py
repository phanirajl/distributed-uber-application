from flask import (Blueprint)

from flask_login import (login_required)

account = Blueprint('account', __name__)


@account.route("/login", methods=["GET", "POST"])
def login():
    pass


@account.route("/logout", methods=["GET", "POST"])
@login_required
def logout():
    pass


@account.route("/register", methods=["GET", "POST"])
def register():
    pass

import random
import string
from flask import Flask, render_template, redirect

app = Flask(__name__)


def guid():  # generate user id
    def g4():
        return ''.join(random.choices(string.ascii_lowercase + string.digits, k=4))

    return g4() + g4() + '-' + g4() + '-' + g4() + '-' + g4() + '-' + g4() + g4() + g4()
    pass


@app.route("/driver", methods=["GET"])
@app.route("/driver/<driver_id>", methods=["GET"])
def driver(driver_id=None):
    return render_template('driver.html', driver_id=driver_id)
    pass


@app.route("/rider", methods=["GET"])
@app.route("/rider/<rider_id>", methods=["GET"])
def rider(rider_id=None):
    # insert into cassandra database
    # handle driver's request
    return render_template('rider.html', rider_id=rider_id)
    pass


@app.route("/")
def main():
    # create user session
    return redirect("/rider/{}".format(guid()))


# sink is incoming driver messages
# source is outgoing rider messages
@app.route("/driver-socket")
def driver_socket():
    # endpoint communicating with kafka server
    pass


# sink is incoming rider messages
# source is outgoing driver messages
@app.route("/rider-socket")
def rider_socket():
    # endpoint communicating with kafka server
    pass


@app.route("/admin")
def admin():
    pass


if __name__ == '__main__':
    app.run(debug=True, host="0.0.0.0")

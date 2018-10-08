from flask import Flask, render_template, redirect, Response


from kafka import KafkaProducer, KafkaConsumer
import time

from utils import guid
app = Flask(__name__)


# initizlize consumers
topic = "riders"
consumer = KafkaConsumer(topic)

bootstrap_servers = ['localhost:9092']

riders_id =[] ## BAD TO STORE ALL INFORMATION ON FLASK_APP CAN CAUSE OVERFLOW

@app.route("/driver", methods=["GET"])
def loginDriver():
    # login driver
    # create session
    # insert to drivers topics

    duid = guid()
    # redirect to /driver/<driver_id> with uuid
    return redirect("/driver/{}".format(duid))
    pass


@app.route("/driver/<driver_id>", methods=["GET"])
def driver(driver_id=None):

    if driver_id is None: # generate driver's ID
        driver_id = guid()

    # create thread safe data structure to consumer messages under topic routes with driver's ID
    msgList = []
    for msg in consumer:
        msgList.append(msg)

    return render_template('driver.html', driver_id=driver_id, uuid=driver_id, response=''.join(msgList))
    pass



@app.route("/rider/<rider_id>", methods=["GET"])
def rider(rider_id=None):
    # insert into cassandra database

    # validate rider_id!!

    if rider_id not in  riders_id: # produce new rider message
        topic = "riders"
        # handle driver's request
        # create kafka-api producer
        producer = KafkaProducer(bootstrap_servers=bootstrap_servers)
        producer.send(topic, bytearray(str(rider_id),'utf-8'))
        time.sleep(1)
        print("published message {}".format(rider_id))
        riders_id.append(rider_id)

    return render_template('rider.html', uuid=rider_id)
    pass



@app.route("/")
@app.route("/rider", methods=["GET"])
def main():


    uid = guid() # unique user id

    # create flask session with uid

    return redirect("/rider/{}".format(uid))


# sink is incoming driver messages
# source is outgoing rider messages
@app.route("/driver-socket")
def driver_socket():
    # endpoint communicating with kafka-api-start server

    pass


# sink is incoming rider messages
# source is outgoing driver messages
@app.route("/rider-socket")
def rider_socket():
    # endpoint communicating with kafka-api-start server
    pass


@app.route("/admin")
def admin():
    pass


if __name__ == '__main__':
    app.run(debug=True, host="0.0.0.0", threaded=True)

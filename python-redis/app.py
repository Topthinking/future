from flask import Flask
from redis import Redis
import os

app = Flask(__name__)
redis = Redis(host=os.environ.get("REDIS_HOST","127.0.0.1"),port=6379)


@app.route("/")
def hello():
	redis.incr("hits")
	return '访问次数: %s \n' % (redis.get("hits")) 

if __name__ == "__main__":
		app.run(host="0.0.0.0", port=8000, debug=True)
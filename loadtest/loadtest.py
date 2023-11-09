from locust import HttpUser, task, between
from locust.env import Environment
from locust.stats import stats_printer, stats_history
from locust.log import setup_logging
import gevent

class QuickPostUser(HttpUser):
    wait_time = lambda self: 0
    host = "http://localhost:8080"

    @task
    def post_url(self):
        payload = {'url': 'https://www.example.com'}
        self.client.post("/shorten", json=payload)

# Setup Locust logging
setup_logging("INFO", None)

# Create an Environment with the host parameter
env = Environment(user_classes=[QuickPostUser], host="http://localhost:8080")
env.create_local_runner()

# Start a greenlet that periodically outputs the current stats
gevent.spawn(stats_printer, env.stats)

# Start a greenlet that saves stats to history
gevent.spawn(stats_history, env.runner)

# Start the test
env.runner.start(user_count=500, spawn_rate=100)

# In 60 seconds stop the runner
gevent.spawn_later(60, lambda: env.runner.quit())

# Wait for the greenlets
env.runner.greenlet.join()

# # Print the stats in the end
# env.stats.log_all()
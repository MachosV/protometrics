import sys, time
class ThreadManager:
    
    threads = []

    @classmethod
    def subscribe(cls, t):
        cls.threads.append(t)
    
    @classmethod
    def unsubscribe(cls, t):
        cls.threads.remove(t)
    
    @classmethod
    def terminate(cls):
        print "ThreadManager terminating"
        for t in cls.threads:
            t.stop()

    @classmethod
    def waitForFinish(cls):
        print "ThreadManager Waiting for finish"
        while len(cls.threads) > 0:
            time.sleep(0.5)
        print "ThreadManager Finished"
        sys.exit(42)
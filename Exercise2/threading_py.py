from threading import Thread, Lock
i = 0
def thread_function_1(lock):
	global i
	for j in range(0, 1000000):
		lock.acquire()
		i = i + 1
		lock.release()
def thread_function_2(lock):
	global i;
	for j in range(0, 1000000):
		lock.acquire()
		i = i - 1
		lock.release()
def main():
	lock=Lock()
	thread_1 = Thread(target = thread_function_1, args = (lock))
	thread_2 = Thread(target = thread_function_2, args = (lock))
	thread_1.start()
	thread_2.start()
	thread_1.join()
	thread_2.join()
	print(i)
main()



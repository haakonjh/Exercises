from threading import Thread, Lock
i = 0
def thread_function_1(NULL,lock):
	global i
	for j in range(0, 1000000):
		lock.acquire(True)
		i = i + 1
		lock.release()
def thread_function_2(NULL,lock):
	global i;
	for j in range(0, 1000000):
		lock.acquire(True)
		i = i - 1
		lock.release()
def main():
	lock=Lock()
	NULL=0
	thread_1 = Thread(target = thread_function_1, args = (NULL,lock))
	thread_2 = Thread(target = thread_function_2, args = (NULL,lock))
	thread_1.start()
	thread_2.start()
	thread_1.join()
	thread_2.join()
	print(i)
main()



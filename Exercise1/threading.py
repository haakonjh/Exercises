from threading import Thread

i = 0
def thread_1_function():
	for x in range(0,1000000):
		i+=1;
def thread_2_function():
	for x in range(0,1000000):
		i-=1;

# Potentially useful thing:
# In Python you "import" a global variable, instead of "export"ing it when you declare it
# (This is probably an effort to make you feel bad about typing the word "global")
def main():
	thread_1 = Thread(target = thread_1_function, args = (),)
	thread_2 = Thread(target = thread_2_function, args = (),)

	thread_1.start()
	thread_2.start()

	thread_1.join()
	thread_2.join()

main()

#include <pthread.h>
#include <stdio.h>
// Note the return type: void*
//gcc -std=gnu99 -Wall -g -o threading_c threading.c -lpthread
static int i=0;
pthread_mutex_t lock;
void* thread_1_function(){
	
	for(int j=0;j<=1000000;j++)
	{
	pthread_mutex_lock(&lock);
		i++;
	pthread_mutex_unlock(&lock);
	}	
	
	return NULL;
}

void* thread_2_function(){
	
	for(int j=0;j<=1000000;j++)
	{
	pthread_mutex_lock(&lock);
		i--;
	pthread_mutex_unlock(&lock);
	}
		
	return NULL;
}

int main(){
	
	
	pthread_t thread_1;
	pthread_create(&thread_1, NULL, thread_1_function,NULL);
	pthread_t thread_2;
	pthread_create(&thread_2, NULL, thread_2_function,NULL);
	
	pthread_join(thread_1, NULL);
	pthread_join(thread_2, NULL);
	printf("%i",i);
	pthread_mutex_destroy(&lock);
	return 0;
}

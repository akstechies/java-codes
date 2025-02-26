# forloop.py
import time

start_time = time.time()

for  i in range(1, 500000):
    # print(i)
    pass

end_time = time.time() 
print(f"Time taken: {end_time - start_time} seconds")
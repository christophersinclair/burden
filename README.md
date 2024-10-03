```
     __                   __                
    / /_  __  ___________/ /__  ____          
   / __ \/ / / / ___/ __  / _ \/ __ \
  / /_/ / /_/ / /  / /_/ /  __/ / / /
 /_.___/\__,_/_/   \__,_/\___/_/ /_/ 
                               
```
                                    
Burden is an automated performance testing tool that will sustain a specified CPU load over time.

It is designed to be run within a Kubernetes cluster, but can be used on any system provided you have set the correct environment variables.



### How it Works
This code will constantly calculate the nth number of the Fibonacci sequence on all cores assigned to the program.


### Altering Behavior
You can get different behavior by modifying the following in `deploy.yaml` (or the equivalent on a non-Kubernetes system):
- The value of `replicas`: determines how many Pods will be scheduled that is executing this code. Increase for more parallel code executions
- The value of `FIB`: determines which position of the Fibonacci sequence to calculate (e.g., `FIB=7` would calculate the 7th number in the sequence, and return 8). The higher the number, the more intense the calculation running on a particular thread
- The value of `MAX_THREADS`: determines how many Goroutines to spawn and sustain. NOTE: the program will constantly add more Goroutines up to this value and keep spawning more as some complete their calculation
- The value of CPU requests and limits (Kubernetes only): by default this is set to 0 to maximize the underlying node CPU consumption, but can tweak it to find the impact of setting requests and limits

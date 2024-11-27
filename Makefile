RatInMaze: RatInMaze.o
	g++ -g -o RatInMaze.exe RatInMaze.o -pthread    

RatInMaze.o: RatInMaze/RatInMaze.cpp
	g++ -g  -c -pthread RatInMaze/RatInMaze.cpp

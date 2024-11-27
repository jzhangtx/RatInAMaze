Rat In A Maze (please see the docx version for a better understanding)


You are given a maze in the form of a matrix of size n * m. Each cell is either clear or blocked denoted by 1 and 0 respectively.
A rat sits at the top-left cell and there exists a block of cheese at the bottom-right cell. Both these cells are guaranteed to be clear.
You need to find if the rat can get the cheese if it can move only in one of the two directions - down and right. It can’t move to blocked cells.


Testing
Input Format
The first line contains an integer ‘T’, denoting the number of test cases.

For each test case the input has the following lines:

The first line contains two space-separated integers ‘n’ and ‘m’ denoting the number of rows and columns of the matrix respectively.
The next n lines contain m space-separated integers which are either 0 or 1.
Output Format 
For each test case, a line containing 1 or 0 based on whether the rat can get the cheese or not respectively.

Sample Input
4
4 4
1 1 0 0
1 1 1 1
0 1 0 1
0 1 0 1
4 4
1 1 0 0
1 1 1 1
0 1 1 0
0 1 0 1
4 5
1 0 1 1 1
1 1 1 0 1
0 1 0 0 1
0 1 1 0 1
3 4
1 0 0 0
0 0 0 0
0 0 1 1
Expected output
1
0
0
0
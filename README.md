**The Small Transport Company**

**Problem**

You are the owner of a small transport company.
You own 1 large bus that can transport a maximum of B passengers, but you can also
sub rent a set of small cars. Each small car can transport up to C passengers at a time.
You need to transport passengers from K different stations to a football stadium.
Elements of vector STATIONS describe the number of passengers waiting to be transported
from each station.

**Input**
B = 12         Bus can transport maximum 12 passengers.
C = 5          Car can transport maximum 5 passengers.
STATIONS = [15, 10]  There are 15 and 12 passengers at two stations.

**Task**
Find the minimal number of small cars you need to rent to transport all
passengers. Passengers need to be transported at the same time (no reuse of
vehicles).

**Constraints**
B, C, STATIONS[i] are integers
1 <= B <= 999
1 <= C <= 1000
1 <= length(STATIONS) <= 100
0 <= STATIONS[i] <= 1000

**Examples**
STATIONS = [15, 10]
B = 12
C = 5
Returns: 3 cars needed

STATIONS = [11, 15, 13]
B = 9
C = 5
Returns: 7 cars needed

# MeetingsScheduler

This is an alpha checkin

## Logic.

 1. Create groups of unique combinations of meetings that do not share participants. The cumulative meeting pool is updated for every addition. Thus the meetings added are unique with respect to the entire group

 2. A 2D array is built for every group in dim1 and the common meetings against every other group in dim2

 3.  The group in dim1 with the least cumulative common meetings of dim2 is selected for a slot as,
 this group in dim1 has the largest number of attendees across groups*, and its meetings are then
 excluded from further consideration. So the next group with a low common overlap is selected.



 4. Repeated from 2 until all meetings are scheduled into different slots.

* These attendees are attending max number of meetings in cumulative terms


# Run

make test

# Results

## Meetings with participants indexed from 0


```
meetings := {"A", "E"},         //0
            {"B", "F"}, 
            {"C", "G"}, 
            {"D", "H"}, 
            {"B", "C", "D"}, 
            {"A", "C", "D"}, 
            {"A", "B", "D"}, 
            {"A", "B", "C"}     //7

schedule::
Num of meetings= 8
Slot[0]=3 7
Slot[1]=2 6
Slot[2]=1 5
Slot[3]=0 4


meetings := 
    {"A0", "B1"},                                                                   //0
    {"A1", "B0", "C2"},
    {"B2", "C1", "E0", "F1"},
    {"C0", "D1", "F2", "G2", "H1"},
    {"D0", "E1", "G0", "H2", "I1", "J0"},
    {"E2", "F0", "I0", "J1", "A2", "B1", "C2"},
    {"G1", "H0", "I2", "J2", "A0", "B2", "C0", "D2"},
    {"A1", "B0", "C1", "D0", "E2", "F1", "G2", "H1", "I0"},
    {"J0", "A2", "B1", "C2", "D1", "E0", "F2", "G0", "H2", "I1"},
    {"J1", "A0", "B2", "C0", "D2", "E1", "F0", "G1", "H0", "I2", "J2"},
    {"A1", "B0", "C1", "D0", "E2", "F1", "G2", "H1", "I0", "J0", "A2", "B1", "C2"},
    {"D1", "E0", "F2", "G0", "H2", "I1", "J1", "A0", "B2", "C0", "D2", "E1", "F0", "G1", "H0"},
    {"B1", "C2", "D1", "E0", "F0", "G1", "H2", "I2", "J0", "A1"},
    {"C0", "D2", "E1", "F2", "G0", "H0", "I0", "J1", "A2", "B0", "B2"},
    {"D0", "E2", "F1", "G2", "H1", "I1", "J2", "A0", "B1", "C1"},
    {"E0", "F0", "G0", "H0", "I0", "J0", "A1", "B2", "C2", "D1", "E1"},
    {"F1", "G1", "H1", "I1", "J1", "A2", "B0", "C0", "D2", "E2", "F2"},
    {"G2", "H2", "I2", "J2", "A0", "B1", "C1", "D0", "E0", "F0", "G0"},
    {"H1", "I1", "J1", "A1", "B2", "C2", "D1", "E1", "F1", "G1", "H2", "I2", "J2"},
    {"I0", "J0", "A0", "B0", "C0", "D0", "E2", "F2", "G2", "H0", "I1", "J1", "A1", "B1", "C1"}

Num of meetings= 20
Slot[0]=19
Slot[1]=12 13
Slot[2]=10 6
Slot[3]=7 8
Slot[4]=0 18
Slot[5]=1 17
Slot[6]=14
Slot[7]=11
Slot[8]=9
Slot[9]=15 16
Slot[10]=4
Slot[11]=2 3 5


```



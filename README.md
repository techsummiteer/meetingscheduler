# MeetingsScheduler

This is an alpha checkin for sharing code with experts.

## Logic.

 1. Create groups of unique combinations of meetings that do not share participants. The cumulative meeting pool is updated for every addition. Thus the meetings added are unique with respect to the entire group

 2. A 2D array is built for every group in dim1 and the common meetings against every other group in dim2

 3. The group in dim1 with the least cumulative common meetings of dim2 is selected for a slot as, this group in dim1 has the largest number of attendees, and its meetings are then excluded from further consideration. So the next group with a low common overlap is selected.

 4. Repeated from 2 until all meetings are scheduled into different slots.

# Run

The test meetings are in-code. The make also runs the code. ( lazy me )

'
make
'
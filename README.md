# MeetingsScheduler

This is an alpha checkin for sharing code with experts.

## Logic.

1. Groups of unique combinations of meetings that do not share participants. The cumulative meeting pool is updated for every meeting addition. The meetings added are unique with respect to the group

2. A 2D array is built for every group in dim1 and the highest common factor of meetings in dim2

3. Group with the least HCF is selected for a slot as, this has the largest number of attendees, and its meetings are then excluded from further consideration.

4. Repeated from 2 until all meetings are scheduled into different slots.


# Run

The test meetings are in-code. The make also runs the code. ( lazy me )

'
make
'
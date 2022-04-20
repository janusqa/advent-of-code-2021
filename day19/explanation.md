# Mental Notes

Between some scanners their will be at most 12 beacons that are shared.

These beacons have a special property. This property is as described below
Let there be two scanners s1, s2 and let there be two becons b1, b2 that s1, and s2 can see
from s1 perspective b1 and b2 will look like s1b1x, s1b1y, s1b1z and likewise s1b2x, s1b2y, s1b2z
from s2 perspective similarly b1 and b2 will look like s2b1x, s1b2y, s2b1z and likewise s2b2x, s2b2y, s2b2z

if we take s1 as our base and we face and rotate s1 in a particular direction as shown in plot.png
something magical happens. Any beacon seen by both, if we measure the distance beween thier perspectives
will, that distance will all be the same.

For example 
d1x, d1y, d1z = s1b1x - s2b1x, s1b1y - s2b1y, s1b1z - s2b1z

and if we do same for second beacon then

d2x, d2y, d2z = s1b2x - s2b2x, s1b2y - s2b2y, s1b2z - s2b2z

for both thes two beaons d1 = d2.  As per problems there will be at least 12 of these between two scannears
when one scanner is rotated and faced in right way.

if we get 12 of these common offsets in a specific orientation of s2, then we know we have faced
and rotated s2 correctly.

If we get 12 of these we know these are the same beacons but seen from two perspectives of s1 and s2

That common offset by the way is also the position of s2 from s1 perspective. that is if we were to move
s2 by that offset it will end up exactly in the position of s1. 

So now we can add this offset to all beacons of s2 when in this special orientation. This will be the positions
of those beacons from s1 perspective.

##Exams from previous years

I have posted exams and solutions from previous years on It's learning. Note that we used a different book and different lab project in 2012, so it may be somewhat less relevant.

(Posted by Hein: Tuesday 21 Oct)

##Use of inline assembly in gcc

```
//make a global lock variable
int asmMutex  = 1;
 
//To lock
asm("spin: lock btr $0, asmMutex ");
asm(" jnc spin ");
.
//To unlock
asm(" bts $0, asmMutex ");
```

##Use of commandline arguments i C
To get commandline arg into c try this:

``` 
#include<stdio.h>
#include<stdlib.h>
 
 
int main( int argc, char *argv[] )
{
        int i;
        printf("Number of arguments entered=%d\n", argc);
        printf("And the arguments are:\n");
        for(i=0;i<argc;i++)
                printf("Arg %d=%s\n", i, argv[i] );
        return 0;
}
``` 
Usage:
```
mortenm@pitter23:~/OpSys/2013/cmdline$ ./main 1 heisan hoppsan
Number of arguments entered=4
And the arguments are:
Arg 0=./main
Arg 1=1
Arg 2=heisan
Arg 3=hoppsan
mortenm@pitter23:~/OpSys/2013/cmdline$
``` 
PS: All arguments are 'strings' (char[]), and may need to be converted to in,float, etc. before furter usage.
Use atoi(), atof() etc....



##Share folder between host and guest on VirtualBox
It is possible to share folders between the host (the machine where the VirtualBox application is running on) and the virtual image.
 
You first need to set up a folder on the host (settings->shared folders) and give the share a name. 
Inside SlackWare you need to run the following command to mount the share:
 
Make a dir where you do the mount:
mkdir share (or another name you choose )
sudo mount -t vboxsf share ¨/share (where the first 'share' is the name given in settings->shared folder, and ~/share is the directory just created)
PS: mount is reserved to 'root' to execute. But by use of sudo we can run 'root' commands as 'user'. Use man sudo for more info


##Lab handin deadlines postponed

Due to the problems with github, we will be postponing the deadline for lab 2 until Sunday 14th, and likewise for the next labs. Please consult the lecture plan for the details. Tomorrow I'll walk through a few examples with git and github to explain a few things. I have also found a new room which should also have better space and facilities. For those who read this, please let others know that we will meet in E458 (across the hall from the eletro-lab).

(Posted by Hein: Sunday 7 Sep)

##Lab repositories update

If you find your repo for lab2 missing, it has most likely been deleted (the remaining will be deleted soon). The reason for deleting, see below.

(Posted by Hein: Friday 5 Sep)

##Lab repositories

Hey all. Due to a misunderstanding of github's instructions, I have created too many repositories (one per student per lab exercise). This does not scale, and yesterday we ran out of private repositories. Thus, we will instead be creating only one repository for all lab exercises, with lab2 and lab3 etc as subfolders. Thus, those who have already been given a repo for lab2, this will be deleted. If you have made changes locally related to lab2, please take care to keep your local copy safe, so that you can move those changes over to the new copy later. I will start to create new repositories soon, and I'm asking that you please begin using the new `username-labs` repos.

(Posted by Hein: Thursday 4 Sep)

##Lab 2 will be private

**We will need your githut account name.** You should not fork the lab2 repo, when you get access to it. Instead we will create a separate repo for each of you. To give us your github account name, clone this repo, and add your account name to the file `students`, and do a *Pull request.* (see the README.md description for lab1 for instructions related to doing a Pull request.)

(Posted by Hein: Tuesday 2 Sep)

##Lab 1 Handin

If you experience trouble with submitting lab 1 due to difficultites with git, don't worry. We will not write down any slip days due to this.

(Posted by Hein: Tuesday 2 Sep)

##External access to the Linux machines

Information about how to do a remote login to the Linux machines can be found here: http://wiki.ux.uis.no/foswiki/Info/WebHome and here: http://wiki.ux.uis.no/foswiki/Info/HvordanLoggeInnP%E5Unix-anlegget

##Important: Sign up for Unix account

To be effective in the lab on Tuesday, it is important that you sign up for an account on the Unix system prior to coming to the lab. The latest time to sign up is **Monday at 15:00**, in order to have your account ready for the lab on Tuesday. But there is no reason to delay; do it now! The instructions for sign up are here http://user.ux.uis.no/.

(Posted by Hein: Friday 22 Aug)

##Lab schedule update

Turns out that the lab room has a scheduling conflict so we won't have lab on Mondays. I've requested that the schedule be updated. You can choose when you come to the lab, but we, the teaching staff, will only be there at some given time (to be announced later). Attendance is not mandatory. Beyond that you can be in the lab when it is reserved for the course or when it is not being used by another course. 


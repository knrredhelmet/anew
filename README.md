## Installation
```
git clone https://github.com/knrredhelmet/anew.git
cd anew
go build anew.go
```

Here, a file called `things.txt` contains a list of numbers. `newthings.txt` contains a second
list of numbers, some of which appear in `things.txt` and some of which do not. `anew` is used
to append the latter to `things.txt`.


```
▶ cat things.txt
Zero
One
Two

▶ cat newthings.txt
One
Two
Three
Four

▶ cat newthings.txt | anew things.txt
Three
Four

▶ cat things.txt
Zero
One
Two
Three
Four

cat newthings.txt | anew things.txt -o added-lines.txt
▶ cat added-lines.txt
Three
Four
```

## Inspired
[tomnomnom/anew](https://github.com/tomnomnom/anew)

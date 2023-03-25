# Forum #

This project consists in creating a web forum that allows :

    communication between users.
    associating categories to posts.
    liking and disliking posts and comments.
    filtering posts.

### How to go run? ####

firstly u need to enter newBackend dir
```
go run main.go
```
secendly u need exit newBackend dir and to enter client dir
```
npm install
```
```
npm start
```
on this moment u start forum on 3000 port(front) and 8080 port(back)

### How to use ###

you must first sign up, secend move is log in with your login and password. Now u can create post comment other person post, like and dislike. Take a filter to post.

### Audit link ###

https://github.com/01-edu/public/tree/master/subjects/forum/audits


### If u wanna use makefile and docker use this command###
```
make dcrun
```
wait the end client image creating
```
make dbrun
```
wait the end back image creating

# go to ling localhost:3000/ #

if u wanna delete all image use this command
```
make dstop
```


```
make dclear
```


### Project made by ###
@aromanov DevOps

@diyar.ildart backend

@aseitkhan frontend



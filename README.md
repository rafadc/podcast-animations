# Animations for podcast

This project is a generator for different animations to be used in podcasts I am recording

It creates a `tmp` folder for the frames and saves the video in `output` folder.

## Cover and author

```shell
animate cover-and-author background.jpg cover.jpg author.jpg 
```


![sample animation with good strategy/bad strategy](doc/book-and-author.gif)

# Video format

At the moment the video exported is hardcoded to h264 with color space yuv420p so free version of Davinci Resolve can open the output files.
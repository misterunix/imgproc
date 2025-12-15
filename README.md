# imgproc

imgproc is a collection of routines that convert image.RGBA into:

- black & white
  - simple
  - bt.601
  - bt.709
  - bt.2020
  - luminance
  - gamma is not fully supported at this time

- ditering - if the image is color, the red value is used and only a grey scale is saved
  - random
  - simple
  - stucki
  - sierra
  - Jarvis-Judice-Nink
  - Floyd-Steinberrg
  - Burks
  - Arkinson
  - unknown1 - I cant remember where I found the error diffusion pattern for this routine

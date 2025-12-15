# imgproc

imgproc is a collection of routines that convert image.RGBA into:

- black & white
  - simple
  - bt.601
  - bt.2020

  - bt.709
  - luminance
    - both bt.709 and luminance are the same

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

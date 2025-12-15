# imgproc

## Ive created an error. Project on hold till I can figure it out.

imgproc is a collection of routines that convert image.RGBA into:

- black & white
  - :heavy_check_mark: simple
  - :heavy_check_mark: bt.601
  - :heavy_check_mark: bt.2020

  - :heavy_check_mark: bt.709
  - :heavy_check_mark: luminance
    - both bt.709 and luminance are the same

  - gamma is not fully supported at this time

- ditering - if the image is color, the red value is used and only a grey scale is saved
  - :heavy_check_mark: random
  - :heavy_check_mark: simple
  - :heavy_check_mark: stucki
  - :heavy_check_mark: sierra1
  - :heavy_check_mark: sierra2
  - :heavy_check_mark: sierra2 4a
  - :heavy_check_mark: Jarvis-Judice-Nink
  - :heavy_check_mark: Floyd-Steinberrg
  - :heavy_check_mark: False Floyd-Steinberrg
  - :heavy_check_mark: Burks
  - :heavy_check_mark: Arkinson
  - :heavy_check_mark: unknown1 - I cant remember where I found the error diffusion pattern for this routine

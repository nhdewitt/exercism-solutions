module Darts (score) where

score :: Float -> Float -> Int
score x y
  | coord > 25 && coord <= 100  = 1
  | coord > 1 && coord <= 25    = 5
  | coord <= 1                  = 10
  | otherwise                   = 0
  where coord = ((x^2) + (y^2))

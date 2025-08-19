module CollatzConjecture (collatz) where

chain :: (Integral a) => a -> [a]
chain 1 = []
chain n
  | even n    = n:chain (n `div` 2)
  | odd n     = n:chain (n*3 + 1)

collatz :: Integer -> Maybe Integer
collatz n
  | n <= 0     = Nothing
  | otherwise  = Just (fromIntegral (length (chain n)) :: Integer)

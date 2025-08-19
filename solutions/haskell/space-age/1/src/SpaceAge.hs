module SpaceAge (Planet(..), ageOn) where

data Planet = Mercury
            | Venus
            | Earth
            | Mars
            | Jupiter
            | Saturn
            | Uranus
            | Neptune

ageOn :: Planet -> Float -> Float
ageOn planet seconds =
  case planet of
    Mercury  -> seconds / (earthSeconds * 0.2408467)
    Venus    -> seconds / (earthSeconds * 0.61519726)
    Earth    -> seconds / earthSeconds
    Mars     -> seconds / (earthSeconds * 1.8808158)
    Jupiter  -> seconds / (earthSeconds * 11.862615)
    Saturn   -> seconds / (earthSeconds * 29.447498)
    Uranus   -> seconds / (earthSeconds * 84.016846)
    Neptune  -> seconds / (earthSeconds * 164.79132)
    _        -> error "Sorry, Pluto."
  where
    earthSeconds = 31557600
# enigma

This is a Go implementation of an Enigma machine

## Design

### Creating an Enigma machine

The parameters needed to create a new instance of an enigma machine are:

Rotors:

- Rotors were marked with Roman numerals to distinguish them, from I to VIII
- Allowed values: "I", "II", "III", "IV", "V", "VI", "VII", "VIII"
- 3 rotors are used

Starting position:

- Each rotor can be set to one of 26 possible starting positions when placed in an Enigma machine
- Allowed values: 1 to 26
- 3 rotors starting positions are used

Ring setting:

- Each rotor contains one or more notches that control rotor stepping. In the military variants, the notches are located
  on the alphabet ring.
- Allowed values: 1 to 26
- 3 rotors ring settings are used

Plugboard cable:

- A cable placed onto the plugboard connected letters in pairs. Can be used up to ten of these cables.
- Allowed values: letters "ABCDEFGHIJKLMNOPQRSTUVWXYZ" in pairs, and cannot repeat letters
- 0 to 10 paired letters are used

## Notes

- Wikipedia Enigma [page](https://en.wikipedia.org/wiki/Enigma_machine)
- How did the Enigma Machine work? [Jared Owen video](https://www.youtube.com/watch?v=ybkkiGtJmkM)
- Cracking Enigma [Computerphile video](https://www.youtube.com/watch?v=RzWB5jL5RX0)
  and [Java code](https://github.com/mikepound/enigma)

